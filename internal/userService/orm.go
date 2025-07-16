package userservice

import taskservice "PantelaPet/internal/taskService"

type User struct {
	ID       string             `gorm:"primaryKey" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskservice.Task `gorm:"foreignKey:UserID"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
