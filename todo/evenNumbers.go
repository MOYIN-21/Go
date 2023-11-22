package main
import "fmt"
func main(){
	 evenNumber(50)
}
	func evenNumber(num int){
		count := 0
		for count <= num {
			count++
			if count % 2 == 0{
				if count == 24{
					continue
				}
				fmt.Print(count)
			}
		}
	}
