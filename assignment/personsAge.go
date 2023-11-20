package main
import ( "fmt"
)

func main(){
	fmt.Println("Enter your age and i will determines the ticket price for your movie theater")
	var age int
	fmt.Scan(&age)

	price := ticketPrice(age)
	fmt.Print("The student grade is", price)
}

func ticketPrice(age int) string{
	switch {
		case age  <= 12:
			return "N5"
		case age > 12 && age <= 64:
			return "N10"
		case age >= 65: 
			return "N7"
		default:
			return "Hi"

	}
}