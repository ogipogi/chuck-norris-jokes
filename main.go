package main

import (
	"fmt"
	"github.com/bndr/gopencils"
	"html"
	"log"
	"net/http"
)

type Joke struct {
	T string `json:"type"`
	V struct {
		ID   int      `json:"id"`
		Joke string   `json:"joke"`
		Cats []string `json:"categories"`
	} `json:"value"`
}

func main() {
	randomJoke := &Joke{}
	url := "http://api.icndb.com/jokes/"
	api := gopencils.Api(url)

	http.HandleFunc("/joke", func(w http.ResponseWriter, r *http.Request) {
		api.Res("random", randomJoke).Get()
		fmt.Fprintf(w, html.UnescapeString(randomJoke.V.Joke))
	})

	log.Fatalln(http.ListenAndServe(":8081", nil))

}
