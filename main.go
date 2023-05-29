package main

import (
	"Booking-app/helper"
	"fmt"
	"strconv"
)

var confrenceName string = "Go Confrence"
	// confrenceName := "Go confrence" another way of defining a variable
const confrenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberofTickets uint
}

func main(){


	greetUsers()


	for {
		firstName,lastName,email,userTickets := getUserInput()
		
		isValidName, isValidEmail,isValidTickets := helper.ValidataUserInput(firstName, lastName, email, uint(userTickets),remainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			
			bookTicket(uint(userTickets),firstName,lastName,email)
			var firstNames = getfirstName()
			fmt.Printf("The first name of bookings are: %v\n", firstNames)
			
			if remainingTickets == 0 {
				// end program
				fmt.Println("our confrence is booked out comeback next year")
				break
			}
		} else{
			if !isValidName{
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address entered doesn't contain @ sign")
			}
			if !isValidTickets{
				fmt.Println("No. of tictets entered is invalid")
			}
		}
		
	}

	
	
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking confrence\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are stil available\n", confrenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getfirstName() []string {
		firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames,booking["firstName"])
			}
			return firstNames
}


func getUserInput() (string,string,string,int) {
	    var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email name:")
		fmt.Scan(&email)

		fmt.Println("Enter no. of tickets")
		fmt.Scan(&userTickets)

		return firstName,lastName,email,int(userTickets)

}

func bookTicket(userTickets uint, firstName string, lastName string,email string){
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	var userData = make(map[string]string)
	userData["firstname"] = firstName
	userData["lastname"] = lastName
	userData["email"] = email
	userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets),10)

	bookings = append(bookings,userData)
	fmt.Printf("List of bookings is %v\n",bookings)	

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets,confrenceName)

}