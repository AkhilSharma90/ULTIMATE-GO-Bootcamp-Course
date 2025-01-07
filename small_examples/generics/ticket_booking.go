package main

import (
	"fmt"
)

// Generic Booking struct
type Booking[T any] struct {
	EventID string
	UserID  string
	Ticket  T
	Booked  bool
}

// MovieTicket struct
type MovieTicket struct {
	SeatNumber string
	ShowTime   string
}

// ConcertTicket struct
type ConcertTicket struct {
	Section string
	Row     int
	Seat    int
}

// SportsTicket struct
type SportsTicket struct {
	Zone   string
	Row    int
	Seat   int
	GameID string
}

// BookingManager manages bookings for various event types
type BookingManager[T any] struct {
	bookings []Booking[T]
}

// AddBooking adds a new booking
func (bm *BookingManager[T]) AddBooking(eventID string, userID string, ticket T) {
	booking := Booking[T]{
		EventID: eventID,
		UserID:  userID,
		Ticket:  ticket,
		Booked:  true,
	}
	bm.bookings = append(bm.bookings, booking)
	fmt.Printf("Booking added for user %s: %+v\n", userID, booking)
}

// ListBookings lists all bookings
func (bm *BookingManager[T]) ListBookings() {
	fmt.Println("\nListing All Bookings:")
	for i, booking := range bm.bookings {
		fmt.Printf("%d. EventID: %s, UserID: %s, Ticket: %+v, Booked: %t\n", i+1, booking.EventID, booking.UserID, booking.Ticket, booking.Booked)
	}
}

// CancelBooking cancels a booking by index
func (bm *BookingManager[T]) CancelBooking(index int) {
	if index < 0 || index >= len(bm.bookings) {
		fmt.Println("Invalid booking index")
		return
	}
	bm.bookings[index].Booked = false
	fmt.Printf("Booking canceled: %+v\n", bm.bookings[index])
}
func main() {
	// Movie ticket booking manager
	movieManager := BookingManager[MovieTicket]{}
	movieManager.AddBooking("MOVIE123", "USER001", MovieTicket{SeatNumber: "A10", ShowTime: "7:00 PM"})
	movieManager.AddBooking("MOVIE123", "USER002", MovieTicket{SeatNumber: "B15", ShowTime: "7:00 PM"})
	// Concert ticket booking manager
	concertManager := BookingManager[ConcertTicket]{}
	concertManager.AddBooking("CONCERT456", "USER003", ConcertTicket{Section: "VIP", Row: 1, Seat: 12})
	concertManager.AddBooking("CONCERT456", "USER004", ConcertTicket{Section: "General", Row: 20, Seat: 45})
	// Sports ticket booking manager
	sportsManager := BookingManager[SportsTicket]{}
	sportsManager.AddBooking("GAME789", "USER005", SportsTicket{Zone: "North", Row: 10, Seat: 25, GameID: "SOCCER2024"})
	sportsManager.AddBooking("GAME789", "USER006", SportsTicket{Zone: "South", Row: 15, Seat: 30, GameID: "SOCCER2024"})
	// List all bookings
	movieManager.ListBookings()
	concertManager.ListBookings()
	sportsManager.ListBookings()
	// Cancel a booking
	movieManager.CancelBooking(0)
	movieManager.ListBookings()
}
