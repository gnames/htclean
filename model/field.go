package model

type Field int

const (
	TimeStampF Field = iota
	TitleF
	PageF
	VernacF
	NameF
	StartF
	EndF
	OddsF
	KindOldF
	MatchTypeF
	MatchNameF
	MatchKindF
)

func (f Field) Int() int {
	return int(f)
}
