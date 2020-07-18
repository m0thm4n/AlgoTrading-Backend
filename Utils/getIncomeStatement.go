package Utils

import (
	"github.com/m0thm4n/AlgoTrading-Backend/Config"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type IncomeStatement struct {
	Symbol			string	`json:"symbol"`
	Financials		[]struct {
		Date									string	`json:"date"`
		Revenue									string	`json:"Revenue"`
		RevenueGrowth							string	`json:"Revenue Growth"`
		CostOfRevenue							string	`json:"Cost of Revenue"`
		GrossProfit								string	`json:"Gross Profit"`
		RnDExpenses								string	`json:"R&D Expenses"`
		SGnAExpense								string	`json:"SG&A Expense"`
		OperatingExpenses						string	`json:"Operating Expenses"`
		OperatingIncome							string	`json:"Operating Income"`
		InterestExpense							string	`json:"Interest Expense"`
		EarningsBeforeTax						string	`json:"Earnings before Tax"`
		IncomeTaxExpense						string	`json:"Income Tax Expense"`
		NetIncomeNonControllingInt				string	`json:"Net Income - Non-Controlling int"`
		NetIncomeDiscontinuedOps				string	`json:"Net Income - Discontinued ops"`
		NetIncome								string	`json:"Net Income"`
		PreferredDividends						string	`json:"Preferred Dividends"`
		NetIncomeCom							string	`json:"Net Income Com"`
		EPS										string	`json:"EPS"`
		EPSDiluted								string	`json:"EPS Diluted"`
		WeightedAverageShsOut					string	`json:"Weighted Average Shs Out"`
		WeightedAverageShsOutDil				string	`json:"Weighted Average Shs Out (Dil)"`
		DividendPerShare						string	`json:"Dividend per Share"`
		GrossMargin								string	`json:"Gross Margin"`
		EBITDAMargin							string	`json:"EBITDA Margin"`
		EBITMargin								string	`json:"EBIT Margin"`
		ProfitMargin							string	`json:"Profit Margin"`
		FreeCashFlowMargin						string	`json:"Free Cash Flow margin"`
		EBITDA									string	`json:"EBITDA"`
		EBIT									string	`json:"EBIT"`
		ConsolidatedIncome						string	`json:"Consolidated Income"`
		EarningsBeforeTaxMargin					string	`json:"Earnings Before Tax Margin"`
		NetProfitMargin							string	`json:"Net Profit Margin"`
	}	`json:"financials"`
}

var incomeUrl = url.URL{
	Scheme: "https",
	Host:	"financialmodelingprep.com",
	Path:	"/api/v3/financials/income-statement/",
}

func GetIncomeStatement(symbol string) IncomeStatement {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(incomeUrl.String() + symbol + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data IncomeStatement

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyIncomeStatement(symbol string) IncomeStatement {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(incomeUrl.String() + symbol + "?period=quarter" + "&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data IncomeStatement

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetCsvIncomeStatement(symbol string) IncomeStatement {
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

	var data IncomeStatement

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyCsvIncomeStatement(symbol string) IncomeStatement {
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

	var data IncomeStatement

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}