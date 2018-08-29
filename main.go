package main // import "github.com/harvey93/pipeline"

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Pipeline")
	devP := flag.Bool("dev", false, "environment")
	flag.Parse()
	if *devP {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	port := os.Getenv("PORT")
	fmt.Println(port)

}
