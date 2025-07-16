package userservice

import (
	"errors"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req UserRequest) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id string, req UserRequest) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) validateUserRequest(req UserRequest) error {
	if req.Email == "" || req.Password == "" {
		return errors.New("email and password cannot be empty")
	}

	if !strings.Contains(req.Email, "@") && !strings.Contains(req.Email, ".") {
		return errors.New("invalid email format")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	if strings.Contains(req.Password, " ") {
		return errors.New("password must not contain spaces")
	}

	var hasLetter, hasDigit bool
	for _, ch := range req.Password {
		switch {
		case unicode.IsLetter(ch):
			hasLetter = true
		case unicode.IsDigit(ch):
			hasDigit = true
		}
	}
	if !hasLetter || !hasDigit {
		return errors.New("password must contain both letters and digits")
	}

	return nil
}

func (s *userService) CreateUser(req UserRequest) (User, error) {
	if err := s.validateUserRequest(req); err != nil {
		return User{}, err
	}

	user := User{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id string) (User, error) {
	if id == "" {
		return User{}, errors.New("user ID is required")
	}
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id string, req UserRequest) (User, error) {
	if id == "" {
		return User{}, errors.New("user ID is required")
	}

	if err := s.validateUserRequest(req); err != nil {
		return User{}, err
	}

	user := User{
		ID:       id,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("user ID is required")
	}
	return s.repo.DeleteUser(id)
}
