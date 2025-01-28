package cronjob

import (
	"context"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/entity"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
)

func (repository *Repository) GetCronjobByID(ctx context.Context, ID int64, UserId int64) (entity.Cronjob, error) {
	var cronjob entity.Cronjob

	err := repository.db.Where("id = ?", ID).Where("user_id = ?", UserId).Find(&cronjob).Error
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "repository.db.Where().Find() error - GetCronjobByID")
		return entity.Cronjob{}, err
	}

	return cronjob, nil
}

func (repository *Repository) GetCronjobs(ctx context.Context, UserId int64) ([]entity.Cronjob, error) {
	var cronjobs []entity.Cronjob
	err := repository.db.Where("user_id = ?", UserId).Order("id asc").Find(&cronjobs).Error

	if err != nil {
		logger.Trace(ctx, nil, err, "repository.db.Order().Find() error - GetCronjobs")
		return nil, err
	}

	return cronjobs, nil
}

func (repository *Repository) InsertCronjob(ctx context.Context, cronjob entity.Cronjob) (entity.Cronjob, error) {
	err := repository.db.Create(&cronjob).Error
	if err != nil {
		logger.Trace(ctx, cronjob, err, "repository.db.Create() error - InsertCronjob")
		return entity.Cronjob{}, err
	}

	return cronjob, nil
}

func (repository *Repository) UpdateCronjob(ctx context.Context, cronjob entity.Cronjob) (entity.Cronjob, error) {
	err := repository.db.Save(&cronjob).Error
	if err != nil {
		logger.Trace(ctx, cronjob, err, "repository.db.Save() error - UpdateCronjob")
		return entity.Cronjob{}, err
	}

	return cronjob, nil
}

func (repository *Repository) DeleteCronjob(ctx context.Context, ID int64, UserId int64) error {
	err := repository.db.Where("user_id = ?", UserId).Delete(&entity.Cronjob{}).Error
	if err != nil {
		logger.Trace(ctx, struct{ ID int64 }{ID}, err, "repository.db.Delete() error - DeleteCronjob")
		return err
	}

	return nil
}
