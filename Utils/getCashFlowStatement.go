package Utils

import (
	"encoding/json"
	"github.com/m0thm4n/AlgoTrading-Backend/Config"
	"log"
	"net/http"
	"net/url"
)

type CashFlow struct {
	Symbol     string `json:"symbol"`
	Financials []struct {
		Date                        string `json:"date"`
		DepreciationNAmortization   string `json:"Depreciation & Amortization"`
		StockBasedCompensation      string `json:"Stock-based compensation"`
		OperatingCashFlow           string `json:"Operating Cash Flow"`
		CapitalExpenditure          string `json:"Capital Expenditure"`
		AcquisitionsAndDisposals    string `json:"Acquisitions and disposals"`
		InvestmentPurchasesAndSales string `json:"Investment purchases and sales"`
		InvestingCashFlow           string `json:"Investing Cash Flow"`
		IssuanceRepaymentOfDebt     string `json:"Issuance (repayment) of debt"`
		IssuanceBuybacksOfShares    string `json:"Issuance (buybacks) of shares"`
		DividendPayments            string `json:"Dividend payments"`
		FinancingCashFlow           string `json:"Financing Cash Flow"`
		EffectOfForexChangesOnCash  string `json:"Effect of forex changes on cash"`
		NetCashFlowChangeInCash     string `json:"Net cash flow / Change in cash"`
		FreeCashFlow                string `json:"Free Cash Flow"`
		NetCashMarketcap            string `json:"Net Cash/Marketcap"`
	} `json:"financials"`
}

var cashFlowUrl = url.URL{
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/financials/cash-flow-statement/",
}

// Series of functions for getting Financial Statements from financialmodelingprep.com
func GetCashFlowStatement(symbol string) CashFlow {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(cashFlowUrl.String() + symbol + "?apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CashFlow

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyCashFlowStatement(symbol string) CashFlow {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(cashFlowUrl.String() + symbol + "?period=quarter&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CashFlow

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetCsvCashFlowStatement(symbol string) CashFlow {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(cashFlowUrl.String() + symbol + "?datatype=csv&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CashFlow

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyCsvCashFlowStatement(symbol string) CashFlow {
	config := Config.LoadConfiguration("config.json")
	key := config.Key

	response, err := http.Get(cashFlowUrl.String() + symbol + "?datatype=csv&period=quarter&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CashFlow

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
