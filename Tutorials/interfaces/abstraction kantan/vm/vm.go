package vm

import "fmt"

type VendingMaching struct{}

func New() *VendingMaching {
	return &VendingMaching{}
}

func (this VendingMaching) GetDrink(money int64, brand string) string {
	return fmt.Sprintf("super hot %s", brand)
}