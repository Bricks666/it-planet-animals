package shared

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ENV map[string]string

func init() {

	var err error
	fmt.Println(os.Getwd())
	ENV, err = godotenv.Read("../.env")

	if err != nil {
		panic(err)
	}
}
