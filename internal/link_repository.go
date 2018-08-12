package main

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// Link database struct
type Link struct {
	Hash      string    `db:"hash"`
	Link      string    `db:"link"`
	CreatedAt time.Time `db:"created_at"`
}

// LinkRepository provides access to database
type LinkRepository struct {
	db     *sqlx.DB
	config Config
}

// SaveLink saves link to database
func (lr *LinkRepository) SaveLink(hash, link string) error {
	fmt.Printf("SaveLink(%s, %s)\n", hash, link)

	tx, err := lr.db.Begin()
	if err != nil {
		return err
	}
	tx.Exec("INSERT INTO LINKS(hash, link) VALUES ($1, $2);", hash, link)
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// FindLink searches link by hash
func (lr *LinkRepository) FindLink(hash string) (string, error) {
	link := Link{}
	err := lr.db.Get(&link, "SELECT link FROM LINKS WHERE hash = $1", hash)
	if err != nil {
		return "", err
	}
	return link.Link, nil
}
