package shared

import (
	"animals/ent"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	Client  *ent.Client
	Context context.Context
}

var Database DB

func init() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)
	client, err := ent.Open("postgres", dsn)

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
