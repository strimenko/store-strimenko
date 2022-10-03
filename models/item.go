package models

import "errors"

type Backpack struct {
	ItemId   int    `json:"itemid" db:"itemid"`
	ItemName string `json:"itemname" db:"itemname"`
	Price    int    `json:"price" db:"price"`
}

type Bodyarmor struct {
	ItemId   int    `json:"itemid" db:"itemid"`
	ItemName string `json:"itemname" db:"itemname"`
	Price    int    `json:"price" db:"price"`
}

type Helmet struct {
	ItemId   int    `json:"itemid" db:"itemid"`
	ItemName string `json:"itemname" db:"itemname"`
	Price    int    `json:"price" db:"price"`
}

type UpdateItem struct {
	ItemName *string `json:"itemname"`
	Price    *int    `json:"price"`
}

func (i UpdateItem) Validate() error {
	if i.ItemName == nil && i.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
