package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Ticket struct {
	ID          int
	Name        string
	Email       string
	Destination string
	FlightTime  time.Time
	Price       int
}

func ParseArrToTicket(ticketStr []string) (ticket Ticket, err error) {

	ID, err := strconv.Atoi(ticketStr[0])

	if err != nil {
		return ticket, err
	}

	Price, err := strconv.Atoi(ticketStr[5])

	if err != nil {
		return ticket, err
	}

	flightTime, err := ParseFlightTime(ticketStr[4])

	if err != nil {
		return ticket, err
	}

	ticket = Ticket{
		ID:          ID,
		Name:        ticketStr[1],
		Email:       ticketStr[2],
		Destination: ticketStr[3],
		FlightTime:  flightTime,
		Price:       Price,
	}

	return ticket, nil

}

var TimeParsingError error = errors.New("Error parsing time")

var (
	InvalidDestionationErr error = errors.New("Invalid destination")
	InvalidTicketDataErr         = errors.New("Invalid data!")
)

func ParseFlightTime(timeString string) (parsedTime time.Time, err error) {

	layout := "15:04"

	parsedTime, err = time.Parse(layout, timeString)

	if err != nil {
		return time.Time{}, TimeParsingError
	}

	return parsedTime, nil

}

func GetTicketsFromCSV(path string) (tickets []Ticket, err error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	ticketList, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading")
		return
	}

	for _, ticketArr := range ticketList {

		ticket, err := ParseArrToTicket(ticketArr)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)

	}
	return tickets, nil
}

func GetTotalTickets(destination string, tickets []Ticket) (int, error) {
	if destination == "" {
		return 0, InvalidDestionationErr
	}

	var flightsCounter int

	for _, ticket := range tickets {
		if ticket.Destination == destination {
			flightsCounter += 1
		}
	}
	return flightsCounter, nil
}

type TimePeriod struct {
	start time.Time
	end   time.Time
}

func NewTimePeriod(startTimeStr, endTimeStr string) (TimePeriod, error) {
	startTime, err := ParseFlightTime(startTimeStr)
	if err != nil {
		return TimePeriod{}, err
	}

	endTime, err := ParseFlightTime(endTimeStr)
	if err != nil {
		return TimePeriod{}, err
	}

	return TimePeriod{start: startTime, end: endTime}, nil
}

func GetCountByPeriod(timePeriod string, tickets []Ticket) (int, error) {

	var ticketCounter int

	dawn, err := NewTimePeriod("0:00", "7:00")
	if err != nil {
		fmt.Println("Error creating dawn TimePeriod:", err)
	}

	morning, err := NewTimePeriod("7:00", "13:00")
	if err != nil {
		fmt.Println("Error creating morning TimePeriod:", err)
	}

	evening, err := NewTimePeriod("13:00", "20:00")
	if err != nil {
		fmt.Println("Error creating evening TimePeriod:", err)
	}

	night, err := NewTimePeriod("20:00", "00:00")
	if err != nil {
		fmt.Println("Error creating night TimePeriod:", err)
	}

	timePeriodMap := map[string]TimePeriod{
		"dawn":    dawn,
		"morning": morning,
		"evening": evening,
		"night":   night,
	}

	period, ok := timePeriodMap[timePeriod]
	if !ok {
		fmt.Println("Invalid timePeriod")
		return 0, errors.New("Invalid timePeriod")
	}

	for _, ticket := range tickets {
		if period.end.Before(period.start) {
			// Handle the case where the period spans midnight
			if ticket.FlightTime.After(period.start) || ticket.FlightTime.Before(period.end) {
				ticketCounter++
			}
		} else {
			// Normal case
			if ticket.FlightTime.After(period.start) && ticket.FlightTime.Before(period.end) {
				ticketCounter++
			}
		}
	}
	return ticketCounter, nil

}

func AverageDestination(destination string, tickets []Ticket) (float64, error) {

	ticketTotalAmount := len(tickets)
	ticketToDestinationAmount := 0

	if destination == "" {
		return 0.00, nil
	}

	if ticketTotalAmount == 0 {
		return 0.00, nil
	}

	for _, ticket := range tickets {
		if ticket.Destination == destination {
			ticketToDestinationAmount++
		}
	}

	average := float64(ticketToDestinationAmount) / float64(ticketTotalAmount)
	return average, nil
}
