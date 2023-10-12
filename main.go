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

	printMessages(repository.GetAll())
	printMessages(repository.GetDuplicatesJoin())
	printMessages(repository.GetDuplicatesCte())
}

func printMessages(messages []db.Entity) {
	fmt.Println("-----------------------------------------------------------------------------")
	for _, v := range messages {
		fmt.Println(v)
	}
}
