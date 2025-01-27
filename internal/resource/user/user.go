package user

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (resource *Resource) GetUserByID(ctx context.Context, ID int64) (entity.User, error) {
	user, err := resource.user.GetUserByID(ctx, ID)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "resource.user.GetUserByID() error - GetUserByID")
		return user, err
	}

	return user, nil
}

func (resource *Resource) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user, err := resource.user.GetUserByEmail(ctx, email)
	if err != nil {
		logger.Trace(ctx, struct{ email string }{email}, err, "resource.user.GetUserByEmail() error - GetUserByEmail")
		return user, err
	}

	return user, nil
}

func (resource *Resource) GetUsers(ctx context.Context) ([]entity.User, error) {
	users, err := resource.user.GetUsers(ctx)
	if err != nil {
		logger.Trace(ctx, nil, err, "resource.user.GetUsers() error - GetUsers")
		return users, err
	}

	return users, nil
}

func (resource *Resource) AddUser(ctx context.Context, user entity.User) (entity.User, error) {
	user, err := resource.user.InsertUser(ctx, user)
	if err != nil {
		logger.Trace(ctx, user, err, "resource.user.InsertUser() error - AddUser")
		return user, err
	}

	return user, nil
}

func (resource *Resource) UpdateUser(ctx context.Context, ID int64, user entity.User) (entity.User, error) {
	newUser, err := resource.user.UpdateUser(ctx, user)
	if err != nil {
		logger.Trace(ctx, user, err, "resource.user.UpdateUser() error - UpdateUser")
		return newUser, err
	}

	return newUser, nil
}

func (resource *Resource) DeleteUser(ctx context.Context, ID int64) error {

	err := resource.user.DeleteUser(ctx, ID)
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "resource.user.DeleteUser() error - DeleteUser")
		return err
	}

	return nil
}
