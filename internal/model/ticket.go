package model

import (
	"encoding/json"
	"strconv"
)

type Ticket struct {
	Id          int    `json:"id"`
	Names       string `json:"nombre"`
	Email       string `json:"correo"`
	Destination string `json:"destino"`
	Date        string `json:"fecha"`
	Price       int    `json:"precio"`
}

func (t *Ticket) ToJson() (string, error) {
	ticket, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(ticket), nil
}

func (t *Ticket) ToArray() []string {
	return []string{strconv.Itoa(t.Id), t.Names, t.Email, t.Destination, t.Date, strconv.Itoa(t.Price)}
}
