// This small program is just a small web server created in static mode
// in order to provide the smallest docker image possible

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
