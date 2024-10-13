package factories

import (
	"myproject/helpers"
	"myproject/models/mysql"
)

func newProduct() *mysql.Product {
	return &mysql.Product{
		Name:             helpers.GenerateLetter(8),
		ChannelProductId: helpers.GenerateString(14),
		CommonColumn: mysql.CommonColumn{
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
	}
}

func DefinitionProduct() *mysql.Product {
	return newProduct()
}

func SeedProduct(number int) []*mysql.Product {
	var data []*mysql.Product

	for i := 0; i < number; i++ {
		data = append(data, DefinitionProduct())
	}

	return data
}
