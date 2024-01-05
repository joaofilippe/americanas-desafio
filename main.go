package main

import "net/http"

func main() {
	http.HandleFunc("/save_lists", SaveLists)
	http.HandleFunc("/merge", Merge)

	http.HandleFunc("/", HelloWorld)
	
	http.ListenAndServe(":8080", nil)
}

// HelloWorld is a simple handler
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// Merge is a simple handler
func Merge(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Merge!"))
}

// SaveLists is a simple handler
func SaveLists(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SaveLists!"))
}