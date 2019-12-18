package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type title struct {
	id    string
	pages []page
	kinds map[string]int
}

type page struct {
	pageNum int
	names   []name
}

type name struct {
	nameVerbatim string
	name         string
	offsetStart  int
	offsetEnd    int
	odds         int
	kind         string
}

func main() {
	var t *title
	var p page
	f, err := os.Open("filtered.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(f)

	for {
		l, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		id, pgNum, verbatim, name, offsetStart, offsetEnd, odds, kind := l[1],
			l[2], l[3], l[4], l[5], l[6], l[7], l[8]

		if t == nil {
			pgNum, err := strconv.Atoi(pgNum)
			if err != nil {
				log.Fatal(err)
			}
			p = page{pageNum: pgNum}
			t = &title{id: id}
		}
		_ = p
		_ = verbatim
		_ = name
		_ = offsetStart
		_ = offsetEnd
		_ = odds
		_ = kind
	}
}
