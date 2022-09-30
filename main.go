package main

import (
	"fmt"
	"spirit-airlines/internal/model"
	"spirit-airlines/internal/service"
)

func main() {
	var tickets []model.Ticket
	// Funcion para obtener tickets del archivo csv
	bookingService := service.NewBookings(tickets)

	/*ticket, err := bookingService.Update(999, model.Ticket{
		Names:       "Andy Perez",
		Email:       "sfallonrq@etsy.com",
		Destination: "Colombia",
		Date:        "13:50",
		Price:       550,
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(ticket.ToJson())*/

	/*status, err := bookingService.Delete(1001)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Eliminado el registro con ID:", status)*/
	//var filename file.File = file.File{}
	//filename.SetPathname("tickets.csv")
	/*records, err := filename.Read()
	if err != nil {
		panic(err.Error())
	}
	err1 := filename.Write(model.Ticket{
		Id:          1002,
		Names:       "Juancho Perez",
		Email:       "meli@melidata.com",
		Destination: "Los Angeles",
		Date:        "13:45",
		Price:       1234,
	})
	if err1 != nil {
		panic(err1.Error())
	}*/

	value, err := bookingService.Create(model.Ticket{
		Names:       "Mercho Vil",
		Email:       "ferxxo@universalmusic.com",
		Destination: "Medallo",
		Date:        "13:45",
		Price:       1234,
	})
	if err != nil {
		panic(err.Error())
	}
	booking, _ := value.ToJson()
	fmt.Println(booking)

	/*ticket, err := bookingService.Read(140)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(ticket.ToJson())*/

	//fmt.Println(records)
}
