package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// BlogPost -- struct for all blogposts
type BlogPost struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"content"`
}

// BlogPosts is a global variable for the list of all blogposts
var BlogPosts []BlogPost

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllBlogPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllBlogPosts")
	json.NewEncoder(w).Encode(BlogPosts)
}

func returnSingleBlogPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSinglePost")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, post := range BlogPosts {
		if post.ID == key {
			json.NewEncoder(w).Encode(post)
		}
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/blogposts", returnAllBlogPosts).Methods("GET")
	router.HandleFunc("/blogposts/{id}", returnSingleBlogPost)
	log.Fatal(http.ListenAndServe(":10001", router))
}
func main() {
	fmt.Println("Standard Rest API - Mux Router")

	BlogPosts = []BlogPost{
		BlogPost{ID: "1", Title: "First Post", Content: "lorem ipsum"},
		BlogPost{ID: "2", Title: "Second Post", Content: "lorem ipsum"},
		BlogPost{ID: "3", Title: "Third Post", Content: "lorem ipsum"},
		BlogPost{ID: "4", Title: "Fourth Post", Content: "lorem ipsum"},
	}
	handleRequests()
}
