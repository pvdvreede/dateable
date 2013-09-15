package controllers

import "github.com/robfig/revel"
import "github.com/pvdvreede/dateable/app/models"

type Date struct {
	*revel.Controller
}

func (c Date) Index(dateParam string) revel.Result {
	date := models.NewDate(dateParam)
	return c.RenderJson(date)
}

func (c Date) Between(from, to string) revel.Result {
	dates := models.NewDates(from, to)
	return c.RenderJson(dates)
}
