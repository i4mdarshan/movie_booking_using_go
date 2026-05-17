package booking

import "sync"

// temporary storage data structure

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {

	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) (Booking, error) {

	// if booking exists in memory store then error
	// else add entry in map
	// return Booking, => nil error for current implementation
	
	s.Lock()
	defer s.Unlock()

	seatId := b.SeatID
	_, bookingExists  := s.bookings[seatId]

	if bookingExists {
		var zero Booking
		return zero, ErrSeatAlreadyBooked
	}

	s.bookings[seatId] = b

	return s.bookings[seatId], nil
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {

	s.RLock()
	defer s.RUnlock()

	bookings := []Booking{}

	for _, b := range s.bookings {
		if b.MovieID == movieID {
			bookings = append(bookings, b)
		}
	}

	return bookings
}
