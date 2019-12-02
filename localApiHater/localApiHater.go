package main

import (
	"net/http"
	"fmt"
	)

func expose() {
	for {
		res, error := http.Get("http://127.0.0.1:8080/beat")
		
		if error != nil {
			fmt.Println(error)
		}
		
		if res.StatusCode != 200 {
			fmt.Println("not work. =(")
			break;
		}

		fmt.Println("beated! >=D")
	}
}

func main() {
	expose()
}