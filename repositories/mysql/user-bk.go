package mysql

import (
	"gorm.io/gorm"
	"myproject/database/connection"
	"myproject/models/mysql"
)

type UserBkRepository struct {
	db *connection.MysqlDatabase
}

func NewUserBkRepository(db *connection.MysqlDatabase) *UserBkRepository {
	return &UserBkRepository{db: db}
}

func (repository *UserBkRepository) Create(record *mysql.User) (uint64, error) {
	if err := repository.db.Master.Create(record).Error; err != nil {
		return 0, err
	}

	return record.Id, nil
}

func (repository *UserBkRepository) CreateMany(records []*mysql.User) error {
	return repository.db.Master.Transaction(func(tran *gorm.DB) error {
		if err := tran.Create(records).Error; err != nil {
			return err
		}

		return nil
	})
}
