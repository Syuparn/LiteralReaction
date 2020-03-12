package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/adj/:id", GetAdj),
		rest.Get("/adverb/:id", GetAdverb),
		rest.Get("/noun/:id", GetNoun),
		rest.Get("/verb/:id", GetVerb),
		rest.Get("/rand/adj", GetRandAdj),
		rest.Get("/rand/adverb", GetRandAdverb),
		rest.Get("/rand/noun", GetRandNoun),
		rest.Get("/rand/verb", GetRandVerb),
		rest.Get("/favs/:page", GetSentences),
		rest.Post("/favs", PostSentence),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":80", api.MakeHandler()))
}
