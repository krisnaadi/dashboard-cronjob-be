package user

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (repository *Repository) GetUserByID(ctx context.Context, ID int64) (entity.User, error) {
	var user entity.User

	err := repository.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "repository.db.Where().Find() error - GetUserByID")
		return entity.User{}, err
	}

	return user, nil
}

func (repository *Repository) GetUserByEmail(ctx context.Context, Email string) (entity.User, error) {
	var user entity.User

	err := repository.db.Where("email = ?", Email).Find(&user).Error
	if err != nil {
		logger.Trace(ctx, struct{ Email string }{Email}, err, "repository.db.Where().Find() error - GetUserByEmail")
		return entity.User{}, err
	}

	return user, nil
}

func (repository *Repository) GetUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := repository.db.Order("id asc").Find(&users).Error

	if err != nil {
		logger.Trace(ctx, nil, err, "repository.db.Order().Find() error - GetUsers")
		return nil, err
	}

	return users, nil
}

func (repository *Repository) InsertUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		logger.Trace(ctx, user, err, "repository.db.Create() error - InsertUser")
		return entity.User{}, err
	}

	return user, nil
}

func (repository *Repository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := repository.db.Save(&user).Error
	if err != nil {
		logger.Trace(ctx, user, err, "repository.db.Save() error - UpdateUser")
		return entity.User{}, err
	}

	return user, nil
}

func (repository *Repository) DeleteUser(ctx context.Context, ID int64) error {
	err := repository.db.Delete(&entity.User{}, ID).Error
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "repository.db.Delete() error - DeleteUser")
		return err
	}

	return nil
}
