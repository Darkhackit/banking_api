package domain

import "github/Darkhackit/banking_api/errs"

type Customer struct {
	Id          string `db:"customer_id" json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `db:"zip_code" json:"zip_code"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	Status      string `db:"status" json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(id string) (*Customer, *errs.AppErrors)
}
