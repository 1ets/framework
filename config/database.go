package config

import (
	"framework/app/repository"
	"os"

	"github.com/1ets/lets/drivers"
	"github.com/1ets/lets/types"
)

func Database() {
	drivers.MySQLConfig = &types.MySQL{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_DATABASE"),
		Charset:   "utf8mb4",
		ParseTime: "True",
		Loc:       "Local",
		Debug:     true,

		// Depedency injection
		Repository: repository.User,
	}
}
