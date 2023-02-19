package shared

import (
	"animals/ent"
	"context"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DB struct {
	Client  *ent.Client
	Context context.Context
}

var Database DB

func init() {
	ENV, _ := godotenv.Read(".env")

	var port = ENV["DB_PORT"]
	var host = ENV["DB_HOST"]
	var name = ENV["DB_NAME"]
	var user = ENV["DB_USER"]
	var password = ENV["DB_PASSWORD"]

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, name, password)

	client, err := ent.Open("postgres", dsn)

	fmt.Println(err, "Error")

	if err != nil {
		panic(err)
	}

	Database = DB{
		Client:  client,
		Context: context.Background(),
	}

	err = Database.Client.Schema.Create(Database.Context)
	fmt.Println(err, "Error")

	if err != nil {
		panic(err)
	}
}
