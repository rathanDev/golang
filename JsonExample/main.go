package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
	fmt.Println()
	fmt.Println("bird => " , bird)


	birdsJson := `[{"species": "pigeon","description": "likes to perch on rocks"}, {"species": "mina","description": "likes to live with people"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdsJson), &birds)
	fmt.Println("birds => ", birds)
}
