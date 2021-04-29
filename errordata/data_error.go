package errordata

import (
	"fmt"
)

type Err struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message string `json:"message"`
}

type ValidateErr struct {
	// code error
	Code int64 `json:"code"`
	// description error
	Message map[string]string `json:"message"`
}

func NewErr(err error) *Err {
	if _, exist := MapErrorCode[err]; exist {
		return &Err{
			Code:    MapErrorCode[err],
			Message: MapDescription[err],
		}
	} else {
		return &Err{
			Code:    401,
			Message: err.Error(),
		}
	}
}

func NewErrByText(err error, text string) *Err {
	return &Err{
		Code:    MapErrorCode[err],
		Message: text,
	}
}

func NewErrWithPara(err error, args ...string) *Err {
	if _, exist := MapErrorCode[err]; exist {
		return &Err{
			Code:    MapErrorCode[err],
			Message: fmt.Sprintf(MapDescription[err], args),
		}
	} else {
		return &Err{
			Code:    401,
			Message: err.Error(),
		}
	}
}
