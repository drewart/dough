package data

import (
	"fmt"
)

var Catagories []Catagory

var CodeCagories map[string]*Catagory

var IdCagories map[int]*Catagory

var RootCatagory = Catagory{ID: 0, Code: "00", Name: "Root", Parent: nil, Order: 0}

var lastID = 0

type Catagory struct {
	ID     int
	Code   string
	Name   string
	Parent *Catagory
	Tags   []string
	Pos  int
}

func NewCatagory(id int, name string, code string, parent *Catagory, tags []string, pos int) (*Catagory, error) {
	if parent == nil {
		parent = &RootCatagory
	}
	c, hasID := IdCagories[id] 
	if hasID {
		err := fmt.Errorf("category %s already has id: %d", c.Name, id)
		return nil, err
	}
	value, hasCode := CodeCagories[code]
	if hasCode {
		err := fmt.Errorf("code %s already taken by %s", code, value.Name)
		return nil, err
	}
	cat := &Catagory{
		ID:     id,
		Code:   code,
		Name:   name,
		Parent: parent,
		Tags:   tags,
		Pos:  pos,
	}

	if !hasCode {
		if !hasID {
			IdCagories[cat.ID] = cat
		}

		CodeCagories[code] = cat
		Catagories = append(Catagories, *cat)
	}

	return cat, nil
}

func GetCatById(ID int) *Catagory {
	if cat, has := IdCagories[ID]; has {
		return cat
	}
	return nil
}

type CatagoryMatch struct {
	ID           int
	Catagory     *Catagory
	IsCheck      bool
	KeyTerm      string
	Terms        []string
	AmmountMatch int
}
