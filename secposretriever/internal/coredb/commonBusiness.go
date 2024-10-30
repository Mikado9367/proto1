package coredb

import (
	"errors"
	"slices"

	"secposretriever/internal/tool"
)

// Handlers for the router yet to be implemented
type BusinessObject interface {
	Filter() BusinessFilter
	New()
	IsInputValid() (err error)
	RetrieeData(*BusinessObject)
}

type BusinessFilter struct {
	Date  string
	Phase string
}

// func (b BusinessFilter) IsFilterValid() (bool bool, err error) {
func (b BusinessFilter) IsFilterValid() (err error) {

	var tmp = []string{"SOD", "NTS", "RTS", "LAST"}
	if !slices.Contains(tmp, b.Phase) {
		err = errors.Join(err, errPhaseFilterNotValid)
	}

	if !tool.IsDateValue(b.Date) {
		err = errors.Join(err, errDateFilterNotValid)
	}

	return err
}
