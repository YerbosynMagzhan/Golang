package pkg

import (
	"encoding/json"
	"myapp/internal"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("ui/html/homepage.html")
	// if err != nil {
	// 	http.Error(w, "Server Error", http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	book := internal.GetBookData()

	json.NewEncoder(w).Encode(book)

	// tmpl.Execute(w, book)
}

func MainChar(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("ui/html/maincharpage.html")
	// if err != nil {
	// 	http.Error(w, "Server Error", http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")

	parametrs := mux.Vars(r)
	charName := parametrs["name"]

	character, found := internal.GetCharacterByName(charName)
	if !found {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(character)

	// tmpl.Execute(w, character)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	healthCheckMap := map[string]string{
		"text":   "My app is about a book. There is a name of title, genre, name of the book, description and characters",
		"author": "Yerbosyn Magzhan",
	}

	json.NewEncoder(w).Encode(healthCheckMap)
}
