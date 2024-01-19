//os.Clearenv() Method 
package main 

import ( 
	"fmt"
	"os"
) 

func main() { 

	// this delete's all environment variables 
	// Don't try this in the home environment, 
	// if you do so you will end up deleting 
	// all env variables 
	os.Clearenv() 

	fmt.Println("All environment variables cleared") 

}
