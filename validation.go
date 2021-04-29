package validation

import (
	"github.com/ltdai010/validation/errordata"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CoinsStruct struct {
	Coin     string
	Currency string
}

func (this CoinsStruct) Validate() error {
	err := validation.ValidateStruct(&this,
		validation.Field(&this.Coin, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))),
		validation.Field(&this.Currency, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))))
	if err != nil {
		return errordata.VALIDATION_COIN
	}
	return nil
}

func SEInt64(number, max int64) error {
	if number > max {
		return errordata.BAD_REQUEST
	}
	return nil
}

func SEFloat64(number, max float64) error {
	if number > max {
		return errordata.BAD_REQUEST
	}
	return nil
}

func AddressBTCFormat(address string) error {
	err := validation.Validate(address, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-zA-Z]{34}$`)))
	if err != nil {
		return errordata.VALIDATION_ADDRESS_BTC
	}
	return nil
}

func CheckPubkey(keys string) error {
	err := validation.Validate(keys, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-f]{66}$`)))
	if err != nil {
		return errordata.PUBKEY_INVALID_ERROR
	}
	return nil
}

func CheckCoin(coin string) error {
	err := validation.Validate(coin, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]$`)))
	if err != nil {
		return errordata.VALIDATION_COIN
	}
	return nil
}

type ValidateListWarning struct {
	Pos       int
	Count     int
	StartTime int64
	EndTime   int64
}

func (this *ValidateListWarning) Validate() error {
	err := validation.ValidateStruct(&this, validation.Field(&this.Pos, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))),
		validation.Field(&this.Count, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))),
		validation.Field(&this.StartTime, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))),
		validation.Field(&this.EndTime, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))))
	if err != nil {
		return errordata.BAD_REQUEST
	}
	return nil
}

type ValidateUserWarning struct {
	PubKey    string
	StartTime int64
	EndTime   int64
}

func (this *ValidateUserWarning) Validate() error {
	err := validation.ValidateStruct(&this, validation.Field(&this.PubKey, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9a-f]{66}$`))),
		validation.Field(&this.StartTime, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))),
		validation.Field(&this.EndTime, validation.Required, validation.NotNil, validation.Match(regexp.MustCompile(`^[0-9]$`))))
	if err != nil {
		return errordata.BAD_REQUEST
	}
	return nil
}
