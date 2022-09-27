package models

type Backpack struct {
	ItemId   int    `json:"itemid"`
	ItemName string `json:"itemname"`
	Price    int    `json:"price"`
}

type Bodyarmor struct {
	ItemId   int    `json:"itemid"`
	ItemName string `json:"itemname"`
	Price    int    `json:"price"`
}

type Helmet struct {
	ItemId   int    `json:"itemid"`
	ItemName string `json:"itemname"`
	Price    int    `json:"price"`
}
