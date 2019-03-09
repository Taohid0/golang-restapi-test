package api

import (
	"cricetAPITest/database"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
)

type Team struct {
	Name string `json:"name"`
	Flag string `json:"flag"`
}

func TeamRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/",handleBulkGet)
	r.Get("/{id}", handleSingleGet)
	r.Post("/", handleCreate)
	r.Delete("/",handleDelete)
	r.Put("/",handlePut)
	return r
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	team := Team{}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &team)
	if err != nil {
		log.Println(err)
	}
	db := database.ConnecToDB()
	defer db.Close()

	_, err = db.Exec("INSERT INTO team(name,flag) VALUES(?,?)", team.Name, team.Flag)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	response, _ := json.Marshal(map[string]string{"data": "team created"})
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}

func handleSingleGet(w http.ResponseWriter, r *http.Request) {
	team := Team{}
	id := chi.URLParam(r, "id")

	db := database.ConnecToDB()
	err := db.QueryRow("SELECT name,flag from team where id=?", id).Scan(&team.Name, &team.Flag)

	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(team)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}

func handleBulkGet(w http.ResponseWriter, r *http.Request) {
	db := database.ConnecToDB()
	team := Team{}
	allTeams := make([]Team, 0)

	rows, err := db.Query("SELECT name,flag from team")

	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		err = rows.Scan(&team.Name,&team.Flag)
		if err != nil {
			fmt.Println(err)
		}
		allTeams = append(allTeams, team)

	}
	err = rows.Err()

	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(allTeams)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write(data)
	if err != nil {
		fmt.Println(err)
	}

}
func handleDelete(w http.ResponseWriter, r *http.Request){
	type deleteObject struct{
		Id int
	}
	var object deleteObject

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &object)
	if err != nil {
		log.Println(err)
	}
	db := database.ConnecToDB()
	defer db.Close()

	_, err = db.Exec("DELETE FROM team where id=?",object.Id)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(204)
	response, _ := json.Marshal(map[string]string{"data": "team deleted"})
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}

func handlePut(w http.ResponseWriter, r *http.Request){
	type updateObject struct{
		Name string
		Flag string
		Id int
	}
	object :=updateObject{}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &object)

	if err != nil {
		log.Println(err)
	}
	db := database.ConnecToDB()
	defer db.Close()
	fmt.Println(object)


	_, err = db.Exec("UPDATE team set name=?, flag=? where id=?",object.Name,object.Flag,object.Id)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(204)
	response, _ := json.Marshal(map[string]string{"data": "team updated"})
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}