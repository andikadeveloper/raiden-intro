package controllers

import (
	"raidenintro/internal/models"

	"github.com/sev-2/raiden"
)

type OrderController struct {
	raiden.ControllerBase
	Http  string `path:"/orders" type:"rest"`
	Model models.Order
}
