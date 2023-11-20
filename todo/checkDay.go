package main
import "fmt"
func main(){
	fmt.println(checkDay(3))
}

func divide(x, y int, user string) (int, string){
	c := x/y
	myName := user
	return c, myName
}

func checkDay(day int) string {
	if day == 1 {
		return "weekday"
	}
	if day == 2 {
		return "weekday"
	}
	if day == 3 {
		return "weekday"
	}
	if day == 4 {
		return "weekday"
	}
	if day == 5 {
		return "weekday"
	}
	if day == 6 {
		return "weekday"
	}
	if day == 7 {
		return "weekday"
	}
	return "weekend"
}