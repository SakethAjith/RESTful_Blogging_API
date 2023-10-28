package database

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/SakethAjith/RESTfulBlog/models"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var details Details
	details.getAuthDetails()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", details.Host, details.User, details.Password, details.DBName, details.Port)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// AutoMigrate will create the "blogs" table if it doesn't exist
	DB.AutoMigrate(&models.Blogs{})

	return DB, nil
}
