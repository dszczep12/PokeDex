package main

import (
    "bufio"
    "os"
	"io"
    "fmt"
	"net/http"
	"strings"
	"encoding/json"
	
)

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`

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

func getPokemonData() (*Pokemon, error) {
    pokemon := Reader()
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
    response, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making GET request: %v", err)
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("received non-200 response status: %d", response.StatusCode)
    }

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }

    var pokemonData Pokemon
    err = json.Unmarshal(body, &pokemonData)
    if err != nil {
        return nil, fmt.Errorf("error unmarshaling response: %v", err)
    }

    return &pokemonData, nil
}

func main() {
	pokemonData, err := getPokemonData()
    if err != nil {
        fmt.Printf("Error fetching Pokémon data: %v\n", err)
        return
    }

    fmt.Printf("ID: %d\n", pokemonData.Id)
    fmt.Printf("Name: %s\n", pokemonData.Name)
}