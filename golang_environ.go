// os.Environ() Method 
package main 

import ( 
	"fmt"
	"os"
) 

func main() { 

	// list all environment variables and their values 
	for _, env := range os.Environ() { 
		fmt.Println(env) 
	} 
} 
