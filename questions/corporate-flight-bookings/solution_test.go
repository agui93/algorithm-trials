package corporate_flight_bookings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		bookings := [][]int{
			{1, 2, 10},
			{2, 3, 20},
			{2, 5, 25},
		}

		answers := corpFlightBookings(bookings, 5)
		assert.Equal(t, []int{10, 55, 45, 25, 25}, answers)
	})

}
