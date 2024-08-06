package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets int = 50
var bookings = []string{}

type BookingRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	UserTickets int    `json:"userTickets"`
}

type BookingResponse struct {
	Message          string `json:"message"`
	RemainingTickets int    `json:"remainingTickets"`
}

func main() {
	http.HandleFunc("/greet", greetUsers)
	http.HandleFunc("/book", bookTicketHandler)
	http.HandleFunc("/bookings", getBookingsHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func greetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": fmt.Sprintf("Welcome to our %v booking application. We have total of %v tickets and %v are still available. Get your tickets to attend.", conferenceName, conferenceTickets, remainingTickets),
	}
	json.NewEncoder(w).Encode(response)
}

func bookTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var bookingReq BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&bookingReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	isValidUsername, isValidEmail, isValidTickets := userInputValid(bookingReq.FirstName, bookingReq.LastName, bookingReq.Email, bookingReq.UserTickets)
	if !isValidUsername || !isValidEmail || !isValidTickets {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	remainingTickets, bookings = bookTicket(bookingReq.UserTickets, bookingReq.FirstName, bookingReq.LastName, bookingReq.Email)

	response := BookingResponse{
		Message:          fmt.Sprintf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v", bookingReq.FirstName, bookingReq.LastName, bookingReq.UserTickets, bookingReq.Email),
		RemainingTickets: remainingTickets,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getBookingsHandler(w http.ResponseWriter, r *http.Request) {
	firstNames := getFirstNames()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(firstNames)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func bookTicket(userTickets int, firstName string, lastName string, email string) (int, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, conferenceName)
	return remainingTickets, bookings
}

func userInputValid(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidUsername := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidUsername, isValidEmail, isValidTickets
}
