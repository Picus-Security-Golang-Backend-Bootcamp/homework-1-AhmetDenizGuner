package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"workspace/my_packages/searchhelper"
)

func main() {
	args := os.Args

	//if user didn't give any argument to the program, return this warning message that includes possible commands and stop program
	if len(args) == 1 {
		printDefaultErrorMessage(args[0])
		return
	}

	if strings.TrimSpace(args[1]) == "list" {
		bookPrinter(getBookSlice())

	} else if strings.TrimSpace(args[1]) == "search" {

		//check is there any search string
		if len(args) < 3 {
			fmt.Println("Search komutunu kullanabilmek için lütfen arama yapılacak kelime veya kelimleri de giriniz!")
			return
		}

		//creating search string using by program arguments
		//BuildSearchItem method get arguments and return the string these arguments
		searchItem := searchhelper.BuildSearchItem(args[2:])

		//check string that will be searched is equal or bigger than  3 char
		if len(searchItem) < 3 {
			fmt.Println("Lütfen daha uzun bir keime giriniz!")
			return
		}

		//Searching string in book list
		//Search method use strings contains method and seek any match in book list and return the all matches
		resultSlices, err := searchhelper.Search(getBookSlice(), searchItem)

		//if resultSlices is empty Serach method will be return error message and program terminated
		if err != nil {
			fmt.Println(err)
			return
		}

		//if there is search result , bookPrinter function print the books that are in given slice
		bookPrinter(resultSlices)

	} else { // If user gives wrong command , terminate program and print error message.
		printDefaultErrorMessage(args[0])
		return
	}

}

func printDefaultErrorMessage(projectPath string) {
	//taking executable project name from path
	projectName := filepath.Base(projectPath)
	fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", projectName)
}

//this function return the book slice that was created by manually
func getBookSlice() []string {
	books := []string{"Sherlock Holmes Red Leech", "Crime And Punishment", "Animal Farm", "Little Prince", "Sapiens", "Harry Potter and the Philosopher's Stone"}
	return books
}

//this function print the list that given
func bookPrinter(books []string) {
	for index, bookName := range books {
		fmt.Printf("%d) %s\n", index+1, bookName)
	}
}
