package err

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	EParam                      = errors.New("param error")
	EParamInvalidID             = errors.New("param error: invalid id")
	EParamInvalidOpenID         = errors.New("param error: invalid open id")
	EParamInvalidName           = errors.New("param error: invalid name")
	EDatabase                   = errors.New("database error")
	ENotFound                   = errors.New("not found")
	EParamInvalidPage           = errors.New("param error: invalid page")
	EParamInvalidLimit          = errors.New("param error: invalid limit")
	EDeleteFailed               = errors.New("delete failed")
	EParamInvalidCategoryID     = errors.New("param error: invalid category id")
	EParamInvalidTitle          = errors.New("param error: invalid title")
	EParamInvalidWholesalePrice = errors.New("param error: invalid wholesale price")
	EParamInvalidImage          = errors.New("param error: invalid image")
	EParamInvalidProductID      = errors.New("param error: invalid product id")
)

func DBError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return ENotFound
	}

	return EDatabase
}
