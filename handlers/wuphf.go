package handlers

import (
	"errors"
	"net/http"

	"github.com/brandenc40/wuphf.com/common"
	"github.com/brandenc40/wuphf.com/controllers"
	"github.com/brandenc40/wuphf.com/models"
	"github.com/gin-gonic/gin"
)

func bindPostToWuphfParams(c *gin.Context) (*controllers.WuphfParams, error) {

	var smsNumber models.PhoneNumber
	var callNumber models.PhoneNumber
	var err error

	rawSmsNumber := c.PostForm("sms_number")
	rawCallNumber := c.PostForm("call_number")
	email := c.PostForm("to_email")
	message := c.PostForm("message")
	fromName := c.PostForm("from_name")

	if smsNumber, err = models.NewPhoneNumber(rawSmsNumber); err != nil {
		return nil, err
	}
	if callNumber, err = models.NewPhoneNumber(rawCallNumber); err != nil {
		return nil, err
	}
	if err = common.ValidateEmailFormat(email); email != "" && err != nil {
		return nil, err
	}
	if common.ContainsCurseWords(message) || common.ContainsCurseWords(fromName) {
		return nil, errors.New("Avoid using curse words.")
	}

	return &controllers.WuphfParams{
		Message:    message,
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

	if err := h.controllers.SendWuphf(wuphfParams); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.String(http.StatusOK, successMessage)
		return
	}
}
