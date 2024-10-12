package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
		log.Println("Error while querying customers", err.Error())
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

func (ch CustomerRepositoryDB) ById(customerId string) (Customer, error) {
	findById := "SELECT customer_id , name , date_of_birth , status , city , zip_code FROM customers WHERE customer_id = ?"

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
