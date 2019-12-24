package model

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
	var names []string
	for k := range d.Title.Names {
		names = append(names, k)
	}
	stats := d.Title.Stats
	if stats[MultiHighK] > 0 || (stats[UniK] > 100*pageCoef && repetition < 3) {
		d.Accept = true
		return
	}
}
