package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
)

// Config struct for webapp config
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"database"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readConfig(cfg *Config) {
	f, err := os.Open("config.yaml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func main() {
	var cfg Config
	readConfig(&cfg)

	// get user hash
	userHash := os.Args[1]
	// fmt.Printf("%T\n", userHash)
	// connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the hash base!")
	// loop through userhash and search in the database for its entry
	sqlStatement := `SELECT plain_text FROM data WHERE hash=$1;`
	var plainText string
	row := db.QueryRow(sqlStatement, userHash)
	switch err := row.Scan(&plainText); err {
	case sql.ErrNoRows:
		fmt.Println("Sorry, that hash is not stored in our database :(")
	case nil:
		fmt.Printf("The plain text password for your hash is => %s", plainText)
	default:
		panic(err)
	}

	// if found then return the plain text for the hash

}
