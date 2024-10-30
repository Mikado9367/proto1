// curl -A "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0" -L ""

package coredb

import (
	"secposretriever/config"
	"secposretriever/internal/tool"

	"encoding/json"
	"errors"
	"log"
	"time"
)

// Handlers for the router yet to be implemented
type RequestPosition struct {
	Bic             string `json:"bic" example:"bicXX123"`
	Isin            string `json:"isin"`
	Account         string `json:"account"`
	Restrictiontype string `json:"restrictiontype"`
	//	Quantity        int64          `json:"quantity"`
	Filter BusinessFilter `json:"filter"`
}

// Struc to return information
type Position struct {
	Isin            string    `json:"isin"`
	Account         string    `json:"account"`
	Restrictiontype string    `json:"restrictiontype"`
	Quantity        int64     `json:"quantity"`
	QuantityFD      string    `json:"quantityfd"`
	LastTimestamp   string    `json:"lasttimestamp"`
	Phase           string    `json:"phase"`
	PartyBic        string    `json:"partybic" example:"bicXX123"`
	BusinessDate    time.Time `json:"businessdate"`
	SystemEntity    string    `json:"systementity"`
	AccountType     string    `json:"accounttype"`
}

// Various error predefined
var errIsinBicRequired = errors.New("isin or bic is required")
var errIsinNotValid = errors.New("isin is not valid")
var errBicNotValid = errors.New("bic is not valid")
var errAccountNotValid = errors.New("account is not valid")
var errRestrictiontypeNotValid = errors.New("restriction type is not valid")
var errDateFilterNotValid = errors.New("date filter is not valid aaaa-mm-dd")
var errPhaseFilterNotValid = errors.New("phase filter is not valid")
var ErrTooMuchRows = errors.New("unexpected data, expected only one position")

func (r RequestPosition) CheckRequest() (err error) {

	err = r.isInputValid()
	err2 := r.isFilterValid()
	err = errors.Join(err, err2)
	return err
}

func (r RequestPosition) isInputValid() (err error) {

	if len(r.Isin) == 0 && len(r.Bic) == 0 {
		errors.Join(err, errIsinBicRequired)
	}
	if !tool.IsAlphaNumeric(r.Bic) && r.Bic != config.Envs.WILDCARD {
		err = errors.Join(err, errBicNotValid)
	}
	if !tool.IsAlphaNumeric(r.Isin) && r.Isin != config.Envs.WILDCARD {
		err = errors.Join(err, errIsinNotValid)
	}
	if !tool.IsAlphaNumeric(r.Account) && r.Account != config.Envs.WILDCARD {
		err = errors.Join(err, errAccountNotValid)
	}
	if !tool.IsAlphaNumeric(r.Restrictiontype) && r.Restrictiontype != config.Envs.WILDCARD {
		err = errors.Join(err, errRestrictiontypeNotValid)
	}

	return err
}

func (r RequestPosition) isFilterValid() (err error) {

	err = r.Filter.IsFilterValid()

	// coding temporaire , juste si on veut sur-ajouter ce que rend la func
	if err != nil {
		err = errors.Join(err, errors.New("oups coding temporaire , juste si on veut sur-ajouter ce que rend la func"))
	}

	return err
}

func (r RequestPosition) GetDataToJson() (data string, err error) {

	// Translate query before asking for Data
	if r.Account == "-" {
		r.Account = ""
	}
	if r.Restrictiontype == "-" {
		r.Restrictiontype = ""
	}
	if r.Bic == "-" {
		r.Bic = ""
	}
	if r.Isin == "-" {
		r.Isin = ""
	}

	rows, err := r.GetPositionsFromDB()

	for i := 0; i < len(rows); i++ {

		tmpPosition := Position(rows[i])
		tmpjson, err := tmpPosition.encodeJson()
		if err != nil {
			log.Fatal("oups json")
			// fmt.Println(tmpPosition)
		}

		if data == "" {
			data = data + tmpjson
		} else {
			data = data + "," + tmpjson
		}
	}

	data = `"result": [` + data + `]`
	return data, err
}

func (a Position) encodeJson() (string, error) {

	if jsonNewValue, err := json.Marshal(a); err != nil {
		return "", err
	} else {

		return string(jsonNewValue), nil
	}
}
