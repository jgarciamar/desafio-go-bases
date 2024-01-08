package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	ticketData, err := tickets.GetTicketsFromCSV("tickets.csv")

	if err != nil {
		fmt.Println(err)
	}
	total, err := tickets.GetTotalTickets("Colombia", ticketData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Total Tickets: ", total)

	eveningFlights, err := tickets.GetCountByPeriod("evening", ticketData)
	nightFlights, err := tickets.GetCountByPeriod("night", ticketData)
	morningFlights, err := tickets.GetCountByPeriod("morning", ticketData)
	dawnFlights, err := tickets.GetCountByPeriod("dawn", ticketData)

	fmt.Println(eveningFlights + nightFlights + morningFlights + dawnFlights)

	someDestination := "Colombia"

	averageToColombia, err := tickets.AverageDestination(someDestination, ticketData)
	fmt.Println("Average to Colombia", averageToColombia)
}
