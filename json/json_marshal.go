package json

import (
	"encoding/json"
	"fmt"
)

type Movie struct{
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`// omitempty express if Color is zero value (eg, bool is false) will not show in the json string
	Actors []string
}

func JsonMarshal() (string, error){
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Bullitt", Year: 1943, Color: true, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1944, Color: true, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	}

	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		fmt.Errorf("JSON marshaling failed: %s", err)
		return "",err
	}
	fmt.Printf("%s\n", data)
	return string(data),err
}

func JsonUnMarshal(jsonStr string) (error){
	var movies []Movie
	if err := json.Unmarshal([]byte(jsonStr), &movies); err != nil{
		fmt.Errorf("JSON unmarshaling failed: %s", err)
		return err
	}
	for _, movie := range movies{
		fmt.Println(movie)
	}
	return nil
}
