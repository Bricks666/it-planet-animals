package shared

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env map[string]string

var ENV = calculateENV()

var DB_PORT, ROUND_COUNT int
var DB_HOST, DB_USER, DB_PASSWORD, DB_NAME string

func calculateENV() Env {

	var err error
	var env Env = Env{}

	env, err = godotenv.Read(".env")

	if err != nil {
		env["DB_HOST"] = os.Getenv("DB_HOST")
		env["DB_PORT"] = os.Getenv("DB_PORT")
		env["DB_USER"] = os.Getenv("DB_USER")
		env["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
		env["DB_NAME"] = os.Getenv("DB_NAME")
		env["ROUND_COUNT"] = os.Getenv("ROUND_COUNT")
	}

	DB_HOST = env["DB_HOST"]
	DB_PORT, _ = strconv.Atoi(env["DB_PORT"])
	DB_USER = env["DB_USER"]
	DB_PASSWORD = env["DB_PASSWORD"]
	DB_NAME = env["DB_NAME"]
	ROUND_COUNT, _ = strconv.Atoi(env["ROUND_COUNT"])

	return env
}
