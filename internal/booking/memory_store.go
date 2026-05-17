package booking

// temporary storage data structure

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {

	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) (Booking, error) {

	// if booking exists in memory store then error
	// else add entry in map
	// return Booking, => nil error for current implementation

	seatId := b.SeatID
	_, bookingExists  := s.bookings[seatId]

	if bookingExists {
		var zero Booking
		return zero, ErrSeatAlreadyBooked
	}

	s.bookings[seatId] = b

	return s.bookings[seatId], nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {

	bookings := []Booking{}

	for _, b := range s.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}

	return bookings
}
