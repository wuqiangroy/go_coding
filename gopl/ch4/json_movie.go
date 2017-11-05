package main

import (
	"fmt"
	"encoding/json"
	"log"
)

type Movie struct {
	Title string
	Year int 
	Color bool
	Actors []string
}

func main() {
	var movie = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors:[]string{"steve MacQueen", "Jacqueline Bisset"}},
	}
	data, err := json.Marshal(movie)
	if err != nil {
		log.Fatalf("JSON marshling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// formating print
	data2, err := json.MarshalIndent(movie, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	// json uncode
	var titles []struct { Title string}
	if err := json.Unmarshal(data2, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}