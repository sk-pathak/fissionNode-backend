package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sk-pathak/fissionNode-backend/internal/app/repository"
	"github.com/sk-pathak/fissionNode-backend/internal/db"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *db.User) error {
	_, err := s.repo.GetUser(ctx, user.ID)
	if err != nil {
		return errors.New("user already exists")
	}

	if err := s.repo.SaveUser(ctx, user); err != nil {
		return errors.New("failed to create user in repository: " + err.Error())
	}
	return nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, errors.New("failed to retrieve users from repository: " + err.Error())
	}
	return users, nil
}

func (s *UserService) GetUser(ctx context.Context, id pgtype.UUID) (db.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return db.User{}, errors.New("failed to retrieve user from repository: " + err.Error())
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id pgtype.UUID) error {
	_, err := s.repo.GetUser(ctx, id)
	if err!=nil {
		return errors.New("user doesn't exist")
	}
	
	err = s.repo.DeleteUser(ctx, id)

	return err
}