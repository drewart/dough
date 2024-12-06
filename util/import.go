package util

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/drewart/dough/data"
)

func DollorStrToCents(dollarStr string) int {
	ammount := 0
	amtParts := strings.SplitN(dollarStr, ".", 2)
	if len(amtParts) > 1 {
		dollar, err := strconv.Atoi(amtParts[0])
		if err != nil {
			log.Fatalf("dollar amount parse %s %s", dollarStr, err)
		}
		cents, err := strconv.Atoi(amtParts[1])
		if err != nil {
			log.Fatalf("cents format error %s %s", dollarStr, err)
		}
		if dollar < 0 {
			ammount = (dollar * 100) - cents
		} else {
			ammount = (dollar * 100) + cents
		}
	}
	return ammount
}

// TODO change to LegerItem
func ImportCSVToAccount(reader io.Reader) ([]data.LedgerEntry, error) {
	var entries []data.LedgerEntry

	csvReader := *csv.NewReader(reader)

	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("unable to parse file as CSV \n %s", err)
		return entries, err
	}

	// Date,Transaction,Payee,Memo,Amount
	//entries = make([]data.LedgerEntry, len(rows))
	transformers := GetTransformers()

	for _, row := range rows {

		date := row[0]
		trans := row[1]
		payee := row[2]
		memo := row[3]
		amountStr := row[4]
		check := ""

		// skip Header
		if date == "Date" {
			continue
		}
		transDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Fatalf("Date format error %s %s", date, err)
		}

		//lower case
		payee = strings.ToLower(payee)
		memo = strings.ToLower(memo)
		// CUSTOM
		if payee == "check" {
			check = trans
			trans = "check"
		}

		for _, tr := range transformers.Transforms {
			if tr.Field == "payee" {
				payee = tr.FindReplace(payee)
			} else if tr.Field == "memo" {
				memo = tr.FindReplace(memo)
			}
		}

		amount := DollorStrToCents(amountStr)

		entries = append(entries, data.LedgerEntry{Date: transDate,
			Account:   nil,
			TransType: trans,
			Check:     check,
			Payee:     payee,
			Memo:      memo,
			Amount:    amount,
			RawRecord: strings.Join(row, ",")})
	}

	return entries, nil
}

func ImportCatagories(reader io.Reader) {
	csvReader := *csv.NewReader(reader)

	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Printf("unable to parse file as CSV for \n %s", err)
	}
	store := data.NewDoughStorage()
	defer store.Close()
	for _, row := range rows {
		if row[0] == "id" {
			continue
		}
		idStr := row[0]
		code := row[1]
		name := row[2]
		parent_id := row[3]
		tagRaw := row[4]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatal(err)
		}

		tags := strings.Split(tagRaw, "|")
		var parentCat *data.Catagory
		if parent_id != "" {
			p_id, err := strconv.Atoi(parent_id)
			if err != nil {
				log.Fatal(err)
			}
			parentCat = data.GetCatById(p_id)
		}
		active := true
		cat, err := data.NewCatagory(id, name, code, parentCat, tags, id, active)
		if err != nil {
			log.Println(err)
		} else {
			store.InsertCatagory(*cat)
		}
	}

}
