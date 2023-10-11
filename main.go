package main

import (
	"fmt"
	"github.com/cwd-nial/mysql-playground/db"
	"github.com/sirupsen/logrus"
)

func main() {
	var logger = logrus.New()

	repository := db.NewRepository(db.Config{
		Uri:      "localhost",
		User:     "neo",
		Password: "white_rabbit",
		Logger:   logger,
	})

	printMessages(repository.GetAll())
	printMessages(repository.GetDuplicatesJoin())
	printMessages(repository.GetDuplicatesCte())
}

func printMessages(op string, messages []db.Entity) {
	fmt.Println("\n=============================================================================")
	fmt.Println(op)
	fmt.Println("-----------------------------------------------------------------------------")
	for _, v := range messages {
		fmt.Println(v)
	}
}
