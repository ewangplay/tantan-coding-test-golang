package main

import (
	"fmt"
	"gopkg.in/pg.v4"
)

type PgAdaptor struct {
	db *pg.DB
}

func NewPgAdaptor(addr, user, pass, dbname string) (*PgAdaptor, error) {
	dbAdaptor := &PgAdaptor{}
	dbAdaptor.db = pg.Connect(&pg.Options{
        Addr:     addr,
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

func (this *PgAdaptor) Query(model, query interface{}, params ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	_, err := this.db.Query(model, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (this *PgAdaptor) QueryOne(model, query interface{}, params ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	_, err := this.db.QueryOne(model, query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (this *PgAdaptor) Exec(query interface{}, params ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	_, err := this.db.Exec(query, params...)
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

func (this *PgAdaptor) ExecTransaction(tx *pg.Tx, query string, params ...interface{}) error {
	if this.db == nil {
		return fmt.Errorf("database object invalid")
	}

	smt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer smt.Close()

	_, err = smt.Exec(params...)
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
