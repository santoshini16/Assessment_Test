package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type employees struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func createNewPerson(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee employees
	json.Unmarshal(reqBody, &employee)
	db.Create(&employee)
	fmt.Println("Create New Person")
	json.NewEncoder(w).Encode(employee)
}
func main() {
	db, err = gorm.Open("mysql", "root:root@123@tcp(127.0.0.1:3306)/sk2?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Connection failed to open!")
	} else {
		fmt.Println("Connection Established!")
	}
	defer db.Close()
	db.AutoMigrate(&employees{})
	router := mux.NewRouter()
	router.HandleFunc("/person", createNewPerson).Methods("POST")
	http.ListenAndServe(":8000", router)
}
