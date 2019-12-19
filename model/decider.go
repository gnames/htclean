package model

import (
	"fmt"
	"strings"
)

type Decider struct {
	Rows   [][]string
	Title  *Title
	Accept bool
}

func NewDecider(rows [][]string) *Decider {
	d := &Decider{Rows: rows}
	d.Title = NewTitle(rows)
	return d
}

func (d *Decider) Decide() {
	repetition := d.Title.OccurNum / d.Title.NamesNum
	pageCoef := int(float32(d.Title.PagesNum) / 100.0)
	if pageCoef < 1 {
		pageCoef = 1
	}
	fmt.Println(d.Title.ID)
	fmt.Println(d.Title.OccurNum, d.Title.NamesNum)
	fmt.Println(d.Title.OccurNum / d.Title.NamesNum)
	var names []string
	for k := range d.Title.Names {
		names = append(names, k)
	}
	fmt.Println(strings.Join(names, ", "))
	stats := d.Title.Stats
	if stats[MultiHighK] > 0 || (stats[UniK] > 100*pageCoef && repetition < 3) {
		fmt.Println(stats)
		d.Accept = true
		return
	}
}
