package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/geo-albin/projects/webdb/db"
	"github.com/geo-albin/projects/webdb/web/webhandler"
)

func main() {

	resetDB := flag.Bool("reset-db", false, "reset the DB")
	populateDB := flag.Bool("populate-db", false, "populate the DB with sample data")
	port := flag.Int("port", 3000, "webserver listening port")
	flag.Parse()

	wh := webhandler.WebHandler{
		Port:       int16(*port),
		Started:    false,
		PopulateDB: *populateDB,
		ResetDB:    *resetDB,
	}

	cont, err := startup(&wh)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if false == cont {
		return
	}

	users, err := db.GetAllUsers()
	if nil != err {
		return
	}

	for _, u := range users {
		fmt.Println(u.ID, " ", u.FirstName, " ", u.LastName, " ", u.Age)
	}

	select {}
}

func startup(wh *webhandler.WebHandler) (bool, error) {

	//connect to DB
	if err := db.ConnectToDB(); err != nil {
		return false, err
	}
	fmt.Println("DB connection successful")

	if wh.ResetDB {
		err := db.DeleteAllUsers()
		fmt.Println("DB users table reset successful")
		return false, err
	}

	if wh.PopulateDB {
		populateDB()
		fmt.Println("DB users table population successful")
		return false, nil
	}

	wh.RegisterWeb()
	go wh.StartWeb()

	return true, nil
}

func populateDB() {
	var u db.User

	for i := 0; i < 100; i++ {
		u.FirstName = "DBUser_" + strconv.FormatInt(int64(i), 10)
		u.LastName = "DBUser_Ln"
		u.Age = uint8(i + 1)

		db.InsertUser(&u)
	}
}
