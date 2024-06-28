package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/celsopires1999/estimation/internal/infra/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	conn    *pgx.Conn
	queries *db.Queries
}

func NewUserService(conn *pgx.Conn) *UserService {
	return &UserService{
		conn:    conn,
		queries: db.New(conn),
	}
}

func (h *UserService) CreateUser(ctx context.Context, input CreateUserInputDTO) (CreateUserOutputDTO, error) {
	_, err := h.queries.GetUserByEmail(ctx, input.Email)
	if err == nil {
		return CreateUserOutputDTO{}, fmt.Errorf("user with email %s already exists", input.Email)
	}
	params := db.CreateUserParams{
		UserID:    uuid.New().String(),
		Email:     input.Email,
		UserName:  input.UserName,
		Name:      input.Name,
		UserType:  input.UserType,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	err = h.queries.CreateUser(ctx, params)
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	created, err := h.queries.GetUser(ctx, params.UserID)
	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	output := CreateUserOutputDTO{
		UserID:    created.UserID,
		Email:     created.Email,
		UserName:  created.UserName,
		Name:      created.Name,
		UserType:  created.UserType,
		CreatedAt: created.CreatedAt.Time,
	}

	return output, nil
}

func (h *UserService) UpdateUser(ctx context.Context, input UpdateUserInputDTO) (UpdateUserOutputDTO, error) {
	current, err := h.queries.GetUser(ctx, input.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return UpdateUserOutputDTO{}, fmt.Errorf("user with id: %s not found", input.UserID)
		}
		return UpdateUserOutputDTO{}, err
	}
	params := db.UpdateUserParams{}

	params.UserID = current.UserID
	params.UpdatedAt = pgtype.Timestamp{Time: time.Now(), Valid: true}

	params.Email = copyPatch(input.Email, current.Email)
	params.UserName = copyPatch(input.UserName, current.UserName)
	params.Name = copyPatch(input.Name, current.Name)
	params.UserType = copyPatch(input.UserType, current.UserType)

	err = h.queries.UpdateUser(ctx, params)
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	updated, err := h.queries.GetUser(ctx, params.UserID)
	if err != nil {
		return UpdateUserOutputDTO{}, err
	}

	output := UpdateUserOutputDTO{
		UserID:    updated.UserID,
		Email:     updated.Email,
		UserName:  updated.UserName,
		Name:      updated.Name,
		UserType:  updated.UserType,
		CreatedAt: updated.CreatedAt.Time,
		UpdatedAt: updated.UpdatedAt.Time,
	}

	return output, nil
}

func (h *UserService) GetUser(ctx context.Context, userID string) (GetUserOutputDTO, error) {
	current, err := h.queries.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return GetUserOutputDTO{}, fmt.Errorf("user with id: %s not found", userID)
		}
		return GetUserOutputDTO{}, err
	}

	output := GetUserOutputDTO{
		UserID:    current.UserID,
		Email:     current.Email,
		UserName:  current.UserName,
		Name:      current.Name,
		UserType:  current.UserType,
		CreatedAt: current.CreatedAt.Time,
		UpdatedAt: current.UpdatedAt.Time,
	}

	return output, nil
}

func (h *UserService) DeleteUser(ctx context.Context, userID string) error {
	err := h.queries.DeleteUser(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("user with id: %s not found", userID)
		}
		return err
	}
	return nil
}

func copyPatch[T any](input *T, current T) (params T) {
	params = current
	if input != nil {
		params = *input
	}
	return
}

type CreateUserInputDTO struct {
	Email    string `json:"email" validate:"required,email"`
	UserName string `json:"user_name" validate:"required"`
	Name     string `json:"name" validate:"required"`
	UserType string `json:"user_type" validate:"required,oneof=manager estimator"`
}

type CreateUserOutputDTO struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	Name      string    `json:"name"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserInputDTO struct {
	UserID   string  `json:"user_id" validate:"required"`
	Email    *string `json:"email" validate:"omitempty,email"`
	UserName *string `json:"user_name" validate:"omitempty"`
	Name     *string `json:"name" validate:"omitempty"`
	UserType *string `json:"user_type" validate:"omitempty,oneof=manager estimator"`
}

type UpdateUserOutputDTO struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	Name      string    `json:"name"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserOutputDTO struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	UserName  string    `json:"user_name"`
	Name      string    `json:"name"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
