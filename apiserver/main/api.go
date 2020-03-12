package main

import (
	"apiserver/model"
	"net/http"
	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

var dbHandler = model.OpenDB()

func GetAdj(w rest.ResponseWriter, r *rest.Request)    { getWord(w, r, model.Adj) }
func GetAdverb(w rest.ResponseWriter, r *rest.Request) { getWord(w, r, model.Adverb) }
func GetNoun(w rest.ResponseWriter, r *rest.Request)   { getWord(w, r, model.Noun) }
func GetVerb(w rest.ResponseWriter, r *rest.Request)   { getWord(w, r, model.Verb) }

func getWord(w rest.ResponseWriter, r *rest.Request, pos model.POS) {
	id, err := strconv.Atoi(r.PathParam("id"))
	if err != nil {
		rest.NotFound(w, r)
		return
	}
	getWordAt(w, r, pos, id)
}

func GetRandAdj(w rest.ResponseWriter, r *rest.Request)    { getRandWord(w, r, model.Adj) }
func GetRandAdverb(w rest.ResponseWriter, r *rest.Request) { getRandWord(w, r, model.Adverb) }
func GetRandNoun(w rest.ResponseWriter, r *rest.Request)   { getRandWord(w, r, model.Noun) }
func GetRandVerb(w rest.ResponseWriter, r *rest.Request)   { getRandWord(w, r, model.Verb) }

func getRandWord(w rest.ResponseWriter, r *rest.Request, pos model.POS) {
	id := dbHandler.GetRandID(pos)
	getWordAt(w, r, pos, id)
}

func getWordAt(w rest.ResponseWriter, r *rest.Request, pos model.POS, id int) {
	word, err := dbHandler.GetWord(id, pos)
	if err != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(word)
}

func PostSentence(w rest.ResponseWriter, r *rest.Request) {
	sentence := model.FavSentence{}
	if err := r.DecodeJsonPayload(&sentence); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if sentence.FormerWord == "" {
		rest.Error(w, "FormerWord required", 400)
		return
	}
	if sentence.LatterWord == "" {
		rest.Error(w, "LatterWord required", 400)
		return
	}
	dbHandler.StoreSentence(sentence)
	// send back json to tell client that post has successed
	w.WriteJson(&sentence)
}

func GetSentences(w rest.ResponseWriter, r *rest.Request) {
	page, err := strconv.Atoi(r.PathParam("page"))
	if err != nil {
		rest.NotFound(w, r)
		return
	}

	sentences, err := dbHandler.GetSentences(page)
	if err != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(sentences)
}
