package main

import (
	"fmt"
	"net/http"
	"words.com/mypkg"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(mypkg.GetWord())
		fmt.Fprintf(w, mypkg.GetWord())
	})

	const hello = "HELLO"

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(hello)
		fmt.Fprintf(w, hello)
	})

	http.ListenAndServe(":8080", nil)

	fmt.Println("Listening on port 8080")

}
