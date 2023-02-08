package setup

import (
	"benchmarks/env"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	g "benchmarks/generator"
)

func Postgres() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		env.DbHost, env.DbUser, env.DbPassword, env.DbName, env.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&g.User{})
	if err != nil {
		return err
	}
	return nil
}
