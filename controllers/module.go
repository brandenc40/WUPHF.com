package controllers

import "github.com/brandenc40/wuphf.com/gateways"

type Controllers struct {
	*gateways.Gateway
}

func New() *Controllers {
	return &Controllers{
		gateways.New(),
	}
}
