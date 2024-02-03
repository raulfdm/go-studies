package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"studies.com/httpexample/data"
)

const apiURL = "https://rickandmortyapi.com/api/character/1"

func main() {
	res, err := http.Get(apiURL)

	if err != nil {
		panic(err)
	}
	if res.StatusCode == http.StatusOK {

		bodyAsByte, _ := io.ReadAll(res.Body)

		var parsedApiResponse RickAndMortyApiCharacterResponse

		err := json.Unmarshal(bodyAsByte, &parsedApiResponse)

		if err != nil {
			panic(err)
		}

		character := data.Character{
			Name:    parsedApiResponse.Name,
			Species: parsedApiResponse.Species,
		}

		fmt.Printf("The name of the character is %s and it is a %s\n", character.Name, character.Species)
	} else {
		panic(errors.New("API request failed"))
	}
}
