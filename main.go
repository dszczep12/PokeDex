package main

import (
    "bufio"
    "os"
    "fmt"
	"net/http"
	"strings"
	"encoding/json"
	
)

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`

}

func Reader() string {
    reader := bufio.NewReader(os.Stdin)   
    fmt.Println("What Pokémon info do you need?:")

    key, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    key = strings.TrimSpace(key)

    return key
}

func getPokemonData() (Pokemon, error) {
    pokemon := "pikachu"
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
	
    response, err := http.Get(url)
    if err != nil {
        return Pokemon{}, err
    }
    defer response.Body.Close()

	var result Pokemon
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
        return Pokemon{}, err
    }

	return result, nil
}

func main() {
    pokemonData, err := getPokemonData()
    if err != nil {
        fmt.Printf("Error fetching Pokémon data: %v\n", err)
        return
    }

    fmt.Println("name and id")
    fmt.Println(pokemonData)
}