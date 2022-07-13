package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Event struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	DateFromRaw string    `json:"date_from"`
	DateToRaw   string    `json:"date_to"`
	DateFrom    time.Time `json:"-"`
	DateTo      time.Time `json:"-"`
}

func (data *Event) ParseJSON(req *http.Request) error {
	err := json.NewDecoder(req.Body).Decode(&data)

	if data.Title == "" {
		return errors.New("empty title")
	}

	if data.DateFrom, err = time.Parse("2006-01-02", data.DateFromRaw); err != nil {
		return err
	}

	if data.DateTo, err = time.Parse("2006-01-02", data.DateToRaw); err != nil {
		return err
	}
	if data.DateTo.Before(data.DateFrom) {
		return errors.New("invalid data")
	}
	return nil
}

func (data *Event) GetParse(req *http.Request) error {
	var err error
	if err = req.ParseForm(); err != nil {
		return err
	}

	if data.DateFrom, err = time.Parse("2006-01-02", req.FormValue("date_from")); err != nil {
		return err
	}
	return nil
}
