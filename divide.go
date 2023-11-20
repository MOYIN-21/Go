package main
import "fmt"
func main(){
	fmt.Println(divide(12, 6))
}

	func divide(x int, y float64)(int, string){
		c := x/ int(y)
		return c, "myName"
	}
