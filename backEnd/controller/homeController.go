package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golangBaseProject/backEnd/model"
	"github.com/golangBaseProject/backEnd/services"
)


func GetAllUsers(c *gin.Context) {
	var users []model.Users
	services.Db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "None found!"})
		return
	}
	RestApiExample()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": users})
}

// A Response struct to map the Entire Response
type Response struct {
 //   Name    string    `json:"name"`
    Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
    EntryNo int            `json:"entry_number"`
    Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
    Name string `json:"name"`
}

func RestApiExample(){
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body) // read json and convert to byte
    if err != nil {
        log.Fatal(err)
    }

    var responseObject Response
    json.Unmarshal(responseData, &responseObject) // convert []byte(json) to struct 


	fmt.Println(" ------------------------------------ ")
	//	fmt.Println(string(responseData))
	fmt.Println(" ------------------------------------ ")
    //   fmt.Println(responseObject.Name)
    //   fmt.Println(len(responseObject.Pokemon))

    for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
    }

	
}