package paymentmethod

import (
	"database/sql"
)

type PaymentMethodRepository struct {
	DB *sql.DB
}

func (r *PaymentMethodRepository) GetAllPaymentMethods() ([]PaymentMethod, error) {
	rows, err := r.DB.Query("SELECT id, name FROM payment_methods")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []PaymentMethod
	for rows.Next() {
		var method PaymentMethod
		if err := rows.Scan(&method.ID, &method.Name); err != nil {
			return nil, err
		}
		methods = append(methods, method)
	}
	return methods, nil
}
