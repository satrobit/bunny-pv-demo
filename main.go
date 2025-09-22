package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type PageData struct {
	Content string
	AppID   string
	PodID   string
}

func main() {
	dir := "/test_pv1"
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal("Failed to create directory:", err)
	}

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/save", saveHandler).Methods("POST")

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	content := loadContent()

	data := PageData{
		Content: content,
		AppID:   os.Getenv("BUNNYNET_MC_APPID"),
		PodID:   os.Getenv("BUNNYNET_MC_PODID"),
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")

	if err := saveContent(content); err != nil {
		http.Error(w, "Failed to save content", http.StatusInternalServerError)
		return
	}

	// Redirect back to home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func loadContent() string {
	filePath := "/test_pv1/content.txt"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return ""
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		return ""
	}

	return string(content)
}

func saveContent(content string) error {
	filePath := "/test_pv1/content.txt"

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, []byte(content), 0644)
}
