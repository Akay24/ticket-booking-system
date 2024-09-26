package main

import "fmt"

func main() {
	showName := "Diljits concert" //cant be used for constants
	const concertTickets int = 1000

	var remainingTickets uint = 1000

	fmt.Printf("showName is of type %T concertticket is of type %T and finally the remainingtickets is of type %T \n", showName, concertTickets, remainingTickets)
	fmt.Printf("Hello User \n")
	fmt.Printf("Get your tickets here for %v \n", showName)
	fmt.Printf("Number of tickets left are %v \n", remainingTickets)

	var booking []string

	var Firstname string
	var Lastname string
	var Email string
	var userTicket int

	//ask for user details
	fmt.Println("Enter your firstname")
	fmt.Scan(&Firstname)
	fmt.Println("Enter your lastname")
	fmt.Scan(&Lastname)
	fmt.Println("Enter your email")
	fmt.Scan(&Email)
	fmt.Println("enter the number of tickets you want")
	fmt.Scan(&userTicket)
	remainingTickets = remainingTickets - uint(userTicket)

	booking = append(booking, Firstname+" "+Lastname)

	fmt.Printf("The whole array is : %v \n"+" ", booking)
	fmt.Printf("the first value is %v:\n", booking[0])
	fmt.Printf("Array type is %T:\n", booking[0])
	fmt.Printf("Array length is %v:\n", len(booking[0]))

	fmt.Printf("Thankyou %v %v for booking %v tickets through %v \n", Firstname, Lastname, userTicket, Email)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, showName)

}
