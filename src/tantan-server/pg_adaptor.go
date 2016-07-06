package main

import (
	"fmt"
	"gopkg.in/pg.v3"
)

type PgAdaptor struct {
	db *pg.DB
}

func NewPgAdaptor(host, port, user, pass, dbname string) (*PgAdaptor, error) {
	dbAdaptor := &PgAdaptor{}
	dbAdaptor.db = pg.Connect(&pg.Options{
		Host:     host,
		Port:     port,
		User:     user,
		Password: pass,
		Database: dbname,
	})
	return dbAdaptor, nil
}

func (this *PgAdaptor) Release() {
	if this.db != nil {
		this.db.Close()
		this.db = nil
	}
}

func (this *PgAdaptor) Query(coll pg.Collection, query string, args ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	_, err := this.db.Query(coll, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (this *PgAdaptor) QueryOne(record interface{}, query string, args ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	_, err := this.db.QueryOne(record, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (this *PgAdaptor) Exec(query string, args ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	_, err := this.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (this *PgAdaptor) BeginTransaction() (*pg.Tx, error) {
	if this.db == nil {
		return nil, fmt.Errorf("database object invalid")
	}

	tx, err := this.db.Begin()
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (this *PgAdaptor) ExecTransaction(tx *pg.Tx, query string, args ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	smt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.Exec(args...)
	if err != nil {
		return err
	}

	return nil

}

func (this *PgAdaptor) CommitTransaction(tx *pg.Tx) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (this *PgAdaptor) RollbackTransaction(tx *pg.Tx) error {

	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	err := tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}
