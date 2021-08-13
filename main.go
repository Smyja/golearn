package main
import "fmt"

func main(){
	// ":-" declares a variable
	// [] is used for a slice(an array that contains the same type of elements)
	
	cards:= []string{"Ace of Diamonds",newCard()}
	cards=append(cards,"Queens Knight")
	for i, card:=range cards{
		fmt.Println(i,card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}