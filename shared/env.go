package shared

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env map[string]string

var ENV Env = calculateENV()

func calculateENV() Env {

	var err error
	var env Env
	fmt.Println(os.Getwd())
	env, err = godotenv.Read(".env")

	if err != nil {
		panic(err)
	}

	return env
}
