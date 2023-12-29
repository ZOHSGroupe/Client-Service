package Models

import (
	_ "gorm.io/gorm"
)

type Client struct {
	ID                   string `json:"id" gorm:"primaryKey"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	NationalID           string `json:"nationalId"`
	Email                string `json:"email"`
	BirthDate            string `json:"birth_date"`
	City                 string `json:"city"`
	Nationality          string `json:"nationality"`
	Gender               string `json:"gender"`
	CreateDate           string `json:"createDate" validate:"require,date"`
	LastModificationDate string `json:"lastModificationDate" validate:"require,date"`
	Status               string `json:"status"`
	Address              string `json:"address"`
}

func (c *Client) TableName() string {
	return "client"
}
