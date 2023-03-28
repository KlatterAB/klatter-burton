package db

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/PatrikOlin/skvs"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfigDatabase struct {
	Host     string `yaml:"KB_HOST"`
	Name     string `yaml:"KB_NAME"`
	Port     string `yaml:"KB_PORT"`
	User     string `yaml:"KB_USER"`
	Password string `yaml:"KB_PASSWORD"`
}

type WorkLogReturn struct {
	TotalHoursWorked float64 `db:"total_hours_worked"`
}

var cfg ConfigDatabase

var Store *skvs.KVStore
var db *sqlx.DB

func InitStore() {
	dir := os.Getenv("HOME") + "/.klatterburton/"
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
	db, err = sqlx.Open("postgres", getPqslInfo())

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
	inputStmt := "INSERT INTO work_log (minutes_worked, project_id, worker_id) VALUES ($1, $2, $3)"
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
	inputStmt := "INSERT INTO work_log (minutes_worked, project_id, worker_id) VALUES ($1, $2, $3)"
	_, err := db.Exec(inputStmt, minutes, projectId, userId)
	if err != nil {
		return err
	}

	return nil
}

func GetWorkLog(projectID, workerID, fromDate, toDate string) (string, error) {
	params := struct {
		ProjectID string `db:"project_id"`
		WorkerID  string `db:"worker_id"`
		FromDate  string `db:"fromDate"`
		ToDate    string `db:"toDate"`
	}{
		ProjectID: projectID,
		WorkerID:  workerID,
		FromDate:  fromDate,
		ToDate:    toDate,
	}

	stmt := buildWorkLogStmt(workerID, fromDate, toDate)
	var workLog WorkLogReturn

	nstmt, err := db.PrepareNamed(stmt)
	if err != nil {
		return "", err
	}

	err = nstmt.Get(&workLog, params)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return strconv.FormatFloat(workLog.TotalHoursWorked, 'f', 2, 64), nil
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
	dir := os.Getenv("HOME") + "/.klatterburton/"
	err := cleanenv.ReadConfig(dir+"kb.yml", &cfg)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error loading kb.yml")
	}
}

func buildWorkLogStmt(workerID, fromDate, toDate string) string {
	var sb strings.Builder

	sb.WriteString("SELECT SUM(minutes_worked) / 60.0 AS total_hours_worked FROM work_log WHERE project_id = :project_id ")
	if workerID != "" {
		sb.WriteString("AND worker_id = :worker_id ")
	}
	if fromDate != "" && toDate != "" {
		sb.WriteString("AND date BETWEEN DATE(:fromDate) AND DATE(:toDate)")
	}

	return sb.String()
}
