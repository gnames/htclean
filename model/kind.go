package model

type Kind int

const (
	UniK Kind = iota
	MultiLowK
	MultiHighK
)

func NewKind(s string, odds int) Kind {
	if s == "Uninomial" {
		return UniK
	}
	if odds > 700 {
		return MultiHighK
	}
	return MultiLowK
}
