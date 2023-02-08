package main

import (
	"benchmarks/actions"
	"benchmarks/env"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	g "benchmarks/generator"

	"github.com/aeroideaservices/gormtnt"
	_ "github.com/aeroideaservices/tnt"
	"gorm.io/gorm"

	"github.com/tarantool/go-tarantool"
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
	case ModeExecute:
		conn, err := tarantool.Connect(fmt.Sprintf("%s:%s", env.TntHost, env.TntPort), tarantool.Opts{
			User: env.TntUser,
			Pass: env.TntPassword,
		})
		if err != nil {
			log.Fatalln(err)
		}
		start = time.Now()
		for i := 0; i < env.Workers; i++ {
			wg.Add(1)
			go actions.TarantoolExecute(&wg, conn, c)
		}
	case ModeIsert:
		conn, err := tarantool.Connect(fmt.Sprintf("%s:%s", env.TntHost, env.TntPort), tarantool.Opts{
			User: env.TntUser,
			Pass: env.TntPassword,
		})
		if err != nil {
			log.Fatalln(err)
		}
		start = time.Now()
		for i := 0; i < env.Workers; i++ {
			wg.Add(1)
			go actions.TarantoolInsert(&wg, conn, c)
		}
	case ModeSql:
		db, err := sql.Open("tnt", fmt.Sprintf("tarantool://%s:%s@%s:%s",
			env.TntUser, env.TntPassword, env.TntHost, env.TntPort))
		if err != nil {
			log.Fatalln(err)
		}
		start = time.Now()
		for i := 0; i < env.Workers; i++ {
			wg.Add(1)
			go actions.Tnt(&wg, db, c)
		}
	case ModeGorm:
		db, err := gorm.Open(gormtnt.Open("tarantool://admin:aeroportal-tarantool@51.250.8.61:3305"), &gorm.Config{
			SkipDefaultTransaction: true,
			// DisableNestedTransaction: true,
		})
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
