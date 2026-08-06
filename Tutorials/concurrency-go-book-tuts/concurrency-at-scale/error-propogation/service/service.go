package service

import (
	"fmt"
	something "practice-error-propogation/err"
	"practice-error-propogation/repository"
)

type ServiceErorr struct {
	Err error
}

func (s ServiceErorr) Error() string {
	if s.Err != nil {
		return "service error: " + s.Err.Error()
	}
	return "service error"
}

func (s ServiceErorr) Unwrap() error {
	return s.Err
}

func DownloadInvoice(id string) ([]byte, error) {
	invoice, err := repository.GetInvoice(id)
	if err != nil {
		return nil, ServiceErorr{
			Err: something.WrapE(err, "service.DownloadService", fmt.Sprintf("invoice %s is currently Unavailable", id)),
		}
	}
	return invoice, nil
}
