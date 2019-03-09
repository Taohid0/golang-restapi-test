package api

import (
	"cricetAPITest/database"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Player struct {
	Team  int    `json:"team"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Age   int    `json:"age"`
}

func PlayerRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", playerHandleBulkGet)
	r.Get("/{id}", playerHandleSingleGet)
	r.Post("/", playerHandleCreate)
	r.Delete("/", playerHandleDelete)
	r.Put("/", playerHandlePut)
	return r
}

func playerHandleCreate(w http.ResponseWriter, r *http.Request) {
	player := Player{}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &player)
	if err != nil {
		log.Println(err)
	}
	db := database.ConnecToDB()
	defer db.Close()

	_, err = db.Exec("INSERT INTO player(name,image,age,team) VALUES(?,?,?,?)", player.Name, player.Image, player.Age, player.Team)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	response, _ := json.Marshal(map[string]string{"data": "player created"})
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}

func playerHandleSingleGet(w http.ResponseWriter, r *http.Request) {
	player := Player{}
	id := chi.URLParam(r, "id")

	db := database.ConnecToDB()
	err := db.QueryRow("SELECT name,image,age,team from player where id=?", id).Scan(&player.Name, &player.Image, &player.Age, &player.Team)

	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(player)
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

func playerHandleBulkGet(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)

	db := database.ConnecToDB()
	player := Player{}
	allPlayers := make([]Player, 0)

	rows, err := db.Query("SELECT name,image,age,team from player")

	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		err = rows.Scan(&player.Name, &player.Image, &player.Age, &player.Team)
		if err != nil {
			fmt.Println(err)
		}
		allPlayers = append(allPlayers, player)

	}
	err = rows.Err()

	if err != nil {
		fmt.Println(err)
	}

	data, err := json.Marshal(allPlayers)
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
func playerHandleDelete(w http.ResponseWriter, r *http.Request) {
	type deleteObject struct {
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

	_, err = db.Exec("DELETE FROM player where id=?", object.Id)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(204)
	response, _ := json.Marshal(map[string]string{"data": "player deleted"})
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}

func playerHandlePut(w http.ResponseWriter, r *http.Request) {
	type updateObject struct {
		Name  string
		Image string
		Age   int
		Team  int
		Id    int
	}
	object := updateObject{}
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

	_, err = db.Exec("UPDATE player set name=?,image=?,age=?,team=?  where id=?", object.Name, object.Image, object.Age, object.Team, object.Id)

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(204)
	response, _ := json.Marshal(map[string]string{"data": "player updated"})
	_, err = w.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}
