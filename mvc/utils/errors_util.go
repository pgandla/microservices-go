package utils

type AppErr struct {
	Message string		`json:"message"`
	StatusCode int		`json:"status_code"`
	ErrCode string		`json:"err_code"`
}
