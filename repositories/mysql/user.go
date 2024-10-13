package mysql

import (
	"myproject/database/connection"
	"myproject/models/mysql"
)

type UserRepository struct {
	BaseRepository[mysql.User]
}

func NewUserRepository(db *connection.MysqlDatabase) *UserRepository {
	return &UserRepository{
		BaseRepository: *NewBaseRepository[mysql.User](db),
	}
}
