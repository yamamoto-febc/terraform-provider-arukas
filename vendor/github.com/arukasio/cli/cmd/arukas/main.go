package main

import (
	arukas "github.com/arukasio/cli"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	arukas.Run(os.Args)
}
