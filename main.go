package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
			)

	func main() {
		r := mux.NewRouter()
		r.HandleFunc("/convert/{hex_value}", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			hex_value := vars["hex_value"]
		rgb, err := strconv.ParseUint(hex_value, 16, 32)
			if err != nil {
				fmt.Printf("Error: %v", err)
				return
			}
			fmt.Println("(",rgb >> 16,",",(rgb >> 8) & 0xFF,",", rgb & 0xFF,")")

		})
		http.ListenAndServe(":80", r)
	}

