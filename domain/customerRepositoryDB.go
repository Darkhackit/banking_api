package domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github/Darkhackit/banking_api/errs"
	"log"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (cr CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSql := "SELECT customer_id , name , date_of_birth , status , city , zip_code FROM customers"
	rows, err := cr.client.Query(findAllSql)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("customer repository not found")
		}
		log.Println("Error while querying customer", err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var customer Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.Status, &customer.City, &customer.ZipCode)
		if err != nil {
			log.Println("Error while scanning customers", err.Error())
			return nil, err
		}
		customers = append(customers, customer)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error while closing customers", err.Error())
			return
		}
	}(rows)
	return customers, nil
}

func (cr CustomerRepositoryDB) ById(customerId string) (*Customer, *errs.AppErrors) {
	findById := "SELECT customer_id , name , date_of_birth , status , city , zip_code FROM customers WHERE customer_id = ?"

	row := cr.client.QueryRow(findById, customerId)
	var customer Customer
	err := row.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.Status, &customer.City, &customer.ZipCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer repository not found")
		} else {
			log.Println("Error while selecting customers", err.Error())
			return nil, errs.NewUnexpectedError("unexpected Database Error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client}
}
