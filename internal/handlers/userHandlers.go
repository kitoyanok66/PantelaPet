package handlers

import (
	userservice "PantelaPet/internal/userService"
	"PantelaPet/internal/web/users"
	"context"
)

type UserHandler struct {
	service userservice.UserService
}

func NewUserHandler(s userservice.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// GetUsers implements users.StrictServerInterface.
func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		user := users.User{
			Id:       &user.ID,
			Email:    &user.Email,
			Password: &user.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := userservice.UserRequest{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.service.CreateUser(userToCreate)
	if err != nil {
		errMsg := err.Error()
		return users.PostUsers422JSONResponse{
			Error: &errMsg,
		}, nil
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

// PatchUsersId implements users.StrictServerInterface.
func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userID := request.Id
	userRequest := request.Body

	userToUpdate := userservice.UserRequest{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	updatedUser, err := h.service.UpdateUser(userID, userToUpdate)
	if err != nil {
		errMsg := err.Error()
		return users.PatchUsersId422JSONResponse{
			Error: &errMsg,
		}, nil
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

// DeleteUsersId implements users.StrictServerInterface.
func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id

	if err := h.service.DeleteUser(userID); err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
