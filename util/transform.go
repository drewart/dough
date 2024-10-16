package util


var Transforms []Transform

type Transform struct {
	ID int
	Name string
	Source string
	Target string
}