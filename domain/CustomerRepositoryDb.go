package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	//rows, err := d.client.Query(findAllSql)
	customers := make([]Customer, 0)
	err := d.client.Select(&customers, findAllSql)
	if err != nil {
		logger.Error("unexpected database error: " + err.Error())
		return nil, errs.NewUnexceptedError("unexpected database error")
	}

	//err = sqlx.StructScan(rows, &customers)

	/*if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scaning customer result: " + sql.ErrNoRows.Error())
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scaning customer result: " + err.Error())
			return nil, errs.NewUnexceptedError("unexpected database error")
		}
	}*/
	/*for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error while scaning customer result: " + sql.ErrNoRows.Error())
				return nil, errs.NewNotFoundError("customer not found")
			} else {
				logger.Error("Error while scaning customer result: " + err.Error())
				return nil, errs.NewUnexceptedError("unexpected database error")
			}
		}

		customers = append(customers, c)
	}*/

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"

	//row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scaning customer result: " + sql.ErrNoRows.Error())
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scaning customer result: " + err.Error())
			return nil, errs.NewUnexceptedError("unexpected database error")
		}

	}

	/*if c == Customer{} {
		logger.Error("No results for customerSql")
		return nil, errs.NewNotFoundError("No customers found with this id: " + id)
	}*/

	return &c, nil

}

func (d CustomerRepositoryDb) ByStatus(status int) ([]Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status=?"

	//rows, err := d.client.Query(customerSql, status)
	customers := make([]Customer, 0)
	err := d.client.Select(&customers, customerSql, status)
	if err != nil {
		logger.Error("Error while scaning customer result qw: " + err.Error())
		return nil, errs.NewUnexceptedError("unexpected database error")
	}

	if cap(customers) == 0 {
		logger.Error("No results for customerSql")
		return nil, errs.NewNotFoundError("No customers found with this stats")
	}

	return customers, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}
