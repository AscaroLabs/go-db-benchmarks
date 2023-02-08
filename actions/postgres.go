package actions

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	g "benchmarks/generator"

	"gorm.io/gorm"
)

func Pgx(wg *sync.WaitGroup, db *sql.DB, c <-chan g.User) {
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

func PostgresGorm(wg *sync.WaitGroup, db *gorm.DB, c <-chan g.User) {
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
