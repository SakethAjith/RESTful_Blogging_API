package database

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type Details struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "david"
	password = "postgres"
	dbname   = "postgres"
)

func getAuthDetails() {
	auth, err := ioutil.ReadFile("Auth.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := Details{}
	err = yaml.Unmarshal(auth, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.Host)
	fmt.Println()
}

func InitDB() (*sqlx.DB, error) {
	psqlDetails := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlDetails)
	if err != nil {
		return nil, err
	}

	return db, nil
}
