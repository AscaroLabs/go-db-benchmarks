package main

import (
	"benchmarks/actions"
	"benchmarks/env"
	g "benchmarks/generator"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	_ "github.com/aeroideaservices/tnt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	ModeExecute = iota
	ModeIsert
	ModeSql
	ModeGorm
)

func main() {

	c := make(chan g.User, g.DATA_LIMIT)

	g.Generator(c)

	var wg sync.WaitGroup

	log.Println("start")
	defer log.Println("done")

	var start time.Time

	switch env.Mode {
	case ModeSql:
		db, err := sql.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.DbName))
		if err != nil {
			log.Fatalln(err)
		}
		start = time.Now()
		for i := 0; i < env.Workers; i++ {
			wg.Add(1)
			go actions.Pgx(&wg, db, c)
		}
	case ModeGorm:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			env.DbHost, env.DbUser, env.DbPassword, env.DbName, env.DbPort)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		start = time.Now()
		for i := 0; i < env.Workers; i++ {
			wg.Add(1)
			go actions.TarantoolGorm(&wg, db, c)
		}
	}

	wg.Wait()

	log.Printf("%s\n", time.Since(start))
}
