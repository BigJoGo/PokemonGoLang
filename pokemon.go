package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Pokemon struct {
	Name    string
	Type    string
	HP      int
	Attack  int
	Defense int
}

// Battle func fur Pokemon.

func Battle(pokemon1, pokemon2 *Pokemon) string {
	rand.Seed(time.Now().UnixNano())

	var number1 = rand.Intn(20)
	var spread1 = rand.Intn(1) - 1
	randomAttack1 := number1 + spread1

	var number2 = rand.Intn(20)
	var spread2 = rand.Intn(1) - 1
	randomAttack2 := number2 + spread2

	var number3 = rand.Intn(20)
	var spread3 = rand.Intn(1) - 1
	randomDefense2 := number3 + spread3

	var number4 = rand.Intn(20)
	var spread4 = rand.Intn(1) - 1
	randomDefense1 := number4 + spread4

	//Damage Pokemon
	damege1 := (pokemon1.Attack + randomAttack1) - (pokemon2.Defense + randomDefense2)
	damege2 := (pokemon2.Attack + randomAttack2) - (pokemon1.Defense + randomDefense1)

	//HP

	pokemon1.HP -= damege2
	pokemon2.HP -= damege1
	fmt.Printf("Name: %v, HP: %v - Damage %v !!! Name: %v, HP: %v - Damage: %v  ", pokemon1.Name, pokemon1.HP, damege2, pokemon2.Name, pokemon2.HP, damege1)
	//Win

	if pokemon1.HP <= 0 && pokemon2.HP <= 0 {
		return " Draw!"
	} else if pokemon1.HP <= 0 {
		return pokemon2.Name + " Winner!"
	} else if pokemon2.HP <= 0 {
		return pokemon1.Name + " Winner!"
	} else {
		return "the fight continues"
	}

}
func main() {
	db, err := sql.Open("sqlite3", "./pokemon.db")
	if err != nil {
		fmt.Println("Error open data base:", err)
		os.Exit(1)
	}
	defer db.Close()

	//Create Table Pokemon
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pokemon (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		type TEXT,
		hp INTEGER,
		attack INTEGER,
		defense INTEGER
	)`)
	if err != nil {
		fmt.Println("Error create table", err)
		os.Exit(1)
	}
	// Table info.

	if err != nil {
		fmt.Println("Error Table Info", err)
		os.Exit(1)
	}

	rows, err := db.Query("SELECT name, type, hp, attack, defense FROM pokemon")
	if err != nil {
		fmt.Println("Erorr select pokemon:", err)
		os.Exit(1)
	}
	defer rows.Close()

	var pokemons []*Pokemon
	for rows.Next() {
		var name, ptype string
		var hp, attack, defense int
		if err := rows.Scan(&name, &ptype, &hp, &attack, &defense); err != nil {
			fmt.Println("Error scan pokemon:", err)
			continue
		}
		pokemons = append(pokemons, &Pokemon{Name: name, Type: ptype, HP: hp, Attack: attack, Defense: defense})
	}

	// pokemons = pokemons.append(pokemon, Abra43BaseSet)

	//Overlay scan pokemon.
	fmt.Println("Select Pokemon:")
	for i, pokemon := range pokemons {
		fmt.Printf("%d. %s (%s)\n", i+1, pokemon.Name, pokemon.Type)
	}

	var choice int
	fmt.Println("Select first pokemon (writte nummer):")
	fmt.Scanln(&choice)
	pokemon1 := pokemons[choice-1]

	fmt.Println("Select second pokemon (wrrite nummer):")
	fmt.Scanln(&choice)
	pokemon2 := pokemons[choice-1]

	for pokemon1.HP > 0 && pokemon2.HP > 0 {
		fmt.Println("Start battle for", pokemon1.Name, "and", pokemon2.Name)
		result := Battle(pokemon1, pokemon2)
		fmt.Println(result)
	}
}
