package model

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const DB_PATH = "./words.sqlite3"

const PAGE_SIZE = 100

func OpenDB() *DBHandler {
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatal(err)
	}
	/* initialize rand for GetRandID */
	rand.Seed(time.Now().UnixNano())

	return &DBHandler{db: db, sizeMap: wordSizeMap(db)}
}

type DBHandler struct {
	db      *sql.DB
	sizeMap map[POS]int
}

func (h *DBHandler) GetWord(id int, pos POS) (*Word, error) {
	row := h.db.QueryRow(
		fmt.Sprintf(`select * from %s where id=?`, posName(pos)+"s"),
		id,
	)

	w := &Word{POS: pos}
	if err := row.Scan(&w.ID, &w.Word); err != nil {
		return nil, err
	}

	return w, nil
}

func (h *DBHandler) StoreSentence(s FavSentence) error {
	_, err := h.db.Exec(
		`insert into favorite_sentences (former_pos, latter_pos, particle, former_word, latter_word) values (?, ?, ?, ?, ?)`,
		posName(s.FormerPOS),
		posName(s.LatterPOS),
		s.Particle,
		s.FormerWord,
		s.LatterWord,
	)
	return err
}

func (h *DBHandler) GetSentences(page int) ([]FavSentence, error) {
	sentences := []FavSentence{}

	minID := (page-1)*PAGE_SIZE + 1
	maxID := page * PAGE_SIZE
	rows, err := h.db.Query(
		fmt.Sprintf(`select * from favorite_sentences where id >= %d and id <= %d`, minID, maxID),
	)
	if err != nil {
		return sentences, err
	}

	for rows.Next() {
		s := FavSentence{}
		var _id int
		var former, latter string
		err := rows.Scan(&_id, &former, &latter, &s.Particle, &s.FormerWord, &s.LatterWord)
		if err != nil {
			log.Fatal(err)
		}
		s.FormerPOS = namePOS(former)
		s.LatterPOS = namePOS(latter)
		sentences = append(sentences, s)
	}

	return sentences, nil
}

func (h *DBHandler) GetRandID(pos POS) int {
	return rand.Intn(h.sizeMap[pos])
}

func wordSizeMap(db *sql.DB) map[POS]int {
	sizeMap := map[POS]int{}

	rows, err := db.Query(`select * from counts`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name string
		var n int
		if err := rows.Scan(&name, &n); err != nil {
			log.Fatal(err)
		}
		// remove plural "s" from table name
		sizeMap[namePOS(name[:len(name)-1])] = n
	}

	return sizeMap
}
