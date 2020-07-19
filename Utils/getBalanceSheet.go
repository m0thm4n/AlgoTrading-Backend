package Utils

import (
	"encoding/json"
	"github.com/m0thm4n/AlgoTrading-Backend/Config"
	"log"
	"net/http"
	"net/url"
)

type BalanceSheet struct {
	Symbol     string `json:"symbol"`
	Financials []struct {
		Date                         string `json:"date"`
		CashAndCashEquivalents       string `json:"Cash and cash equivalents"`
		ShortTermInvestments         string `json:"Short-term investments"`
		CashAndShortTermInvestments  string `json:"Cash and short-term investments"`
		Receivables                  string `json:"Receivables"`
		Inventories                  string `json:"Inventories"`
		TotalCurrentAssets           string `json:"Total current assets"`
		PropertyPlantAndEquipmentNet string `json:"Property, Plant & Equipment Net"`
		GoodwillAndIntangibleAssets  string `json:"Goodwill and Intangible Assets"`
		LongTermInvestments          string `json:"Long-term investments"`
		TaxAssets                    string `json:"Tax assets"`
		Payables                     string `json:"Payables"`
		ShortTermDebt                string `json:"Short-term debt"`
		TotalCurrentLiabilities      string `json:"Total current liabilities"`
		LongTermDebt                 string `json:"Long-term debt"`
		TotalDebt                    string `json:"Total debt"`
		DeferredRevenue              string `json:"Deferred revenue"`
		TaxLiabilities               string `json:"Tax Liabilities"`
		DepositLiabilities           string `json:"Deposit Liabilities"`
		TotalNonCurrentLiabilities   string `json:"Total non-current liabilities"`
		TotalLiabilities             string `json:"Total liabilities"`
		OtherComprehensiveIncome     string `json:"Other comprehensive income"`
		RetainedEarnings             string `json:"Retained  (deficit)"`
		TotalShareholdersEquity      string `json:"Total shareholders equity"`
		Investments                  string `json:"Investments"`
		NetDebt                      string `json:"Net Debt"`
		OtherAssets                  string `json:"Other Assets"`
		OtherLiabilities             string `json:"Other Liabilities"`
	} `json:"financials"`
}

var balanceSheetUrl = url.URL{
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/financials/balance-sheet-statement/",
}

// Series of functions for getting Financial Statements from financialmodelingprep.com
func GetBalanceSheet(symbol string) BalanceSheet {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(balanceSheetUrl.String() + symbol + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data BalanceSheet

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyBalanceSheet(symbol string) BalanceSheet {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(balanceSheetUrl.String() + symbol + "?period=quarter&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data BalanceSheet

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetCsvBalanceSheet(symbol string) BalanceSheet {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(balanceSheetUrl.String() + symbol + "?datatype=csv&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data BalanceSheet

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyCsvBalanceSheet(symbol string) BalanceSheet {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(balanceSheetUrl.String() + symbol + "?datatype=csv&period=quarter&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data BalanceSheet

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
