package corporate_flight_bookings

func corpFlightBookings(bookings [][]int, n int) []int {
	if len(bookings) == 0 {
		return nil
	}

	answer := make([]int, n)
	for _, b := range bookings {

		answer[b[0]-1] = answer[b[0]-1] + b[2]
		if b[1] < n {
			answer[b[1]] = answer[b[1]] - b[2]
		}
	}

	for i := 1; i < n; i++ {
		answer[i] = answer[i] + answer[i-1]
	}

	return answer
}
