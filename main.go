package main
import (
	"fmt"
	"strings"
)

func main(){

	var confrenceName string = "Go Confrence"
	// confrenceName := "Go confrence" another way of defining a variable
	const confrenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have total of %v tickets and %v are stil available\n", confrenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email name:")
		fmt.Scan(&email)

		fmt.Println("Enter no. of tickets")
		fmt.Scan(&userTickets)

		var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(email, "@")
		var isValidTickets = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings,firstName + " "+ lastName)
			
			
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets,confrenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames,names[0])
			}
			fmt.Printf("The first name of bookings are: %v\n",firstNames)
			
			
			if remainingTickets == 0 {
				// end program
				fmt.Println("our confrence is booked out comeback next year")
				break
			}
		} else{
			fmt.Println("Your input data is invalid")
		}
		
	}

	
	

	
}