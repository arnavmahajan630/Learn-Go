package repository

import (
	"fmt"
	something "practice-error-propogation/err"
	"practice-error-propogation/fs"
)

type RepositoryError struct {
	Err error
}

func (r RepositoryError) Error() string {
	if r.Err != nil {
		return "repository error: " + r.Err.Error()
	}
	return "repository error"
}

func (r RepositoryError) Unwrap() error {
	return r.Err
}


func GetInvoice(id string) ([]byte, error) {
	path := fmt.Sprintf("invoices%s.pdf", id)
	data , err := fs.ReadInvoice(path)

	if err != nil {
		return nil, RepositoryError{
			Err: something.WrapE(err, "repository.GetInvoice", "invoice could not be loaded"),
		}
	}

	return data, nil

}