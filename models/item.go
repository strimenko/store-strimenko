package models

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
