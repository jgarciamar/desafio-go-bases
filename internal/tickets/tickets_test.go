package tickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// We use testify for all this testings

func TestGetTotalTickets(t *testing.T) {

	// Test case 1: Valid destination, multiple matching tickets
	// This is temporary and need to be asked to the tutors about how to do it better.

	t1, _ := ParseFlightTime("5:56")
	t2, _ := ParseFlightTime("9:44")
	t3, _ := ParseFlightTime("12:30")
	t4, _ := ParseFlightTime("15:15")

	tickets1 := []Ticket{
		{ID: 1, Name: "John", Email: "john@example.com", Destination: "Colombia", FlightTime: t1, Price: 500},
		{ID: 2, Name: "Alice", Email: "alice@example.com", Destination: "Colombia", FlightTime: t2, Price: 600},
		{ID: 3, Name: "Bob", Email: "bob@example.com", Destination: "Peru", FlightTime: t3, Price: 700},
		{ID: 4, Name: "Eve", Email: "eve@example.com", Destination: "Colombia", FlightTime: t4, Price: 800},
	}

	total, err := GetTotalTickets("Colombia", tickets1)
	assert.NoError(t, err, "Unnexpected error")

	expectedTotal := 3
	assert.Equal(t, expectedTotal, total, "Expected total")

}

func TestGetCountByPeriod(t *testing.T) {

	t1, _ := ParseFlightTime("7:01")
	t2, _ := ParseFlightTime("8:02")
	t3, _ := ParseFlightTime("12:30")
	t4, _ := ParseFlightTime("15:15")

	tickets := []Ticket{
		{ID: 1, Name: "John", Email: "john@example.com", Destination: "Colombia", FlightTime: t1, Price: 500},
		{ID: 2, Name: "Alice", Email: "alice@example.com", Destination: "Colombia", FlightTime: t2, Price: 600},
		{ID: 3, Name: "Bob", Email: "bob@example.com", Destination: "Peru", FlightTime: t3, Price: 700},
		{ID: 4, Name: "Eve", Email: "eve@example.com", Destination: "Colombia", FlightTime: t4, Price: 800},
	}

	expectedTotalFlights := len(tickets)

	expecterMorningFlights := 3

	dawnFlights, err := GetCountByPeriod("dawn", tickets)
	morningFlights, err := GetCountByPeriod("morning", tickets)
	eveningFlights, err := GetCountByPeriod("evening", tickets)
	nightFlights, err := GetCountByPeriod("night", tickets)

	if err != nil {
		t.Errorf("There´s an error in the computation of flights for")
	}

	assert.Equal(t, dawnFlights+morningFlights+eveningFlights+nightFlights, expectedTotalFlights, "The sum of flights by periods is not equal to the total flights")

	assert.Equal(t, expecterMorningFlights, morningFlights, tickets, "The amount of flicts is not the expected for the morning case.")
}

func TestAverageDestination(t *testing.T) {

	t1, _ := ParseFlightTime("7:01")
	t2, _ := ParseFlightTime("8:02")
	t3, _ := ParseFlightTime("12:30")
	t4, _ := ParseFlightTime("15:15")

	tickets := []Ticket{
		{ID: 1, Name: "John", Email: "john@example.com", Destination: "Colombia", FlightTime: t1, Price: 500},
		{ID: 2, Name: "Alice", Email: "alice@example.com", Destination: "Colombia", FlightTime: t2, Price: 600},
		{ID: 3, Name: "Bob", Email: "bob@example.com", Destination: "Peru", FlightTime: t3, Price: 700},
		{ID: 4, Name: "Eve", Email: "eve@example.com", Destination: "Colombia", FlightTime: t4, Price: 800},
	}

	expectedPercentageToColombia := float64(3) / float64(4)
	PercentageToColombia, err := AverageDestination("Colombia", tickets)

	if err != nil {
		t.Errorf("Found error %e", err)
	}

	assert.Equal(t, expectedPercentageToColombia, PercentageToColombia, "Error in AverageDestination test ")
}
