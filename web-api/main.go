package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	Item string
}

var data []Todo

func apiResponse(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(data)
	case http.MethodPost:
		if r.Body == nil {
			json.NewEncoder(w).Encode("Please add some value")
		} else {
			var newTodo Todo
			err := json.NewDecoder(r.Body).Decode(&newTodo)
			if err != nil {
				json.NewEncoder(w).Encode(err)
			}
			data = append(data, newTodo)
			fmt.Println(data)
			json.NewEncoder(w).Encode("Data has been saved successfully")

		}
	case http.MethodDelete:
		fmt.Fprintf(w, "THIS is a DELETE Request")
	default:
		fmt.Fprintf(w, "THIS IS unhandled request")

	}

}
func main() {
	fmt.Println("Let's built our own api")

	mux := http.NewServeMux()
	mux.HandleFunc("/", apiResponse)

	http.ListenAndServe(":3000", mux)

}
