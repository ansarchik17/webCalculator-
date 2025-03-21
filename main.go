package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/substract", handleSubstract)
	http.HandleFunc("/multiply", handleMultiply)
	http.HandleFunc("/divide", handleDivide)

	http.ListenAndServe(":8080", nil)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'a'", 400)
		return
	}
	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'b'", 400)
		return
	}
	result := a + b
	fmt.Fprint(w, result)

}

func handleSubstract(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'a'", 400)
		return
	}
	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'b'", 400)
		return
	}
	result := a - b
	fmt.Fprint(w, result)
}

func handleMultiply(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'a'", 400)
		return
	}
	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'b'", 400)
		return
	}
	result := a * b
	fmt.Fprint(w, result)
}

func handleDivide(w http.ResponseWriter, r *http.Request) {
	p1 := r.URL.Query().Get("a")
	p2 := r.URL.Query().Get("b")

	a, err := strconv.Atoi(p1)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'a'", 400)
		return
	}
	if a == 0 {
		http.Error(w, "Ошибка в параметре 'a', ноль нельзя делить", 400)
		return
	}
	b, err := strconv.Atoi(p2)
	if err != nil {
		http.Error(w, "Ошибка в параметре 'b'", 400)
		return
	}
	if b == 0 {
		http.Error(w, "Ошибка в параметре 'b', ноль нельзя делить", 400)
		return
	}
	result := a / b
	fmt.Fprint(w, result)
}
