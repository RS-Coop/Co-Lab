package go_dev

import (
	"database/sql"
	//"fmt"
	// "os/exec"
	"os"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "sdev_db"
// )

/*Initializes connection to database
If succesful, return database object
If failed, return nil and throw error
*/
func Initialize() *sql.DB {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	psqlInfo := os.Getenv("POSTGRES_CONNECTION")
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	return db
}

// func Build() bool {
//
// 	var cmdOut []byte
//
// 	_, err = exec.Command("which", "psql").Run()
//
// 	if err == nil {
// 		cmdOut, err = exec.Command("echo", "$?").Output()
// 	} else {
// 		return false
// 	}
//
// 	exit := ""
//
// 	if err == nil {
// 		exit = string(cmdOut)
// 	} else {
// 		return false
// 	}
//
// 	if exit == "1" {
// 		err = exec.Command("sudo", "apt-get", "update").Run()
// 		cmdOut, err = exec.Command("sudo", "apt-get", "install", "postgresql", "postgresql-contrib").Output()
// 		exit := string(cmdOut)
//
// 		if err != nil || exit != "0" {
// 			return false
// 		}
// 	}
//
// 	err = exec.Command("sudo", "-u", "postgres", "psql").Run()
//
// 	if err != nil {
// 		return false
// 	}
//
// 	//err = exec.Command()
// 	return true
// }
