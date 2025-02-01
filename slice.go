/*Write a program which prompts the user to enter integers and stores the integers in a sorted slice. 
The program should be written as a loop. Before entering the loop, the program should create an empty integer 
slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer 
to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents 
of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user d
ecides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ 
instead of an integer.*/
package main
import (
	"fmt"
	"strconv"
	"sort"
)

func main(){
	s1:= make([]int,0,3)
	var input string
	for{
	    fmt.Println("Enter any integer or 'X' to quit: ")
		fmt.Scanln(&input) 

		if input == "X" {
			break
	    }else{
                num, err := strconv.Atoi(input)
		        if err != nil {
			    fmt.Println("Invalid")
			  continue
		       }
				fmt.Print("Added to slice " )
				s1 = append(s1, num)
				sort.Ints(s1)
					fmt.Println("s1:", s1)
	    }
		
	}

}
