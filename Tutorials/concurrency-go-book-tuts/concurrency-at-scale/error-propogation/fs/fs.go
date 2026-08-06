package fs

import (
	"os"
	something "practice-error-propogation/err"
)

type StorageError struct {
	Err error
}

func (s StorageError) Error() string {
	if s.Err != nil  {
		return "storage error: " + s.Err.Error()
	}
	return "storage error"
}

func (s StorageError) Unwrap() error {
	return s.Err
}


func ReadInvoice(path string )([]byte, error) {
	data ,err := os.ReadFile(path)
	if err != nil {
		return nil, StorageError{
			Err: something.WrapE(err, "storage.ReadInvoice", err.Error()),
		}
	}
	return data, nil 
}

