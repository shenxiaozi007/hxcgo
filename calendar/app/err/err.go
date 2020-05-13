package err

import "errors"

var (
	ERequest   = errors.New("err_request")
	ESign      = errors.New("err_sign")
	EParam     = errors.New("err_param")
	EProductID = errors.New("err_product_id")
	ECode      = errors.New("err_code")
	ENotFound  = errors.New("not found")
	EResponse  = errors.New("err_response")
)
