package err

import "errors"

var (
	EParam                        = errors.New("param error")
	EParamInvalidID               = errors.New("param error: invalid id")
	EParamInvalidMobile           = errors.New("param error: invalid mobile")
	EParamInvalidEmail            = errors.New("param error: invalid email")
	EParamInvalidName             = errors.New("param error: invalid name")
	EParamInvalidRoleID           = errors.New("param error: invalid roleID")
	EParamInvalidGroupID          = errors.New("param error: invalid groupID")
	EParamInvalidPassword         = errors.New("param error: invalid password")
	EParamInvalidPage             = errors.New("param error: invalid page")
	EParamInvalidLimit            = errors.New("param error: invalid limit")
	EParamInvalidPID              = errors.New("param error: invalid pid")
	EDeleteFailed                 = errors.New("delete failed")
	EAssociateRolePrivilegeFailed = errors.New("associate role privilege failed")
	EAssociateAdminRoleFailed     = errors.New("associate admin role failed")
	EDatabase                     = errors.New("database error")
)
