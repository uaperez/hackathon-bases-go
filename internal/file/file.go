package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"spirit-airlines/internal/model"
	"strconv"
)

var (
	PATH_FILENAME = "tickets.csv"
)

type File struct {
	path string
}

func (f *File) SetPathname(name string) {
	f.path = name
}

func (f *File) Read() ([]model.Ticket, error) {
	file, err := os.Open(PATH_FILENAME)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	tickets := make([]model.Ticket, 0)

	for _, record := range records {
		ticket, err := recordToTicket(record)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (f *File) Write(ticket model.Ticket) error {
	file, err := os.OpenFile(PATH_FILENAME, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write(ticket.ToArray())
	if err := writer.Error(); err != nil {
		return err
	}
	writer.Flush() // limpiar el buffer del writer y permitir que pueda terminar de escribirse (1.5 horas buscando esto, fuck)
	return nil
}

func (f *File) GetById(id int) (model.Ticket, error) {
	tickets, err := f.Read()
	if err != nil {
		return model.Ticket{}, err
	}
	for _, ticket := range tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return model.Ticket{}, errors.New(fmt.Sprintf("El ticket con el ID %v no fue encontrado", id))
}

func (f *File) Update(id int, ticket model.Ticket) (model.Ticket, error) {
	tickets, err := f.Read()
	if err != nil {
		return model.Ticket{}, err
	}
	var updatedTicket model.Ticket
	var idx int

	for index, currentTicket := range tickets {
		if currentTicket.Id == id {
			idx = index
			break
		}
	}

	if idx == 0 {
		return model.Ticket{}, errors.New("No fue posible actualizar lo que se deseaba :/")
	}

	tickets[idx] = model.Ticket{
		Id:          id,
		Names:       ticket.Names,
		Email:       ticket.Email,
		Destination: ticket.Destination,
		Date:        ticket.Date,
		Price:       ticket.Price,
	}
	updatedTicket = tickets[idx]

	err = rewrite(tickets, PATH_FILENAME)

	if err != nil {
		return model.Ticket{}, err
	}
	return updatedTicket, nil
}

func (f *File) Delete(id int) (bool, error) {
	tickets, _ := f.Read()
	var newTickets []model.Ticket

	for index, ticket := range tickets {
		if ticket.Id == id {
			newTickets = removeIndex(tickets, index)
			break
		}
	}

	err := rewrite(newTickets, PATH_FILENAME)

	if err != nil {
		return false, err
	}

	return true, nil
}

func recordToTicket(record []string) (model.Ticket, error) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		return model.Ticket{}, errors.New("Ha ocurrido un error al convertir el ID a entero")
	}
	price, err1 := strconv.Atoi(record[5])
	if err1 != nil {
		return model.Ticket{}, errors.New("Ha ocurrido un error al convertir el precio a entero")
	}
	return model.Ticket{
		Id:          id,
		Names:       record[1],
		Email:       record[2],
		Destination: record[3],
		Date:        record[4],
		Price:       price,
	}, nil
}

func removeIndex(tickets []model.Ticket, index int) []model.Ticket {
	newArray := make([]model.Ticket, 0)
	newArray = append(newArray, tickets[:index]...)
	return append(newArray, tickets[index+1:]...)
}

func rewrite(data []model.Ticket, pathname string) error {
	file, err := os.Create(pathname)
	if err != nil {
		return err
	}
	defer file.Close()

	var dataAsArrayString [][]string

	for _, v := range data {
		dataAsArrayString = append(dataAsArrayString, v.ToArray())
	}

	writer := csv.NewWriter(file)
	writer.WriteAll(dataAsArrayString)

	if err := writer.Error(); err != nil {
		return err
	}

	writer.Flush()
	return nil
}
