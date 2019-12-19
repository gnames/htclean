package model

import (
	"log"
	"strconv"
	"strings"
)

type Title struct {
	ID       string
	Names    map[string]*Name
	Stats    map[Kind]int
	PagesNum int
	NamesNum int
	OccurNum int
}

type Name struct {
	OccurNum int
	Odds     int
	Match    string
	Kind     string
}

func NewTitle(rows [][]string) *Title {
	var err error
	t := &Title{
		ID:    rows[0][TitleF.Int()],
		Names: make(map[string]*Name),
		Stats: map[Kind]int{
			UniK:       0,
			MultiLowK:  0,
			MultiHighK: 0,
		},
	}
	for _, r := range rows {
		nval := r[MatchNameF.Int()]
		if n, ok := t.Names[nval]; ok {
			n.OccurNum += 1
		} else {
			odds, err := strconv.Atoi(r[OddsF.Int()])
			if err != nil {
				log.Fatal(err)
			}
			t.Names[nval] = &Name{
				OccurNum: 1,
				Odds:     odds,
				Match:    r[MatchTypeF.Int()],
				Kind:     r[MatchKindF.Int()],
			}
		}
	}
	if len(rows) == 0 {
		return t
	}
	p := rows[len(rows)-1][PageF.Int()]
	t.PagesNum, err = strconv.Atoi(p)
	if err != nil {
		t.PagesNum = 0
	}
	t.NamesNum = len(t.Names)
	t.getStats()
	return t
}

func (t *Title) getStats() {
	for _, v := range t.Names {
		t.OccurNum += v.OccurNum
		if v.Kind == "Uninomial" {
			t.updateStat(UniK)
			continue
		}
		if (v.Odds <= 0 || v.Odds > 1000000) && !strings.HasPrefix("Fuzzy", v.Match) { //walk around bugs in gnfinder
			t.updateStat(MultiHighK)
			continue
		}
		t.updateStat(MultiLowK)
	}
}

func (t *Title) updateStat(k Kind) {
	if _, ok := t.Stats[k]; ok {
		t.Stats[k] += 1
	} else {
		log.Fatal("Should not happen...")
	}
}
