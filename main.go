package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {

	keyAspectsOfGo()
	keyAspectsOfGoWebApp()
	//fmt.Println("Hello World!")
}

var messages = []string{
	"Now is the time for all good Devops to come the aid of their servers.",
	"Alas poor Altair 8800; I knew it well!",
	"In the beginning there was ARPA Net and its domain was limited.",
	// assume many more
	"A blog a day helps keep the hacker away.",
	"These nutzs",
}

var spec = ":8080" // means localhost:8080

func sendRandomMessage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add(http.CanonicalHeaderKey("content-type"),
		"text/plain")
	w.Write([]byte(messages[rand.Intn(len(messages))] + "\n"))
}

func keyAspectsOfGoWebApp() {
	fmt.Println("-----------------------------")
	fmt.Println("Web Application")

	http.HandleFunc("/message", sendRandomMessage)
	if err := http.ListenAndServe(spec, nil); err != nil {
		log.Fatalf("Failed to start server on %s: %v", spec, err)
	}

}

func keyAspectsOfGo() {
	fmt.Println("-----------------------------")
	fmt.Println("Key ASpects Of Go")

	//  Java style of development
	fmt.Println("Java stlye development")
	for index, arg := range os.Args {
		if index == 0 {
			fmt.Printf("Program Name = %s \n", arg)
		} else {
			fmt.Printf("arg : %s\n", arg)
		}
	}

	//  Idiomatic go style
	// less nesting of if statements

	fmt.Println("\nIdiomatic Go stlye development")
	for index, arg := range os.Args {
		if index == 0 {
			fmt.Printf("Program Name = %s \n", arg)
			continue
		}
		fmt.Printf("arg : %s\n", arg)
		// less nesting use of continue and break
	}

	fmt.Println("\nIdiomatic Go stlye #2 development")
	fmt.Printf("Program: %s\n", os.Args[0])
	for index := 1; index < len(os.Args); index++ {
		fmt.Printf("Argument %d: %s\n", index, os.Args[index])
	}
	// no if statement

	fmt.Println("\nIdiomatic Go stlye #3 development")
	for index, arg := range os.Args {
		switch {
		case index == 0:
			fmt.Printf("Program Name = %s \n", arg)
		default:
			fmt.Printf("Argument %d: %s\n", index, os.Args[index])
		}
	}

	fmt.Println("\nIdiomatic Go stlye #4 development")
	for index, arg := range os.Args {
		switch index {
		case 0:
			fmt.Printf("Program Name = %s \n", arg)
		default:
			fmt.Printf("Argument %d: %s\n", index, os.Args[index])
		}
	}
}
