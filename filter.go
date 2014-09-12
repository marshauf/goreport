package goreport

type Filter interface {
	Filter(Entry) bool
}
