package main 

import (
	"net/http"
	"strings"
	"fmt"
	)

  func tryEvade(w http.ResponseWriter, r *http.Request) {
	
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	
	fmt.Println("evaaaaading!")
	
	if strings.EqualFold(message, "beat") {
		w.Write([]byte("fuck you!"))
	}
  }

  func main() {
	http.HandleFunc("/beat", tryEvade)

	if err := http.ListenAndServe(":8080", nil); err != nil {
	  panic(err)
	}
  }

  











