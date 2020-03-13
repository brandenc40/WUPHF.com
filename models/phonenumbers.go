package models

import (
	"fmt"

	"github.com/ttacon/libphonenumber"
)

type PhoneNumber string

const minPhoneNumLength = 4

func NewPhoneNumber(pn string) (PhoneNumber, error) {
	if pn == "" {
		return "", nil
	}
	num, err := libphonenumber.Parse(pn, "US")
	switch {
	case err == libphonenumber.ErrNotANumber:
		return "", fmt.Errorf("Invalid phone number: %s", pn)
	case err == libphonenumber.ErrInvalidCountryCode:
		return "", fmt.Errorf("Invalid country code for number: %s", pn)
	case err != nil:
		return "", err
	}
	if len(num.String()) < minPhoneNumLength {
		return "", fmt.Errorf("Invlid number, must be longer than %d digits", minPhoneNumLength)
	}
	return PhoneNumber(libphonenumber.Format(num, libphonenumber.E164)), nil
}

// Friendly returns a friendly international representation of the phone
// number, for example, "+14105554092" is returned as "+1 410-555-4092". If the
// phone number is not in E.164 format, we try to parse it as a US number. If
// we cannot parse it as a US number, it is returned as is.
func (pn PhoneNumber) Friendly() string {
	num, err := libphonenumber.Parse(string(pn), "US")
	if err != nil {
		return string(pn)
	}
	return libphonenumber.Format(num, libphonenumber.INTERNATIONAL)
}

// Local returns a friendly national representation of the phone number, for
// example, "+14105554092" is returned as "(410) 555-4092". If the phone number
// is not in E.164 format, we try to parse it as a US number. If we cannot
// parse it as a US number, it is returned as is.
func (pn PhoneNumber) Local() string {
	num, err := libphonenumber.Parse(string(pn), "US")
	if err != nil {
		return string(pn)
	}
	return libphonenumber.Format(num, libphonenumber.NATIONAL)
}
