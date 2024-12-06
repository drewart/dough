package data

import (
	"fmt"
)

var (
	Catagories   []Catagory
	CodeCagories map[string]*Catagory
	IdCagories   map[int]*Catagory
	RootCatagory = Catagory{ID: 0, Code: "00", Name: "Root", Parent: nil, Pos: 0}
	lastID       int
)

type Catagory struct {
	ID     int
	Code   string
	Name   string
	Parent *Catagory
	Tags   []string
	Pos    int
	Active bool
}

func init() {
	IdCagories = make(map[int]*Catagory, 50)
	CodeCagories = make(map[string]*Catagory, 50)
	Catagories = make([]Catagory, 50)

	AddCat(&RootCatagory)

}

func AddCat(cat *Catagory) error {
	c, hasID := IdCagories[cat.ID]
	if hasID {
		err := fmt.Errorf("category %s already has id: %d", c.Name, cat.ID)
		return err
	}
	value, hasCode := CodeCagories[cat.Code]
	if hasCode {
		err := fmt.Errorf("code %s already taken by %s", cat.Code, value.Name)
		return err
	}
	if !hasCode {
		if !hasID {
			IdCagories[cat.ID] = cat
		}

		CodeCagories[cat.Code] = cat
		Catagories = append(Catagories, *cat)
		lastID = cat.ID
	}
	return nil
}

func NewCatagory(id int, name string, code string, parent *Catagory, tags []string, pos int, active bool) (*Catagory, error) {
	if parent == nil {
		parent = &RootCatagory
	}
	cat := &Catagory{
		ID:     id,
		Code:   code,
		Name:   name,
		Parent: parent,
		Tags:   tags,
		Pos:    pos,
		Active: active,
	}

	err := AddCat(cat)
	if err != nil {
		return nil, err
	}

	return cat, nil
}

func GetCatById(ID int) *Catagory {
	if cat, has := IdCagories[ID]; has {
		return cat
	}
	return nil
}

var KeyTermCat map[string]CatagoryMatch

type CatagoryMatch struct {
	ID           int
	Name         string
	Catagory     *Catagory
	IsCheck      bool
	KeyTerm      string
	Terms        []string
	Replace      string
	AmmountMatch string
}
