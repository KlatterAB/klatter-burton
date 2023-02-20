package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/PatrikOlin/skvs"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

type ConfigDatabase struct {
	Host     string `yaml:"KB_HOST"`
	Name     string `yaml:"KB_NAME"`
	Port     string `yaml:"KB_PORT"`
	User     string `yaml:"KB_USER"`
	Password string `yaml:"KB_PASSWORD"`
}

var cfg ConfigDatabase

var Store *skvs.KVStore
var db *sql.DB

func InitStore() {
	dir := os.Getenv("HOME") + "/.klatter-burton/"
	makeDirectoryIfNotExists(dir)

	dbfile := path.Join(dir, "data.db")

	var err error
	Store, err = skvs.Open(dbfile)
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB() error {
	var err error
	db, err = sql.Open("postgres", getPqslInfo())

	if err != nil {
		fmt.Println("ingen db")
		log.Fatalln(err)
	}

	return db.Ping()
}

func GetCheckinTime() (int64, error) {
	var valUnix int64
	if err := Store.Get("checkinUnix", &valUnix); err == skvs.ErrNotFound {
		log.Println("not found")
		return -1, err
	} else if err != nil {
		log.Fatal(err)
		return -1, err
	}

	return valUnix, nil
}

func SetMinutesWorked(minutes int, project, user string) error {
	inputStmt := "INSERT INTO minutes_worked (minutes, project_id, worker_id) VALUES ($1, $2, $3)"
	_, err := db.Exec(inputStmt, minutes, project, user)
	if err != nil {
		return err
	}

	return nil
}

func AddProject(name, id string) error {
	inputStmt := "INSERT INTO projects (project_id, name, status) VALUES ($1, $2, $3)"
	_, err := db.Exec(inputStmt, id, name, "started")
	if err != nil {
		return err
	}

	return nil
}

func AddTimeToProject(minutes int, projectId, userId string) error {
	inputStmt := "INSERT INTO minutes_worked (minutes, project_id, worker_id) VALUES ($1, $2, $3)"
	_, err := db.Exec(inputStmt, minutes, projectId, userId)
	if err != nil {
		return err
	}

	return nil
}

func makeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}

func getPqslInfo() string {
	readEnvFromFile()

	connStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)
	return connStr
}

func readEnvFromFile() {
	dir := os.Getenv("HOME") + "/.klatter-burton/"
	err := cleanenv.ReadConfig(dir+"kb.yml", &cfg)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error loading kb.yml")
	}
}
