package mysql

import (
	"gorm.io/gorm"
	"myproject/database/connection"
)

type Repository[model any] interface {
	Create(record *model) (*model, error)
	CreateMany(records []*model) error
}

type BaseRepository[model any] struct {
	db *connection.MysqlDatabase
}

func NewBaseRepository[model any](db *connection.MysqlDatabase) *BaseRepository[model] {
	return &BaseRepository[model]{db: db}
}

func (repository *BaseRepository[model]) Create(record *model) (*model, error) {
	if err := repository.db.Master.Create(record).Error; err != nil {
		return nil, err
	}

	return record, nil
}

func (repository *BaseRepository[model]) CreateMany(records []*model) error {
	return repository.db.Master.Transaction(func(tran *gorm.DB) error {
		if err := tran.Create(records).Error; err != nil {
			return err
		}

		return nil
	})
}
