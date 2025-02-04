package models

type Grade struct {
	ActivityCode string
	Score        int
}

type SchoolType int

const (
	Primary SchoolType = iota
	Secondary
)
