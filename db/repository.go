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
	Path     string
	Logger   *logrus.Logger
}

type Repository interface {
	Insert(e Entity) (operation string, rowsAffected int64)
	GetAll() (operation string, entities []Entity)
	GetDuplicatesJoin() (operation string, entities []Entity)
	GetDuplicatesCte() (operation string, entities []Entity)
	Truncate() (operation string, rowsAffected int64)
}

type repository struct {
	db     *sql.DB
	logger *logrus.Logger
	path   string // to help to locate the sql files no matter if code is executed from main or the test files
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
		path:   c.Path,
	}
}

func (r repository) Insert(e Entity) (operation string, rowsAffected int64) {
	op := "Insert"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/insert.sql", r.path)); err == nil {
		res, err := r.db.ExecContext(context.TODO(), string(query), e.MessageType, e.SpaceType, e.ReceiverId, e.LogData, e.CreateTime)
		if err != nil {
			r.logger.Error(fmt.Sprintf("%s DB query failed!", operation))
			return op, 0
		}
		rows, _ := res.RowsAffected()

		return op, rows
	}
	return op, 0
}

func (r repository) GetAll() (operation string, entities []Entity) {
	op := "GetAll"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/get-all.sql", r.path)); err == nil {
		return r.execQuery(op, string(query))
	}
	return op, nil
}

func (r repository) GetDuplicatesJoin() (operation string, entities []Entity) {
	op := "GetDuplicatesJoin"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/get-duplicates-join.sql", r.path)); err == nil {
		return r.execQuery(op, string(query))
	}
	return op, nil
}

func (r repository) GetDuplicatesCte() (operation string, entities []Entity) {
	op := "GetDuplicatesCte"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/get-duplicates-cte.sql", r.path)); err == nil {
		return r.execQuery(op, string(query))
	}
	return op, nil
}

func (r repository) execQuery(op, sql string) (operation string, entities []Entity) {
	rows, err := r.db.QueryContext(context.TODO(), sql)
	if err != nil {
		r.logger.Error(fmt.Sprintf("%s DB query failed!", operation))
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

func (r repository) Truncate() (operation string, rowsAffected int64) {
	op := "Truncate"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/truncate.sql", r.path)); err == nil {
		res, err := r.db.ExecContext(context.TODO(), string(query))
		if err != nil {
			r.logger.Error(fmt.Sprintf("%s DB query failed!", op))
		}
		rows, _ := res.RowsAffected()

		return op, rows
	}
	return op, 0
}
