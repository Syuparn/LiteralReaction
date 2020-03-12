package model

import (
	"encoding/json"
)

type Word struct {
	ID   int
	Word string
	POS  POS `json:"POS"`
}

type FavSentence struct {
	FormerPos  POS `json:"FormerPos"`
	LatterPos  POS `json:"LatterPos"`
	FormerWord string
	LatterWord string
}

func (s *FavSentence) UnmarshalJSON(data []byte) (err error) {
	var value map[string]string
	json.Unmarshal(data, &value)
	s.FormerPos = namePos(value["FormerPos"])
	s.LatterPos = namePos(value["LatterPos"])
	s.FormerWord = value["FormerWord"]
	s.LatterWord = value["LatterWord"]
	return nil
}

type POS int

const (
	Adj POS = iota
	Adverb
	Noun
	Verb
)

func (pos POS) MarshalJSON() ([]byte, error) {
	return []byte(`"` + posName(pos) + `"`), nil
}

func posName(pos POS) string {
	return map[POS]string{
		Adj:    "adjective",
		Adverb: "adverb",
		Noun:   "noun",
		Verb:   "verb",
	}[pos]
}

func namePos(name string) POS {
	return map[string]POS{
		"adjective": Adj,
		"adverb":    Adverb,
		"noun":      Noun,
		"verb":      Verb,
	}[name]
}
