package main

import (
	"fmt"
	"myproject/database/connection"
	"myproject/helpers"
	mongomodels "myproject/models/mongo"
	mysqlmodels "myproject/models/mysql"
	mongorepositories "myproject/repositories/mongo"
	mysqlrepositories "myproject/repositories/mysql"
)

func main() {
	dbMysql, err := connection.ConnectMysql()
	if err != nil {
		fmt.Printf("Could not connect to the database: %v\n", err)
		return
	}

	defer func() {
		if err = dbMysql.Close(); err != nil {
			fmt.Printf("Error closing the database connections: %v\n", err)
		} else {
			fmt.Println("Closed the database connections")
		}
	}()

	// Call the seed function
	seedUsersMysql(dbMysql)
	seedProductsMysql(dbMysql)

	dbMongo, err := connection.ConnectMongo()
	if err != nil {
		fmt.Printf("Could not connect to the database: %v\n", err)
		return
	}

	defer func() {
		if err = dbMongo.Close(); err != nil {
			fmt.Printf("Error closing the database connection: %v\n", err)
		} else {
			fmt.Println("Closed the database connection")
		}
	}()

	// Call the seed function
	seedStorefrontsMongo(dbMongo)
	seedProductsMongo(dbMongo)

}

func seedStorefrontsMongo(db *connection.MongoDatabase) {
	documents := []*mongomodels.Storefront{
		&mongomodels.Storefront{
			Name:                helpers.GenerateLetter(6),
			ChannelStorefrontId: helpers.GenerateString(6),
			//ChannelStorefrontId: "28FQ82",
			Data:      "{'key1':'value1'}",
			ExtraData: "{'key2':'value2'}",
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
	}

	storefrontRepo := mongorepositories.NewStorefrontRepository(db)

	for _, document := range documents {
		result, err := storefrontRepo.Insert(document)
		if err != nil {
			fmt.Printf("Error seeding storefront %s: %v\n", document.ChannelStorefrontId, err)
		} else {
			fmt.Printf("Seeded storefront: %s with %d\n", document.ChannelStorefrontId, result.InsertedID)
		}
	}

	bulkDocuments := []interface{}{
		&mongomodels.Storefront{
			Name:                helpers.GenerateLetter(6),
			ChannelStorefrontId: helpers.GenerateString(6),
			//ChannelStorefrontId: "NkFlL5",
			Data:      "{'key1':'value1'}",
			ExtraData: "{'key2':'value2'}",
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
		&mongomodels.Storefront{
			Name:                helpers.GenerateLetter(6),
			ChannelStorefrontId: helpers.GenerateString(6),
			Data:                "{'key1':'value1'}",
			ExtraData:           "{'key2':'value2'}",
			CreatedAt:           helpers.GetTimeNow(),
			UpdatedAt:           helpers.GetTimeNow(),
		},
	}

	if err := storefrontRepo.InsertMany(bulkDocuments); err != nil {
		fmt.Printf("Error seeding storefronts: %v\n", err)
	} else {
		fmt.Println("Seeded storefronts successfully")
	}

}

func seedProductsMongo(db *connection.MongoDatabase) {
	documents := []*mongomodels.Product{
		&mongomodels.Product{
			Name:             helpers.GenerateLetter(6),
			ChannelProductId: helpers.GenerateString(6),
			//ChannelProductId: "6zLk7X",
			Data:      "{'key1':'value1'}",
			ExtraData: "{'key2':'value2'}",
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
	}

	productRepo := mongorepositories.NewProductRepository(db)

	for _, document := range documents {
		result, err := productRepo.Insert(document)
		if err != nil {
			fmt.Printf("Error seeding product %s: %v\n", document.ChannelProductId, err)
		} else {
			fmt.Printf("Seeded product: %s with %d\n", document.ChannelProductId, result.InsertedID)
		}
	}

	bulkDocuments := []interface{}{
		&mongomodels.Product{
			Name:             helpers.GenerateLetter(6),
			ChannelProductId: helpers.GenerateString(6),
			//ChannelProductId: "LmC5CV",
			Data:      "{'key1':'value1'}",
			ExtraData: "{'key2':'value2'}",
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
		&mongomodels.Product{
			Name:             helpers.GenerateLetter(6),
			ChannelProductId: helpers.GenerateString(6),
			Data:             "{'key1':'value1'}",
			ExtraData:        "{'key2':'value2'}",
			CreatedAt:        helpers.GetTimeNow(),
			UpdatedAt:        helpers.GetTimeNow(),
		},
	}

	if err := productRepo.InsertMany(bulkDocuments); err != nil {
		fmt.Printf("Error seeding products: %v\n", err)
	} else {
		fmt.Println("Seeded products successfully")
	}

}

func seedUsersMysql(db *connection.MysqlDatabase) {
	records := []*mysqlmodels.User{
		&mysqlmodels.User{
			//Id:        1,
			CompanyId: 100,
			Name:      helpers.GenerateLetter(6),
			Email:     helpers.GenerateString(6) + "_user@example.com",
			CommonColumn: mysqlmodels.CommonColumn{
				CreatedAt: helpers.GetTimeNow(),
				UpdatedAt: helpers.GetTimeNow(),
			},
		},
	}

	userRepo := mysqlrepositories.NewUserRepository(db)

	for _, record := range records {
		recordOutput, err := userRepo.Create(record)
		if err != nil {
			fmt.Printf("Error seeding user %s: %v\n", record.Email, err)
		} else {
			fmt.Printf("Seeded user: %s with %d\n", record.Email, recordOutput.Id)
		}
	}

	records = []*mysqlmodels.User{
		&mysqlmodels.User{
			//Id:        2,
			CompanyId: 100,
			Name:      helpers.GenerateLetter(6),
			Email:     helpers.GenerateString(6) + "_user@example.com",
			CommonColumn: mysqlmodels.CommonColumn{
				CreatedAt: helpers.GetTimeNow(),
				UpdatedAt: helpers.GetTimeNow(),
			},
		},
	}

	if err := userRepo.CreateMany(records); err != nil {
		fmt.Printf("Error seeding users: %v\n", err)
	} else {
		fmt.Println("Seeded users successfully")
	}

}

func seedProductsMysql(db *connection.MysqlDatabase) {
	records := []*mysqlmodels.Product{
		&mysqlmodels.Product{
			//Id:               1,
			Name:             helpers.GenerateLetter(6),
			ChannelProductId: helpers.GenerateString(8),
			CommonColumn: mysqlmodels.CommonColumn{
				CreatedAt: helpers.GetTimeNow(),
				UpdatedAt: helpers.GetTimeNow(),
			},
		},
	}

	productRepo := mysqlrepositories.NewProductRepository(db)

	for _, record := range records {
		recordOutput, err := productRepo.Create(record)
		if err != nil {
			fmt.Printf("Error seeding product %s: %v\n", record.Name, err)
		} else {
			fmt.Printf("Seeded product: %s with %d\n", record.Name, recordOutput.Id)
		}
	}

	records = []*mysqlmodels.Product{
		&mysqlmodels.Product{
			//Id:               2,
			Name:             helpers.GenerateLetter(6),
			ChannelProductId: helpers.GenerateString(8),
			CommonColumn: mysqlmodels.CommonColumn{
				CreatedAt: helpers.GetTimeNow(),
				UpdatedAt: helpers.GetTimeNow(),
			},
		},
	}

	if err := productRepo.CreateMany(records); err != nil {
		fmt.Printf("Error seeding products: %v\n", err)
	} else {
		fmt.Println("Seeded products successfully")
	}

}
