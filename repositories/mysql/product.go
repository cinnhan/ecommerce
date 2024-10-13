package mysql

import (
	"myproject/database/connection"
	"myproject/models/mysql"
)

type ProductRepository struct {
	BaseRepository[mysql.Product]
}

func NewProductRepository(db *connection.MysqlDatabase) *ProductRepository {
	return &ProductRepository{
		BaseRepository: *NewBaseRepository[mysql.Product](db),
	}
}
