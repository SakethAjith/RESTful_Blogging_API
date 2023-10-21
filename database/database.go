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

func (d *Details) getAuthDetails() *Details {
	file, err := ioutil.ReadFile("Config.yaml")
	if err != nil {
		log.Printf("yaml file get err #%v", err)
	}
	err = yaml.Unmarshal(file, d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return d
}

func InitDB() (*sqlx.DB, error) {
	var details Details
	details.getAuthDetails()

	psqlDetails := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		details.Host, details.Port, details.User, details.Password, details.DBName)

	db, err := sqlx.Connect("postgres", psqlDetails)
	if err != nil {
		return nil, err
	}

	return db, nil
}
