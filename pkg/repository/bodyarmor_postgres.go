package repository

import (
	"fmt"
	"strings"

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

func (r *BodyarmorPostgres) Update(itemId int, input models.UpdateItem) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.ItemName != nil {
		setValues = append(setValues, fmt.Sprintf("itemname=$%d", argId))
		args = append(args, *input.ItemName)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	setQuerry := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE itemid = $%d", bodyarmorsTable, setQuerry, argId)
	args = append(args, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BodyarmorPostgres) Delete(itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE itemid = $1", bodyarmorsTable)
	_, err := r.db.Exec(query, itemId)

	return err
}
