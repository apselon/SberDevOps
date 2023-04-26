package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Id             int    `json:"id"`
	IsDefault      bool   `json:"is_default"`
	Name           string `json:"name"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

func GetPokemon(i int) Pokemon {
	const url_format = "https://pokeapi.co/api/v2/pokemon/%d"
	url := fmt.Sprintf(url_format, i)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Bad response for URL:%s!\n", url)
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error while reading responce for URL:%s!\n", url)
		panic(err)
	}

	p := Pokemon{}

	err = json.Unmarshal(body, &p)

	if err != nil {
		fmt.Printf("Error while decoding json for URL:%s!\n", url)
		panic(err)
	}

	return p
}

func HandleGetBestPokemon(w http.ResponseWriter, r *http.Request) {
	bestPokemon := GetPokemon(rand.Intn(100))
	w.Write([]byte("The best pokemon is " + bestPokemon.Name + "!\n"))
}

func main() {
	mux := http.NewServeMux()
	handler := http.HandlerFunc(HandleGetBestPokemon)

	mux.Handle("/best", handler)

	fmt.Println("Start")

	http.ListenAndServe(":8088", mux)

	fmt.Println("Finish")
}
