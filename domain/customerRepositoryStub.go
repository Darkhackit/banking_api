package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (r CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return r.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Emmanuel", "Accra", "23 33", "15-07-1994", "active"},
		{"1002", "James", "Koforidua", "2333", "15-07-1994", "active"},
		{"1003", "John", "Kumasi", "2333", "15-07-1994", "active"},
		{"1004", "Faustina", "Takoradi", "2333", "15-07-1994", "active"},
	}

	return CustomerRepositoryStub{customers}
}
