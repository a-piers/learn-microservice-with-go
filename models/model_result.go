package models

import "lib/utils"

type Result struct {
	Success          bool        `json:"success"`
	ErrorCode        utils.Error `json:"error_code"`
	ErrorDescription string      `json:"error_description"`
	ErrorException   string      `json:"error_exception"`
}
