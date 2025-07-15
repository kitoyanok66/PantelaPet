package handlers

import (
	taskservice "PantelaPet/internal/taskService"
	"PantelaPet/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	service taskservice.TaskService
}

func NewTaskHandler(s taskservice.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *TaskHandler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskservice.TaskRequest{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

// PatchTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	taskToUpdate := taskservice.TaskRequest{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	updatedTask, err := h.service.UpdateTask(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

// DeleteTasksId implements tasks.StrictServerInterface.
func (h *TaskHandler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	if err := h.service.DeleteTask(taskID); err != nil {
		return nil, err
	}

	return tasks.DeleteTasksId204Response{}, nil
}
