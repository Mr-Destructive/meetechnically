package main

import (
	"fmt"
	"weak"
)

var tickets = make(map[int]*int)

func bookTicket(n int) *int {
	if ptr, ok := tickets[n]; ok {
		fmt.Println("Already booked")
		return ptr
	}
	fmt.Println("Booking new ticket")
	tickets[n] = &n
	return &n
}

var tickets2 = make(map[int]weak.Pointer[int])

func getOrBookTicket(n int) *int {
	if ptr, ok := tickets2[n]; ok {
		if ticket := ptr.Value(); ticket != nil {
			return ticket
		}
		delete(tickets, n)
	}

	ticket := n
	tickets2[n] = weak.Make(&ticket)
	return &ticket
}

func main1() {
	a := bookTicket(5)
	b := bookTicket(5)
	fmt.Println(*a, *b, a == b)
	fmt.Println(tickets)

	c := getOrBookTicket(5)
	d := getOrBookTicket(5)
	fmt.Println(*c, *d, c == d)

	fmt.Println(tickets2)
}
