package repository

import (
	"errors"
	"eventori/internal/schedule"

	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) schedule.Repository {
	return &scheduleRepository{db}
}

func (sr *scheduleRepository) GetByModelID(modelID int) ([]schedule.Core, error) {
	scheduleModels := []ModelSchedule{}
	tx := sr.db.Where("Model_id = ?", modelID).Find(&scheduleModels)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return ToCores(scheduleModels...), nil
}

func (sr *scheduleRepository) Create(scheduleCore schedule.Core) error {
	scheduleModel := ToModel(scheduleCore)
	tx := sr.db.Create(&scheduleModel)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (sr *scheduleRepository) Update(scheduleID int, scheduleCore schedule.Core) error {
	scheduleModel := ToModel(scheduleCore)
	tx := sr.db.Where("Id = ?", scheduleID).Updates(&scheduleModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed because no rows affected")
	}

	return nil
}

func (sr *scheduleRepository) Delete(scheduleID int) error {
	tx := sr.db.Delete(&ModelSchedule{}, scheduleID)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed because no rows affected")
	}

	return nil
}
