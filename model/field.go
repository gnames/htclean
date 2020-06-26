package model

type Field int

// the constants are according results from May 2020
// TimeStamp,ID,PageID,Verbatim,NameString,AnnotNomen,OffsetStart,OffsetEnd,Odds,Kind
// MatchType, MatchName, Cardinality
const (
	TimeStampF Field = iota
	TitleF
	PageF
	VerbatF
	NameF
	AnnotNomenF
	StartF
	EndF
	OddsF
	KindOldF
	MatchTypeF
	MatchNameF
	MatchCardinF
)

func (f Field) Int() int {
	return int(f)
}
