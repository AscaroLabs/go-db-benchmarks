package main

import (
	"benchmarks/setup"
	"log"
)

func main() {
	log.Println("setup")
	defer log.Println("setup done")
	err := setup.Postgres()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("postgres done")
	err = setup.Tarantool()
	if err != nil {
		log.Fatalln(err)
	}
}
