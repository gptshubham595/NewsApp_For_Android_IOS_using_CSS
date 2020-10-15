package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type article struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allArticles []article

var articles = allArticles{
	{
		ID:          "1",
		Title:       "First Article",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home! \n \n GET: /articles gives all articles \n\n POST: /articles with Body>Raw as {\"id\":\"NUMBER\",\"title\":\"Give your title\",\"description\":\"GIve your Description\",} \n\n DELETE:/articles/2 to delete article number 2 \n\n PATCH: /articles/1 in BODY>raw {\"title\":\"Write new title for the article\",\"description\":\"Write here updated description for the article\"} \n SEARCH: /articles/search/<query> It will List every where found")
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	var newArticle article
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event id, title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newArticle)

	// Add the newly created event to the array of events
	articles = append(articles, newArticle)

	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)
	// Return the newly created event
	json.NewEncoder(w).Encode(newArticle)
}

var decoder = schema.NewDecoder()

func searchArticle(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]
	words := []string{articleID}
	for _, singleArticle := range articles {
		if singleArticle.ID == articleID {
			json.NewEncoder(w).Encode("FOUND AT ID:")
			json.NewEncoder(w).Encode(singleArticle)

		}
		if len(WordsInSentence(words, singleArticle.Title)) > 0 {
			json.NewEncoder(w).Encode("FOUND AT TITLE")
			json.NewEncoder(w).Encode(singleArticle)
		}
		if len(WordsInSentence(words, singleArticle.Description)) > 0 {
			json.NewEncoder(w).Encode("FOUND AT Description:")
			json.NewEncoder(w).Encode(singleArticle)
		}
	}

}
func WordsInSentence(words []string, sentence string) []string {
	var in []string

	dict := make(map[string]string, len(words))
	for _, word := range words {
		dict[strings.ToLower(word)] = word
	}

	f := func(r rune) bool { return !unicode.IsLetter(r) }
	for _, word := range strings.FieldsFunc(sentence, f) {
		if word, ok := dict[strings.ToLower(word)]; ok {
			in = append(in, word)
			delete(dict, word)
		}
	}

	return in
}

func getOneArticle(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	articleID := mux.Vars(r)["id"]

	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for _, singleArticle := range articles {

		if singleArticle.ID == articleID {
			json.NewEncoder(w).Encode(singleArticle)
		}
	}
}

func getAllArticle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	articleID := mux.Vars(r)["id"]
	var updatedArticle article
	// Convert r.Body into a readable formart
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &updatedArticle)

	for i, singleArticle := range articles {
		if singleArticle.ID == articleID {
			singleArticle.Title = updatedArticle.Title
			singleArticle.Description = updatedArticle.Description
			articles[i] = singleArticle
			json.NewEncoder(w).Encode(singleArticle)
		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	articleID := mux.Vars(r)["id"]

	// Get the details from an existing event
	// Use the blank identifier to avoid creating a value that will not be used
	for i, singleArticle := range articles {
		if singleArticle.ID == articleID {
			articles = append(articles[:i], articles[i+1:]...)
			fmt.Fprintf(w, "The article with ID %v has been deleted successfully", articleID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles", getAllArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", getOneArticle).Methods("GET")
	router.HandleFunc("/articles/search/{id}", searchArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PATCH")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
