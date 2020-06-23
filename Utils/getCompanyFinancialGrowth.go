package Utils

import (
	"AlgoTrading-Backend/Config"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type CompanyFinancialGrowth struct {
	Symbol string `json:"symbol"`
	Growth []struct {
		Date                                      string `json:"date"`
		GrossProfitGrowth                         string `json:"Gross Profit Growth"`
		EBITGrowth                                string `json:"EBIT Growth"`
		OperatingIncomeGrowth                     string `json:"Operating Income Growth"`
		EPSGrowth                                 string `json:"EPS Growth"`
		EPSDilutedGrowth                          string `json:"EPS Diluted Growth"`
		WeightedAverageSharesGrowth               string `json:"Weighted Average Shares Growth"`
		WeightedAverageSharesDilutedGrowth        string `json:"Weighted Average Shares Diluted Growth"`
		DividendsPerShareGrowth                   string `json:"Dividends per Share Growth"`
		OperatingCashFlowGrowth                   string `json:"Operating Cash Flow growth"`
		FreeCashFlowGrowth                        string `json:"Free Cash Flow growth"`
		TenYearRevenueGrowthPerShare              string `json:"10Y Revenue Growth (per Share)"`
		FiveYearRevenueGrowthPerShare             string `json:"5Y Revenue Growth (per Share)"`
		ThreeYearRevenueGrowthPerShare            string `json:"3Y Revenue Growth (per Share)"`
		TenYearOperatingCFGrowthPerShare          string `json:"10Y Operating CF Growth (per Share)"`
		FiveYearOperatingCFGrowthPerShare         string `json:"5Y Operating CF Growth (per Share)"`
		ThreeYearOperatingCFGrowthPerShare        string `json:"3Y Operating CF Growth (per Share)"`
		TenYearNetIncomeGrowthPerShare            string `json:"10Y Net Income Growth (per Share)"`
		FiveYearNetIncomeGrowthPerShare           string `json:"5Y Net Income Growth (per Share)"`
		ThreeYearNetIncomeGrowthPerShare          string `json:"3Y Net Income Growth (per Share)"`
		TenYearShareholdersEquityGrowthPerShare   string `json:"10Y Shareholders Equity Growth (per Share)"`
		FiveYearShareholdersEquityGrowthPerShare  string `json:"5Y Shareholders Equity Growth (per Share)"`
		ThreeYearShareholdersEquityGrowthPerShare string `json:"3Y Shareholders Equity Growth (per Share)"`
		TenYearDividendPerShareGrowthPerShare     string `json:"10Y Dividend per Share Growth (per Share)"`
		FiveYearDividendPerShareGrowthPerShare    string `json:"5Y Dividend per Share Growth (per Share)"`
		ThreeYearDividendPerShareGrowthPerShare   string `json:"3Y Dividend per Share Growth (per Share)"`
		ReceivablesGrowth                         string `json:"Receivables growth"`
		InventoryGrowth                           string `json:"Inventory Growth"`
		AssetGrowth                               string `json:"Asset Growth"`
		BookValuePerShareGrowth                   string `json:"Book Value per Share Growth"`
		DebtGrowth                                string `json:"Debt Growth"`
		RAndDExpenseGrowth                        string `json:"R&D Expense Growth"`
		SGnAExpensesGrowth                        string `json:"SG&A Expenses Growth"`
	}	`json:"growth"`
}

var companyFinancialGrowth = url.URL {
	Scheme:	"https",
	Host:	"financialmodelingprep.com",
	Path:	"/api/v3/financials/financial-statement-growth/",
}

// Series of functions for getting Company Financial Growth from financialmodelingprep.com
func GetCompanyFinancialGrowth(symbol string) CompanyFinancialGrowth {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(companyFinancialGrowth.String() + symbol + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyFinancialGrowth

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyCompanyFinancialGrowth(symbol string) CompanyFinancialGrowth {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(companyFinancialGrowth.String() + symbol + "?period=quarter&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyFinancialGrowth

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}