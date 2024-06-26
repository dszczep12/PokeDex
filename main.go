package main

import (
    "bufio"
    "os"
    "fmt"
	"net/http"
	"strings"
    "github.com/TheZoraiz/ascii-image-converter/aic_package"
	"encoding/json"
    "github.com/briandowns/spinner"
	"time"
)

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Sprites struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`

}

func Reader() string {
    fmt.Println("What Pokémon info do you need?:")
    reader := bufio.NewReader(os.Stdin)   

    key, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    key = strings.TrimSpace(key)

    return key
}

func getPokemonData() (Pokemon, error) {
    pokemon := Reader()
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

func terminalOutput() {
    data, err := getPokemonData()
    if( err != nil) {
        fmt.Print("Api no have data :( ")
        getPokemonData()
    }
    
	filePath := data.Sprites.FrontDefault

	flags := aic_package.DefaultFlags()

	flags.Dimensions = []int{50, 25}
	flags.Colored = true
	flags.CustomMap = " .-=+#@"
	flags.SaveBackgroundColor = [4]int{50, 50, 50, 100}

	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		fmt.Println(err)
	}

    s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)  
	s.Start()                                                    
	time.Sleep(4 * time.Second)          
	s.Stop()

	fmt.Printf("%v\n", asciiArt)

}


func main() {
    terminalOutput()
}