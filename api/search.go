package api

import (
	"cricetAPITest/database"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type playerData struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Age int `json:"age"`
	Team string `json:"team"`
	Flag string `json:"flag"`

}

func SearchRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/",handleSearch)

	return r
}


func handleSearch(w http.ResponseWriter, r *http.Request) {
	team_id := r.URL.Query().Get("team")
	age := r.URL.Query().Get("age")
	team_name := r.URL.Query().Get("team_name")
	fmt.Println(team_name)
	db := database.ConnecToDB()

	player := playerData{}
	allPlayers := make([]playerData, 0)
	queryString := "SELECT p.name,p.image,p.age,t.name,t.flag from team as t,player as p where p.team=t.id AND"

	//if team_id != "" || age!="" || team_name!=""{
	//	queryString +=" where"
	//}
	if team_id!=""{
		queryString+=" p.team="+team_id +" AND"// need to ask about it
	}
	if age!=""{
		queryString+=" p.age<"+age+" AND"
	}
	if team_name!=""{
		queryString+=fmt.Sprintf(" t.name LIKE %s AND","\"%b%\"")
	}
	queryString +=" 1"
	fmt.Println(queryString)
	rows, err := db.Query(queryString)

	if err != nil {
		fmt.Print(err)
	}
	for rows.Next() {
		err = rows.Scan(&player.Name,&player.Image,&player.Age,&player.Team,&player.Flag)
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
