package actions

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	g "benchmarks/generator"

	"github.com/tarantool/go-tarantool"
	"gorm.io/gorm"
)

var timeout = 1 * time.Second

func TarantoolExecute(wg *sync.WaitGroup, conn *tarantool.Connection, c <-chan g.User) {
	defer wg.Done()
	var u g.User
	for {
		select {
		case u = <-c:
			_, err := conn.Execute(`insert into "Test" values (?, ?, ?, ?, ?)`, []interface{}{u.ID, u.Name1, u.Name2, u.Name3, u.Name4})
			if err != nil {
				log.Fatalln(err)
			}
		case <-time.After(timeout):
			return
		}
	}
}

func TarantoolInsert(wg *sync.WaitGroup, conn *tarantool.Connection, c <-chan g.User) {
	defer wg.Done()
	var u g.User
	for {
		select {
		case u = <-c:
			_, err := conn.Insert(interface{}("Test"), []interface{}{u.ID, u.Name1, u.Name2, u.Name3, u.Name4})
			if err != nil {
				log.Fatalln(err)
			}
		case <-time.After(timeout):
			return
		}
	}
}

func Tnt(wg *sync.WaitGroup, db *sql.DB, c <-chan g.User) {
	defer wg.Done()
	var u g.User
	for {
		select {
		case u = <-c:
			_, err := db.ExecContext(context.Background(), `insert into "Test" values (?, ?, ?, ?)`, u.ID, u.Name1, u.Name2, u.Name3, u.Name4)
			if err != nil {
				log.Fatalln(err)
			}
		case <-time.After(timeout):
			return
		}
	}
}

func TarantoolGorm(wg *sync.WaitGroup, db *gorm.DB, c <-chan g.User) {
	defer wg.Done()
	var u g.User
	for {
		select {
		case u = <-c:
			err := db.Table("Test").Create(&u).Error
			if err != nil {
				log.Fatalln(err)
			}
		case <-time.After(timeout):
			return
		}
	}
}
