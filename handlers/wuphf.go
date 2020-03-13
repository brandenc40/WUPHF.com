package handlers

import (
	"errors"
	"net/http"

	"github.com/brandenc40/wuphf.com/controllers"
	"github.com/brandenc40/wuphf.com/models"
	"github.com/gin-gonic/gin"
)

const minPhoneNumLength = 4

func bindPostToWuphfParams(c *gin.Context) (*controllers.WuphfParams, error) {

	var smsNumber models.PhoneNumber
	var callNumber models.PhoneNumber
	var err error

	rawSmsNumber := c.PostForm("sms_number")
	rawCallNumber := c.PostForm("call_number")
	if !isValidPhoneNumber(rawSmsNumber) || !isValidPhoneNumber(rawCallNumber) {
		return nil, errors.New("Invalid phone number.")
	}
	if smsNumber, err = models.NewPhoneNumber(rawSmsNumber); err != nil {
		return nil, err
	}
	if callNumber, err = models.NewPhoneNumber(rawCallNumber); err != nil {
		return nil, err
	}

	return &controllers.WuphfParams{
		Message:    c.PostForm("message"),
		FromName:   c.PostForm("from_name"),
		SMSNumber:  smsNumber,
		CallNumber: callNumber,
		ToEmail:    c.PostForm("to_email"),
	}, nil
}

// WUPHF -
func (h *Handlers) WUPHF(c *gin.Context) {

	var wuphfParams *controllers.WuphfParams
	var err error

	if wuphfParams, err = bindPostToWuphfParams(c); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.SendWuphf(wuphfParams); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.String(http.StatusOK, successMessage)
		return
	}
}

// Avoid usage of numbers like 911
func isValidPhoneNumber(rawNumber string) bool {
	if rawNumber != "" && len(rawNumber) < minPhoneNumLength {
		return false
	}
	return true
}
