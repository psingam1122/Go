/*
Write a program which prompts the user to enter a string. The program searches through the entered string for 
the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character 
‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. 
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, 
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, 
“ihhhhhn”, “ina”, “xian”. 

Submit your source code for the program,
“findian.go”.
*/
package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var input string;
	var input1 string;
	var input2 string;
	
	fmt.Printf("Please enter a string : ")
	reader := bufio.NewReader(os.Stdin)


	input, _ = reader.ReadString('\n')
	//fmt.Scanln(&input)
	input1=strings.ToLower(input)
    input2= strings.Join(strings.Fields(input1), " ")
	if strings.HasPrefix(input2,"i") && strings.Contains(input2,"a") && strings.HasSuffix(input2,"n") {
		fmt.Printf("Found")
	}else{
		fmt.Printf("Not Found")
	}
	
}