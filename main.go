package main // import "github.com/harvey93/pipeline"

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Pipeline")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}

	port := os.Getenv("PORT")
	fmt.Println(port)

}
