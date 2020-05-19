package Server

import (
	"AlgoTrading-Backend/API"
	"AlgoTrading-Backend/Config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	config := Config.LoadConfiguration("config.json")
	host := config.Host
	port := config.Port

	router := mux.NewRouter().StrictSlash(true)
	// Quote
	router.HandleFunc("/api/quote/{symbol}", API.GetQuote).Methods("GET")
	// Profile
	router.HandleFunc("/api/company/profile/{symbol}", API.GetProfile).Methods("GET")
	// Ticker
	router.HandleFunc("/api/ticker/{query}/{limit}/{exchange}", API.GetTicker).Methods("GET")
	// Income Statements
	router.HandleFunc("/api/financials/income-statement/{symbol}", API.GetIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/income-statement/quarterly/{symbol}", API.GetQuarterlyIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/income-statement/csv/{symbol}", API.GetCsvIncomeStatement).Methods("GET")
	router.HandleFunc("/api/financials/income-statement/quarterly/csv/{symbol}", API.GetQuarterlyCsvIncomeStatement).Methods("GET")
	// Balance Sheet Statement
	router.HandleFunc("/api/financials/balance-sheet-statement/{symbol}", API.GetBalanceSheet).Methods("GET")
	router.HandleFunc("/api/financials/balance-sheet-statement/quarterly/{symbol}", API.GetQuarterlyBalanceSheet).Methods("GET")
	router.HandleFunc("/api/financials/balance-sheet-statement/csv/{symbol}", API.GetCsvBalanceSheet).Methods("GET")
	router.HandleFunc("/api/financials/balance-sheet-statement/quarterly/csv/{symbol}", API.GetQuarterlyCsvBalanceSheet).Methods("GET")
	// Cash Flow Statement
	router.HandleFunc("/api/financials/cash-flow-statement/{symbol}", API.GetCashFlowStatement).Methods("GET")
	router.HandleFunc("/api/financials/cash-flow-statement/quarterly/{symbol}", API.GetQuarterlyCashFlowStatement).Methods("GET")
	router.HandleFunc("/api/financials/cash-flow-statement/csv/{symbol}", API.GetCsvCashFlowStatement).Methods("GET")
	router.HandleFunc("/api/financials/cash-flow-statement/quarterly/csv/{symbol}", API.GetQuarterlyCsvCashFlowStatement).Methods("GET")
	// Financial Ratios
	router.HandleFunc("/api/financial-ratios/{symbol}", API.GetFinancialRatios).Methods("GET")
	// Company Enterprise Value
	router.HandleFunc("/api/financials/company-enterprise-value/{symbol}", API.GetCompanyEnterpriseValue).Methods("GET")
	router.HandleFunc("/api/financials/company-enterprise-value/quarterly/{symbol}",
		API.GetQuarterlyCompanyEnterpriseValue).Methods("GET")
	// Company Key Metrics
	router.HandleFunc("/api/financials/company-key-metrics/{symbol}", API.GetCompanyEnterpriseValue).Methods("GET")
	router.HandleFunc("/api/financials/company-key-metrics/quarterly/{symbol}",
		API.GetQuarterlyCompanyEnterpriseValue).Methods("GET")
	// Company Financial Growth
	router.HandleFunc("/api/financials/company-financial-growth/{symbol}", API.GetCompanyFinancialGrowth).Methods("GET")
	router.HandleFunc("/api/financials/company-financial-growth/quarterly/{symbol}",
		API.GetQuarterlyCompanyFinancialGrowth).Methods("GET")
	// Company Rating
	router.HandleFunc("/api/financials/company-rating/{symbol}", API.GetCompanyRating).Methods("GET")
	// Discounted Cash Flow
	router.HandleFunc("/api/financials/discounted-cash-flow/{symbol}", API.GetDiscountedCashFlow).Methods("GET")
	// Historical Discounted Cash Flow
	router.HandleFunc("/api/financials/historical-discounted-cash-flow/{symbol}",
		API.GetHistoricalDiscountedCashFlow).Methods("GET")
	router.HandleFunc("/api/financials/historical-discounted-cash-flow/{symbol}",
		API.GetQuarterlyHistoricalDiscountedCashFlow).Methods("GET")
	// Stock Symbols List
	router.HandleFunc("/api/financials/stock/list", API.GetSymbolsList).Methods("GET")

	fmt.Printf("Spinning up server on  %s:%s", host, port)
	log.Fatal(http.ListenAndServe(host + ":" + port, router))
}
