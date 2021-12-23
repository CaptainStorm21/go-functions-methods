package main  

import (
	"fmt"
	//"errors"
)

//function declaration
func main () {
	port := 3000
	// port, err := startWebServer(port)
	_, err := startWebServer(port)
	// fmt.Println(port, err)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting over...")

	// do importnat stuff
	fmt.Println("Started server on ", port)
	// fmt.Println("Number of tries ", numberofRetries)
	// return errors.New("somthing is wrong")
	return port, nil
}