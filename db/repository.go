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
	UseDatabase()
	Insert(e Entity) (rowsAffected int)
	GetAll() (entities []Entity)
	GetDuplicatesJoin() (entities []Entity)
	GetDuplicatesCte() (entities []Entity)
	DeleteDuplicatesJoin() (rowsAffected int)
	DeleteDuplicatesCte() (rowsAffected int)
	Truncate() (rowsAffected int)
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

func (r repository) UseDatabase() {
	op := "UseDatabase"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/base/use-database.sql", r.path)); err == nil {
		r.exec(op, string(query))
	}
}

func (r repository) Insert(e Entity) (rowsAffected int) {
	op := "Insert"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/base/insert.sql", r.path)); err == nil {
		res, err := r.db.ExecContext(context.TODO(), string(query), e.MessageType, e.SpaceType, e.ReceiverId, e.LogData, e.CreateTime)
		if err != nil {
			r.logger.Error(fmt.Sprintf("%s DB query failed!", op))
			return 0
		}
		rows, _ := res.RowsAffected()

		return int(rows)
	}
	return 0
}

func (r repository) GetAll() (entities []Entity) {
	op := "GetAll"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/base/get-all.sql", r.path)); err == nil {
		return r.queryContext(op, string(query))
	}
	return nil
}

func (r repository) GetDuplicatesJoin() (entities []Entity) {
	op := "GetDuplicatesJoin"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/get-duplicates-join.sql", r.path)); err == nil {
		return r.queryContext(op, string(query))
	}
	return nil
}

func (r repository) GetDuplicatesCte() (entities []Entity) {
	op := "GetDuplicatesCte"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/get-duplicates-cte.sql", r.path)); err == nil {
		return r.queryContext(op, string(query))
	}
	return nil
}

func (r repository) DeleteDuplicatesJoin() (rowsAffected int) {
	op := "DeleteDuplicatesJoin"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/delete-duplicates-join.sql", r.path)); err == nil {
		return r.execContext(op, string(query))
	}
	return 0
}

func (r repository) DeleteDuplicatesCte() (rowsAffected int) {
	op := "DeleteDuplicatesCte"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/delete-duplicates-cte.sql", r.path)); err == nil {
		return r.execContext(op, string(query))
	}
	return 0
}

func (r repository) Truncate() (rowsAffected int) {
	op := "Truncate"
	if query, err := os.ReadFile(fmt.Sprintf("%s/sql/base/truncate.sql", r.path)); err == nil {
		return r.execContext(op, string(query))
	}
	return 0
}

func (r repository) queryContext(op, sql string) (entities []Entity) {
	rows, err := r.db.QueryContext(context.TODO(), sql)
	if err != nil {
		r.logger.Error(fmt.Sprintf("%s DB query failed!", op))
		return nil
	}
	var res []Entity
	for rows.Next() {
		e := Entity{}
		err := rows.Scan(&e.Id, &e.MessageType, &e.SpaceType, &e.ReceiverId, &e.LogData, &e.CreateTime)
		if err != nil {
			r.logger.Error(err)
			return nil
		}

		res = append(res, e)
	}

	return res
}

func (r repository) execContext(op, sql string) (rowsAffected int) {
	res, err := r.db.ExecContext(context.TODO(), sql)
	if err != nil {
		r.logger.Error(fmt.Sprintf("%s DB query failed!", op))
	}
	rows, _ := res.RowsAffected()

	return int(rows)
}

func (r repository) exec(op, sql string) {
	_, err := r.db.Exec(sql)
	if err != nil {
		r.logger.Error(fmt.Sprintf("%s DB query failed!", op))
	}
}
