package main

import (
	"fmt"
	"github.com/cwd-nial/mysql-playground/db"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	var logger = logrus.New()

	p, err := os.Getwd()
	if err != nil {
		panic("failed to get global path")
	}

	repository := db.NewRepository(db.Config{
		Uri:      "localhost",
		User:     "neo",
		Password: "white_rabbit",
		Path:     p,
		Logger:   logger,
	})

	fmt.Println("\nget all entries")
	printMessages(repository.GetAll())

	fmt.Println("\nget duplicates by join")
	printMessages(repository.GetDuplicatesJoin())

	fmt.Println("\nget duplicates with CTE")
	printMessages(repository.GetDuplicatesCte())
}

func printMessages(messages []db.Entity) {
	fmt.Println("=============================================================================")
	for _, v := range messages {
		fmt.Println(v)
	}
	fmt.Println("-----------------------------------------------------------------------------")
}
