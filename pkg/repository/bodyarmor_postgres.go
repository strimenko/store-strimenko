package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/strimenko/store-strimenko/models"
)

type BodyarmorPostgres struct {
	db *sqlx.DB
}

func NewBodyarmorPostgres(db *sqlx.DB) *BodyarmorPostgres {
	return &BodyarmorPostgres{db: db}
}

func (r *BodyarmorPostgres) Create(models models.Bodyarmor) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var id int
	createBodyarmorQuery := fmt.Sprintf("INSERT INTO %s (itemname, price) VALUES ($1, $2) RETURNING itemid", bodyarmorsTable)
	row := tx.QueryRow(createBodyarmorQuery, models.ItemName, models.Price)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BodyarmorPostgres) GetAll() ([]models.Bodyarmor, error) {
	var lists []models.Bodyarmor

	query := fmt.Sprintf("SELECT * FROM %s", bodyarmorsTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *BodyarmorPostgres) GetById(itemId int) (models.Bodyarmor, error) {
	var lists models.Bodyarmor

	query := fmt.Sprintf("SELECT * FROM %s WHERE itemid = $1", bodyarmorsTable)
	err := r.db.Get(&lists, query, itemId)

	return lists, err
}
