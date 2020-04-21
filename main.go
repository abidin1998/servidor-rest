package main

import (
	"io/ioutil"
	"log"
	"os"
	"net/http"
  "github.com/go-chi/cors"
  "github.com/gorilla/mux"
)

func getperfil(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
  resp, err := http.Get("https://euw1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/"+vars["id"]+"?api_key=RGAPI-2b76108a-349f-45db-bd9d-4133c9699658")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  b, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  w.Write([]byte(string(b)))
}
func indexRoute(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  resp, err := http.Get("https://euw1.api.riotgames.com/lol/summoner/v4/summoners/by-name/"+vars["nombre"]+"?api_key=RGAPI-2b76108a-349f-45db-bd9d-4133c9699658")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  b, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }


  w.Write([]byte(string(b)))
}

func historial(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  resp, err := http.Get("https://euw1.api.riotgames.com/lol/match/v4/matchlists/by-account/"+vars["idaccount"]+"?api_key=RGAPI-2b76108a-349f-45db-bd9d-4133c9699658")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  b, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }


  w.Write([]byte(string(b)))
}

func partidainfo(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  resp, err := http.Get("https://euw1.api.riotgames.com/lol/match/v4/matches/"+vars["idmach"]+"?api_key=RGAPI-2b76108a-349f-45db-bd9d-4133c9699658")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  b, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }


  w.Write([]byte(string(b)))
}

func perfilinfo(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  resp, err := http.Get("https://euw1.api.riotgames.com/lol/league/v4/entries/by-summoner/"+vars["id"]+"?api_key=RGAPI-2b76108a-349f-45db-bd9d-4133c9699658")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  b, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }


  w.Write([]byte(string(b)))
}

func main() {
  r := mux.NewRouter().StrictSlash(true)
  // Basic CORS
  // for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
  cors := cors.New(cors.Options{
    // AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: true,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  })
  r.Use(cors.Handler)
  r.HandleFunc("/perfil/{nombre}", indexRoute)
	r.HandleFunc("/perfil/historial/{idaccount}", historial)
	r.HandleFunc("/perfil/maxmastery/{id}", getperfil)
	r.HandleFunc("/perfil/partida/{idmach}", partidainfo)
	r.HandleFunc("/perfil/info/{id}", perfilinfo)
  //r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//    getperfil(w)
//  }
//)
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}
  http.ListenAndServe(":"+port, r)

}
