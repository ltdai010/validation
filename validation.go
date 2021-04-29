package main

import (
	"regexp"
	"gitlab.123xe.vn/dailt/validation/errordata"

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
	return errordata.SUCCESS
}

func AddressBTCFormat(address string) error {
	err := validation.Validate(address, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-zA-Z]{34}$`)))
	if err != nil {
		return errordata.VALIDATION_ADDRESS_BTC
	}
	return errordata.SUCCESS
}

func CheckPubkey(keys string) error {
	err := validation.Validate(keys, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-f]{66}$`)))
	if err != nil {
		return errordata.PUBKEY_INVALID_ERROR
	}
	return errordata.SUCCESS
}

func CheckCoin(coin string) error {
	err := validation.Validate(coin, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]$`)))
	if err != nil {
		return errordata.VALIDATION_COIN
	}
	return errordata.SUCCESS
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
	return errordata.SUCCESS
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
	return errordata.SUCCESS
}
