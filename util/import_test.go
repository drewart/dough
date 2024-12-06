package util

import (
	"strings"
	"testing"
)

func TestCSVImportCents(t *testing.T) {

	csvData := `"Date","Transaction","Name","Memo","Amount"
"2024-08-01","DEBIT","DEBIT PURCHASE -VISA Amazon.com*RF867Amzn.com/bilWA","Download from usbank.com.4738","-15.30"
"2024-08-01","DEBIT","DEBIT PURCHASE -VISA AMAZON MKTPL*RV4Amzn.com/bilWA","Download from usbank.com.4738","-20.75"
"2024-08-01","DEBIT","DEBIT PURCHASE -VISA GFU CAFE        NEWBERG     OR","Download from usbank.com.4738","-3.55"
"2024-08-01","DEBIT","DEBIT PURCHASE -VISA SQ *CHAPTERS LLCNewberg     OR","Download from usbank.com.4738","-5.25"
"2024-08-01","DEBIT","DEBIT PURCHASE -VISA SQ *PULP & CIRCUNewberg     OR","Download from usbank.com.4738","-9.90"
"2024-08-01","DEBIT","DEBIT PURCHASE QFC #5838 14130 KIRKLAND    WA","Download from usbank.com.7765","-95.87"
"2024-08-01","DEBIT","DEBIT PURCHASE QFC #5838 14130 KIRKLAND    WA","Download from usbank.com.7765","-7.16"
"2024-08-01","DEBIT","ELECTRONIC WITHDRAWAL JPMORGAN CHASE","Download from usbank.com.","-2666.76"
"2024-08-01","5735","CHECK","Download from usbank.com.","-200.00"
"2024-08-01","5736","CHECK","Download from usbank.com.","-245.00"`

	var tests = []struct {
		want int
	}{
		{-1530},
		{-2075},
		{-355},
		{-525},
		{-990},
		{-9587},
		{-716},
		{-266675},
		{-20000},
		{-24500},
	}

	reader := strings.NewReader(csvData)

	data, err := ImportCSVToAccount(reader)
	if err != nil {
		t.Error(err)
	}

	if len(tests) > len(data) {
		t.Error()
	}

	for i, item := range data {
		if i >= len(tests) {
			break
		}
		if tests[i].want != item.Amount {

			t.Errorf("payee %s %d != %d", item.Payee, tests[i].want, item.Amount)
			t.Failed()

		}
	}

}
