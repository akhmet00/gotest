package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gotest/db"
	"gotest/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {

	_db := db.OpenDBConnection()

	db.CreateArticleTable(_db)

	handleRequests()

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ke")
	fmt.Println("eeee")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(model.ArticlesArray())
	if err != nil {
		return
	}
}

func handleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	//router.HandleFunc("/articles", returnAllArticles).Methods("GET")
	//router.HandleFunc("/articles/{id}", returnById)
	//router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/db", createArticleInDB).Methods("POST")
	router.HandleFunc("/articles/db", getDBArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDBArticles(w http.ResponseWriter, r *http.Request) {

	log.Println("Get Articles IN Method")


	w.Header().Set("Content-Type", "application/json")

	_db := db.OpenDBConnection()

	result, _ := _db.Query("SELECT * FROM articles")

	_db.Close()

	var article model.Article

	var jsonResult []model.Article

	for result.Next() {

		err := result.Scan(&article.Id, &article.Desc, &article.Content, &article.Title)
		if err != nil {
			log.Fatal(err)
		}

		jsonResult = append(jsonResult, model.Article{Id: article.Id, Desc: article.Desc, Content: article.Content, Title: article.Title})

	}
	log.Println(jsonResult)
	json.NewEncoder(w).Encode(&jsonResult)

}

func createArticleInDB(w http.ResponseWriter, r *http.Request) {

	log.Println("Create Articles IN Method")

	requestBody, _ := ioutil.ReadAll(r.Body)

	var article model.Article

	json.Unmarshal(requestBody, &article)

	log.Println(&article)

	_db := db.OpenDBConnection()

	query, err := _db.Query("INSERT INTO articles VALUES($1, $2, $3, $4)",
		&article.Id, &article.Desc, &article.Content, &article.Title)
	if err != nil {
		log.Fatal(err)
		return
	}
	query.Close()
	_db.Close()

}

func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requestBody, _ := ioutil.ReadAll(r.Body)

	articleList := model.ArticlesArray()

	newArticleList := make([]model.Article, len(*articleList), cap(*articleList)+1)

	var newArt model.Article

	json.Unmarshal(requestBody, &newArt)

	newArticleList = append(newArticleList, newArt)

	copy(newArticleList, *articleList)

	//articleList = append(articleList, json.Unmarshal(requestBody, &newArt))

	json.NewEncoder(w).Encode(newArticleList)
}

func returnById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	key := vars["id"]

	articleList := model.ArticlesArray()

	for _, article := range *articleList {
		if intKey, _ := strconv.Atoi(key); article.Id == intKey {
			json.NewEncoder(w).Encode(article)
		}

	}
}
