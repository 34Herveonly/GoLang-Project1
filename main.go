package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello Go World")
	//favBooks := []string{"Get rich or die trying", "Time Is Money", "Holly Bible", "Rich Dad Poor Dad"}
	//printTask(favBooks)
	//
	//fmt.Println()
	//favBooks = addTask(favBooks, "Power of one Thing")
	//printTask(favBooks)
	//
	//fmt.Println()
	//favBooks = addTask(favBooks, "Ego is the enemy")
	//
	//printTask(favBooks)

	http.HandleFunc("/enter", EnterBook)

	http.ListenAndServe(":8008", nil)
}

func EnterBook(writer http.ResponseWriter, request *http.Request) {
	book := "Holy Bible"
	fmt.Fprintf(writer, "Book: %s", book)

}

func printTask(favBooks []string) {
	fmt.Println("List Of Books")
	fmt.Print()

	for index, favBook := range favBooks {
		//fmt.Println(favBook)
		fmt.Printf("%d. %s\n", index, favBook)
	}
}

func addTask(favBooks []string, newBook string) []string {
	var newB = append(favBooks, newBook)
	return newB
}
