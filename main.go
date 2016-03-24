package main

import (
	"fmt"
	"net/http"
	"time"

	a "github.com/djosephsen/answers/lib"
	chicken "github.com/djosephsen/answers/chicken"
	knocknock "github.com/djosephsen/answers/knocknock"
	metrics "github.com/djosephsen/answers/metrics"
)

func helpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the answer service:\nValid answers:\n\n")
	defer metrics.Time("answer.handler.help", time.Now())
	for name,answer := range a.Answers {
		fmt.Fprintf(w, "/get/%s :: %s\n",name,answer.Desc)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	answerType := r.URL.Path[len("/get/"):]
	defer metrics.Time("answer.handler."+answerType, time.Now())
	fmt.Fprintf(w, "%s",a.Answers[answerType].Rand())
}

func initAnswers(){
	a.Init()
	a.Register(chicken.Answers)
	a.Register(knocknock.Answers)
}
	
func main() {
	initAnswers()
	metrics.Connect()
	http.HandleFunc("/", helpHandler)
	http.HandleFunc("/get/", getHandler)
	http.ListenAndServe(":8080", nil)
}
