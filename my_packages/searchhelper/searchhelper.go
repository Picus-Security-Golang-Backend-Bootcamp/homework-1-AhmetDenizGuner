package searchhelper

import (
	"errors"
	"strings"
)

//this method search the matches between the given string and bookList element, and return the matches list
func Search(bookList []string, searchItem string) ([]string, error) {
	searchResult := []string{}

	for _, book := range bookList {
		if strings.Contains(strings.ToLower(book), strings.ToLower(searchItem)) {
			searchResult = append(searchResult, book)
		}
	}

	//set error, if there is no result
	if len(searchResult) == 0 {
		err := errors.New("Aradiginiz kitap bulunamadi lutfen baska bir kelime/kelimelerle deneyiniz!")
		return nil, err
	}
	return searchResult, nil
}

//this method return a string using by elements of given string list
func BuildSearchItem(argumentSlice []string) string {
	return strings.TrimSpace(strings.Join(argumentSlice, " "))
}
