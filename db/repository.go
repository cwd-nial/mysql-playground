package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const readTimeout = 2 * time.Second

type Config struct {
	Uri      string
	User     string
	Password string
	Logger   *logrus.Logger
}

type Repository interface {
	GetAll() (operation string, entities []Entity)
	GetDuplicatesJoin() (operation string, entities []Entity)
	GetDuplicatesCte() (operation string, entities []Entity)
}

type repository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewRepository(c Config) Repository {
	config := mysql.NewConfig()
	config.Addr = c.Uri
	config.User = c.User
	config.Passwd = c.Password
	config.ParseTime = true // IMPORTANT: mysql datetime won't be parsed into time.Time without this!
	config.ReadTimeout = readTimeout

	connector, err := mysql.NewConnector(config)
	if err != nil {
		panic("failed to open mysql connection")
	}

	return &repository{
		db:     sql.OpenDB(connector),
		logger: c.Logger,
	}
}

func (r repository) GetAll() (operation string, entities []Entity) {
	op := "GetAll"
	if query, err := os.ReadFile("sql/get-all.sql"); err == nil {
		return r.execQuery(op, string(query))
	}
	return op, nil
}

func (r repository) GetDuplicatesJoin() (operation string, entities []Entity) {
	op := "GetDuplicatesJoin"
	if query, err := os.ReadFile("sql/get-duplicates-join.sql"); err == nil {
		return r.execQuery(op, string(query))
	}
	return op, nil
}

func (r repository) GetDuplicatesCte() (operation string, entities []Entity) {
	op := "GetDuplicatesCte"
	if query, err := os.ReadFile("sql/get-duplicates-cte.sql"); err == nil {
		return r.execQuery(op, string(query))
	}
	return op, nil
}

func (r repository) execQuery(op, sql string) (operation string, entities []Entity) {
	rows, err := r.db.QueryContext(context.TODO(), sql)
	if err != nil {
		r.logger.Error(fmt.Sprintf("% DB query failed!", operation))
		return op, nil
	}
	var res []Entity
	for rows.Next() {
		e := Entity{}
		err := rows.Scan(&e.Id, &e.MessageType, &e.SpaceType, &e.ReceiverId, &e.LogData, &e.CreateTime)
		if err != nil {
			r.logger.Error(err)
			return op, nil
		}

		res = append(res, e)
	}

	return op, res
}
