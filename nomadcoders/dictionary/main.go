package main

import (
	"fmt"
	"nomadcoders/dictionary/mydict"
)

func printSearch(d mydict.Dictionary, word string) {
	definition, err := d.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Search : ", definition)
	}
}

func printAdd(d mydict.Dictionary, word string, def string) {
	err := d.Add(word, def)
	if err != nil {
		fmt.Println(err)
	} else {
		printSearch(d, word)
	}
}

func printUadate(d mydict.Dictionary, baseWord string, word string) {
	err := d.Update(baseWord, word)
	if err != nil {
		fmt.Println(err)
	} else {
		printSearch(d, baseWord)
	}
}

func printDelete(d mydict.Dictionary, word string) {
	err := d.Delete(word)
	if err != nil {
		fmt.Println(err)
	} else {
		printSearch(d, word)
	}
}

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	baseWord := "hello"

	printSearch(dictionary, "first")
	printSearch(dictionary, "second")

	printAdd(dictionary, baseWord, "Greeting")
	printUadate(dictionary, baseWord, "Second")

	printDelete(dictionary, baseWord)

	fmt.Println(dictionary)
}
