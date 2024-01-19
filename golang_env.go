//Setenv(), Getenv and Unsetenv method 
package main 

import ( 
	"fmt"
	"os"
) 
func main() { 

	// set environment variable Name
	os.Setenv("Name", "Jo") 

	// returns value of Name
	fmt.Println("Name:", os.Getenv("Name")) 

	// Unset environment variable GEEKS 
	os.Unsetenv("Name")
	// returns empty string and false, 
	// because we removed the GEEKS variable 
	value, ok := os.LookupEnv("Name")

	fmt.Println("Name:", value, " Is present:", ok) 

} 
