package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/strimenko/store-strimenko/models"
)

type HelmetPostgres struct {
	db *sqlx.DB
}

func NewHelmetPostgres(db *sqlx.DB) *HelmetPostgres {
	return &HelmetPostgres{db: db}
}

func (r *HelmetPostgres) Create(models models.Helmet) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int
	createHelmetQuery := fmt.Sprintf("INSERT INTO %s (itemname, price) VALUES ($1, $2) RETURNING itemid", helmetsTable)
	row := tx.QueryRow(createHelmetQuery, models.ItemName, models.Price)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *HelmetPostgres) GetAll() ([]models.Helmet, error) {
	var lists []models.Helmet

	query := fmt.Sprintf("SELECT * FROM %s", helmetsTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *HelmetPostgres) GetById(itemId int) (models.Helmet, error) {
	var lists models.Helmet

	query := fmt.Sprintf("SELECT * FROM %s WHERE itemid = $1", helmetsTable)
	err := r.db.Get(&lists, query, itemId)

	return lists, err
}
