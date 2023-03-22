package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Movie struct {
	Name    string
	MovieId int
	Rating  int
	Tags    []string
}

func main() {
	fmt.Println("Hello, World!")
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/list", getMovieList).Methods("GET")
	router.HandleFunc("/addMovie/{id}", postOneMovie).Methods("POST")

	postSomethingToUrl("http://localhost:9200/movie/_doc")
	//getSomethingFromUrl("http://localhost:9200/movie/_search")
	//log.Fatal(http.ListenAndServe(":8000", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Accessing Homepage")
	w.Header().Set("Content-Type", "application/json")
}

func postSomethingToUrl(url string) {
	movie := Movie{"movie1", 2, 1, []string{"Science Fiction", "Aldults"}}
	movieJson, err := json.Marshal(movie)
	if err != nil {
		panic(err)
	}
	requestBody := strings.NewReader(string(movieJson))

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func getSomethingFromUrl(url string) {
	requestBody := strings.NewReader(`
	{
        "query": {
            "term": {
                "Name": "Iron man"
            }
        }
	}
	`)
	request, err := http.NewRequest("GET", url, requestBody)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	response, _ := client.Do(request)

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func postOneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Posting a movie")
	w.Header().Set("Content-Type", "application/json")

	//var movie Movie

}

func getMovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting movie list")
	w.Header().Set("Content-Type", "application/json")

	//myurl := "http://0.0.0.0:5601/app/home#/"

}
