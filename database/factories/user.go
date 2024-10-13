package factories

import (
	"myproject/helpers"
	"myproject/models/mysql"
)

func newUser(companyId uint64) *mysql.User {
	return &mysql.User{
		CompanyId: companyId,
		Name:      helpers.GenerateLetter(8),
		Email:     helpers.GenerateString(10) + "_user@example.com",
		CommonColumn: mysql.CommonColumn{
			CreatedAt: helpers.GetTimeNow(),
			UpdatedAt: helpers.GetTimeNow(),
		},
	}
}

func DefinitionUser(companyId uint64) *mysql.User {
	return newUser(companyId)
}

func SeedUser(number int, companyId uint64) []*mysql.User {
	//data := make([]*mysql.User, 0)
	var data []*mysql.User

	for i := 0; i < number; i++ {
		data = append(data, DefinitionUser(companyId))
	}

	return data
}
