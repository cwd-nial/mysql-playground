package tests

import (
	"database/sql"
	"encoding/csv"
	"github.com/cwd-nial/mysql-playground/db"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

const mysqlTime = "2006-01-02 15:04:05"

type Framework struct {
	t          *testing.T
	repository db.Repository
}

func newFramework(t *testing.T) Framework {
	p, err := os.Getwd()
	if err != nil {
		panic("failed to get global path")
	}
	mp := strings.ReplaceAll(p, "\\tests", "")

	fr := Framework{
		t: t,
		repository: db.NewRepository(db.Config{
			Uri:      "localhost",
			User:     "neo",
			Password: "white_rabbit",
			Path:     mp,
			Logger:   logrus.New(),
		}),
	}
	fr.repository.UseDatabase()
	fr.tearDownDb()
	fr.initDb()

	return fr
}

func (fr *Framework) initDb() {
	f, err := os.Open("data.csv")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	lines, err := csv.NewReader(f).ReadAll()
	for _, v := range lines {
		receiverId, _ := strconv.Atoi(v[2])
		t, err := time.Parse(mysqlTime, v[4])
		if err != nil {
			panic("failed to parse the mysql value from the CSV file")
		}
		fr.repository.Insert(db.Entity{
			MessageType: v[0],
			SpaceType:   v[1],
			ReceiverId:  uint64(receiverId),
			LogData:     v[3],
			CreateTime:  sql.NullTime{t, true},
		})
	}
}

func (fr *Framework) tearDownDb() {
	fr.repository.Truncate()
}
