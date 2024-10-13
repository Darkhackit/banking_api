package domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github/Darkhackit/banking_api/errs"
	"github/Darkhackit/banking_api/logger"
	"log"
	"time"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (cr CustomerRepositoryDB) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)
	findAllSql := "SELECT customer_id , name , date_of_birth , status , city , zip_code FROM customers"
	//rows, err := cr.client.Query(findAllSql)
	err := cr.client.Select(&customers, findAllSql)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("customer repository not found" + err.Error())
			return nil, errors.New("customer repository not found")
		}
		log.Println("Error while querying customer", err.Error())
		return nil, err
	}

	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		logger.Error("customer repository not found" + err.Error())
	//		return nil, err
	//	} else {
	//		logger.Error("An error occurred" + err.Error())
	//		return nil, err
	//	}
	//}
	//for rows.Next() {
	//	var customer Customer
	//	err := rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.Status, &customer.City, &customer.ZipCode)
	//	if err != nil {
	//		logger.Error("Error while scanning customers" + err.Error())
	//		return nil, err
	//	}
	//	customers = append(customers, customer)
	//}
	//defer func(rows *sql.Rows) {
	//	err := rows.Close()
	//	if err != nil {
	//		logger.Error("Error while closing customers" + err.Error())
	//		return
	//	}
	//}(rows)
	return customers, nil
}

func (cr CustomerRepositoryDB) ById(customerId string) (*Customer, *errs.AppErrors) {
	var customer Customer
	findById := "SELECT customer_id , name , date_of_birth , status , city , zip_code FROM customers WHERE customer_id = ?"
	err := cr.client.Get(&customer, findById, customerId)
	//row := cr.client.QueryRow(findById, customerId)
	//err := row.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.Status, &customer.City, &customer.ZipCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("customer repository not found" + err.Error())
			return nil, errs.NewNotFoundError("customer repository not found")
		} else {
			logger.Error("Error while selecting customers" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected Database Error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{client}
}
