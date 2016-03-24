package main

import (
	"fmt"
	"net/http"
	a "github.com/djosephsen/answers/lib"
	chicken "github.com/djosephsen/answers/chicken"
	knocknock "github.com/djosephsen/answers/knocknock"
)

func helpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the answer service:\nValid answers:\n\n")
	for name,answer := range a.Answers {
		fmt.Fprintf(w, "/get/%s :: %s\n",name,answer.Desc)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	answerType := r.URL.Path[len("/get/"):]
	fmt.Fprintf(w, "%s",a.Answers[answerType].Rand())
}

func initAnswers(){
	a.Init()
	a.Register(chicken.Answers)
	a.Register(knocknock.Answers)
}
	
func main() {
	initAnswers()
	http.HandleFunc("/", helpHandler)
	http.HandleFunc("/get/", getHandler)
	http.ListenAndServe(":8080", nil)
}
