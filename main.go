package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Articles []Article

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticleById).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// GET BY ID
func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
	for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

// DELETE BY ID
func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

// GET ALL
func returnAllArticles(w http.ResponseWriter, r *http.Request){
    json.NewEncoder(w).Encode(Articles)
}

func updateArticleById(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
    var updatedArticle Article 
    json.Unmarshal(reqBody, &updatedArticle)
	vars := mux.Vars(r)
    id := vars["id"]

    for _, article := range Articles {
        if article.Id == id {
			article = updatedArticle
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    Articles = append(Articles, article)
    json.NewEncoder(w).Encode(article)
}

func main() {
	Articles = []Article{
        {Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        {Id: "2", Title: "Teste Secundário", Desc: "Alguna Description", Content: "Article Content"},
    }
	handleRequests()
}