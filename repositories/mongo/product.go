package mongo

import (
	"myproject/database/connection"
	"myproject/models/mongo"
)

type ProductRepository struct {
	BaseRepository[mongo.Product]
}

func NewProductRepository(db *connection.MongoDatabase) *ProductRepository {
	return &ProductRepository{
		BaseRepository: *NewBaseRepository[mongo.Product](db, mongo.Product{}.CollectionName()),
	}
}
