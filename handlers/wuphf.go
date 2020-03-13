package handlers

import (
	"net/http"

	"github.com/brandenc40/wuphf.com/controllers"
	"github.com/brandenc40/wuphf.com/models"
	"github.com/gin-gonic/gin"
)

func bindPostToWuphfParams(c *gin.Context) (*controllers.WuphfParams, error) {

	smsNumber, _ := models.NewPhoneNumber(c.PostForm("sms_number"))
	callNumber, _ := models.NewPhoneNumber(c.PostForm("call_number"))

	return &controllers.WuphfParams{
		Message:    c.PostForm("message"),
		FromName:   c.PostForm("from_name"),
		SMSNumber:  smsNumber,
		CallNumber: callNumber,
		ToEmail:    c.PostForm("to_email"),
	}, nil
}

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
