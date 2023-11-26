package database

import "gorm.io/gorm"

type Model[M any] struct {
	Stg *gorm.DB
}

func (m *Model[M]) Get(where ...interface{}) (M, error) {
	var model M
	if err := m.Stg.First(&model, where...); err.Error != nil {
		return model, err.Error
	}

	return model, nil
}

func (m *Model[M]) GetList() ([]M, error) {
	var models []M
	if err := m.Stg.Find(&models); err.Error != nil {
		return nil, err.Error
	}

	return models, nil
}

func (m *Model[M]) Insert(model M) error {
	if err := m.Stg.Create(&model); err.Error != nil {
		return err.Error
	}

	return nil
}

func (m *Model[M]) Update(updated_model M, conds interface{}) (M, error) {
	var model M
	if err := m.Stg.Model(&updated_model).Where(conds).Updates(updated_model); err.Error != nil {
		return model, err.Error
	}

	return updated_model, nil
}

func (m *Model[M]) Delete(id string) (M, error) {
	var model M
	if err := m.Stg.Where("id = ?", id).Delete(&model); err.Error != nil {
		return model, err.Error
	}

	return model, nil
}
