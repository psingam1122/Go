/*
Write a program which prompts the user to first enter a name, and then enter an address.
 Your program should create a map and add the name and address to the map using the keys “name” and “address”, 
 respectively. Your program should use Marshal() to create a JSON object from the map, and then 
 your program should print the JSON object.
*/
package main
import (
	"encoding/json"
	"fmt"
)

func main(){
	var name string
	var address string
	    fmt.Println("Enter your name: ")
		fmt.Scanln(&name) 
		fmt.Println("Enter your address: ")
		fmt.Scanln(&address) 
	idMap := map[string]string{name : address}
	barr,_ :=json.Marshal(idMap)
    fmt.Println("values in json are:",string(barr) )	

	}