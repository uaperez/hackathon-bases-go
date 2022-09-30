package main

import (
	"fmt"
	"spirit-airlines/internal/model"
	"spirit-airlines/internal/service"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			println(err)
		}
	}()
	var tickets []model.Ticket
	// Funcion para obtener tickets del archivo csv
	bookingService := service.NewBookings(tickets)

	/* ======================== Crear ======================== */
	value, errorCreate := bookingService.Create(model.Ticket{
		Names:       "Mercho Vil",
		Email:       "ferxxo@universalmusic.com",
		Destination: "Medallo",
		Date:        "13:45",
		Price:       1234,
	})
	if errorCreate != nil {
		panic(errorCreate.Error())
	}
	booking, _ := value.ToJson()
	fmt.Println("Ticket recién creado:", booking)

	/* ======================== Leer ======================== */
	ticket, errorReading := bookingService.Read(140)
	if errorReading != nil {
		panic(errorReading.Error())
	}
	ticketReaded, _ := ticket.ToJson()
	fmt.Println("Ticket leído:", ticketReaded)

	/* ======================== Actualizar ======================== */
	ticketUpdated, errorUpdate := bookingService.Update(98, model.Ticket{
		Names:       "Andy Perez",
		Email:       "sfallonrq@etsy.com",
		Destination: "Colombia",
		Date:        "13:50",
		Price:       550,
	})
	if errorUpdate != nil {
		panic(errorUpdate.Error())
	}
	ticketUpdatedAsJson, _ := ticketUpdated.ToJson()
	fmt.Println("Ticket recién actualizado", ticketUpdatedAsJson)

	/* ======================== Eliminar ======================== */
	ticketDeletedId, errorAtDelete := bookingService.Delete(99)
	if errorAtDelete != nil {
		panic(errorAtDelete.Error())
	}
	fmt.Println("Eliminado el registro con ID:", ticketDeletedId)
}
