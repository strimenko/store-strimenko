package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/strimenko/store-strimenko/models"
)

type BackpackPostgres struct {
	db *sqlx.DB
}

func NewBackpackPostgres(db *sqlx.DB) *BackpackPostgres {
	return &BackpackPostgres{db: db}
}

func (r *BackpackPostgres) Create(models models.Backpack) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int
	createBackpackQuery := fmt.Sprintf("INSERT INTO %s (itemname, price) VALUES ($1, $2) RETURNING itemid", backpacksTable)
	row := tx.QueryRow(createBackpackQuery, models.ItemName, models.Price)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BackpackPostgres) GetAll() ([]models.Backpack, error) {
	var lists []models.Backpack

	query := fmt.Sprintf("SELECT * FROM %s", backpacksTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *BackpackPostgres) GetById(itemId int) (models.Backpack, error) {
	var lists models.Backpack

	query := fmt.Sprintf("SELECT * FROM %s WHERE itemid = $1", backpacksTable)
	err := r.db.Get(&lists, query, itemId)

	return lists, err
}
