package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Movie struct {
	Title    string
	Genre    string
	Runtime  string
	Year     string
	Released string
}

func getMovie(title string) (Movie, error) {
	resp, err := http.Get(fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=923a1fcc", url.QueryEscape(title)))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return Movie{}, err
	}
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data[:]))
	dat := Movie{}
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(data, &dat)
		if err != nil {
			fmt.Println(err)
			return Movie{}, err
		}
	}

	return dat, nil
}

func main() {
	data, _ := getMovie("Avengers: Endgame")
	fmt.Printf("%+v", data)
}
