package main

import (
	"cricetAPITest/api"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func main()  {
	//db := database.ConnecToDB()
	//fmt.Print(db)

	r:=chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(Recovery)
	r.Mount("/team",api.TeamRouter())
	r.Mount("/player",api.PlayerRouter())
	r.Mount("/search",api.SearchRouter())

	err := http.ListenAndServe(":3001",r)
	if err != nil {
		log.Fatal(err)
	}
}

func Recovery(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func() {
			err := recover()
			if err !=nil{
				fmt.Print(err)
				errosResponse,_ := json.Marshal(map[string]string{
					"error":"Internal server error",
				})
				w.Header().Set("Content-Type","application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_,err:=w.Write(errosResponse)

				if err != nil {
					fmt.Println(err)
				}
			}
		}()
		next.ServeHTTP(w,r)
	})
}