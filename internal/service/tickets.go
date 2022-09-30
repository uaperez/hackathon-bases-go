package service

import (
	"spirit-airlines/internal/file"
	"spirit-airlines/internal/model"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t model.Ticket) (model.Ticket, error)
	// Read read a Ticket by id
	Read(id int) (model.Ticket, error)
	// Update update values of a Ticket
	Update(id int, t model.Ticket) (model.Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	// Get all tickets
	GetAll() []model.Ticket
}

type bookings struct {
	Tickets    []model.Ticket
	Repository file.File
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []model.Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t model.Ticket) (model.Ticket, error) {
	t.Id = generateNextId(b.Repository.Read())
	err := b.Repository.Write(t)
	if err != nil {
		return model.Ticket{}, err
	}
	return t, nil
}

func (b *bookings) Read(id int) (model.Ticket, error) {
	ticket, err := b.Repository.GetById(id)
	if err != nil {
		return model.Ticket{}, err
	}
	return ticket, nil
}

func (b *bookings) Update(id int, t model.Ticket) (model.Ticket, error) {
	_, err := b.Repository.GetById(id)
	if err != nil {
		return model.Ticket{}, err
	}
	ticket, err := b.Repository.Update(id, t)
	if err != nil {
		return model.Ticket{}, err
	}
	return ticket, nil
}

func (b *bookings) Delete(id int) (int, error) {
	_, err := b.Repository.GetById(id)
	if err != nil {
		return 0, err
	}
	_, err1 := b.Repository.Delete(id)

	if err1 != nil {
		return 0, err1
	}
	return id, nil
}

func (b *bookings) GetAll() []model.Ticket {
	return b.Tickets
}

func generateNextId(tickets []model.Ticket, err error) int {
	totalTickets := len(tickets)
	if totalTickets > 0 {
		lastTicket := tickets[totalTickets-1]
		return lastTicket.Id + 1
	}
	return 1
}
