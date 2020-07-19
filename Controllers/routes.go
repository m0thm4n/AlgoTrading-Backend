package Controllers

import (
	"github.com/m0thm4n/AlgoTrading-Backend/API"
	"github.com/m0thm4n/AlgoTrading-Backend/Middlewares"
)

func (server *Server) initializeRoutes() {
	// config := Config.LoadConfiguration("config.json")
	// host := config.Host
	// port := config.Port

	// router := mux.NewRouter().StrictSlash(true)
	// Quote
	server.Router.HandleFunc("/api/quote/{symbol}", Middlewares.SetMiddlewareJSON(Middlewares.SetMiddlewareAuthentication(API.GetQuote))).Methods("GET")
	// Profile
	server.Router.HandleFunc("/api/company/profile/{symbol}", API.GetProfile).Methods("GET")
	// Ticker
	server.Router.HandleFunc("/api/ticker/{query}/{limit}/{exchange}", API.GetTicker).Methods("GET")
	// Income Statements
	server.Router.HandleFunc("/api/financials/income-statement/{symbol}", API.GetIncomeStatement).Methods("GET")
	server.Router.HandleFunc("/api/financials/income-statement/quarterly/{symbol}", API.GetQuarterlyIncomeStatement).Methods("GET")
	server.Router.HandleFunc("/api/financials/income-statement/csv/{symbol}", API.GetCsvIncomeStatement).Methods("GET")
	server.Router.HandleFunc("/api/financials/income-statement/quarterly/csv/{symbol}", API.GetQuarterlyCsvIncomeStatement).Methods("GET")
	// Balance Sheet Statement
	server.Router.HandleFunc("/api/financials/balance-sheet-statement/{symbol}", API.GetBalanceSheet).Methods("GET")
	server.Router.HandleFunc("/api/financials/balance-sheet-statement/quarterly/{symbol}", API.GetQuarterlyBalanceSheet).Methods("GET")
	server.Router.HandleFunc("/api/financials/balance-sheet-statement/csv/{symbol}", API.GetCsvBalanceSheet).Methods("GET")
	server.Router.HandleFunc("/api/financials/balance-sheet-statement/quarterly/csv/{symbol}", API.GetQuarterlyCsvBalanceSheet).Methods("GET")
	// Cash Flow Statement
	server.Router.HandleFunc("/api/financials/cash-flow-statement/{symbol}", API.GetCashFlowStatement).Methods("GET")
	server.Router.HandleFunc("/api/financials/cash-flow-statement/quarterly/{symbol}", API.GetQuarterlyCashFlowStatement).Methods("GET")
	server.Router.HandleFunc("/api/financials/cash-flow-statement/csv/{symbol}", API.GetCsvCashFlowStatement).Methods("GET")
	server.Router.HandleFunc("/api/financials/cash-flow-statement/quarterly/csv/{symbol}", API.GetQuarterlyCsvCashFlowStatement).Methods("GET")
	// Financial Ratios
	server.Router.HandleFunc("/api/financial-ratios/{symbol}", API.GetFinancialRatios).Methods("GET")
	// Company Enterprise Value
	server.Router.HandleFunc("/api/financials/company-enterprise-value/{symbol}", API.GetCompanyEnterpriseValue).Methods("GET")
	server.Router.HandleFunc("/api/financials/company-enterprise-value/quarterly/{symbol}",
		API.GetQuarterlyCompanyEnterpriseValue).Methods("GET")
	// Company Key Metrics
	server.Router.HandleFunc("/api/financials/company-key-metrics/{symbol}", API.GetCompanyEnterpriseValue).Methods("GET")
	server.Router.HandleFunc("/api/financials/company-key-metrics/quarterly/{symbol}",
		API.GetQuarterlyCompanyEnterpriseValue).Methods("GET")
	// Company Financial Growth
	server.Router.HandleFunc("/api/financials/company-financial-growth/{symbol}", API.GetCompanyFinancialGrowth).Methods("GET")
	server.Router.HandleFunc("/api/financials/company-financial-growth/quarterly/{symbol}",
		API.GetQuarterlyCompanyFinancialGrowth).Methods("GET")
	// Company Rating
	server.Router.HandleFunc("/api/financials/company-rating/{symbol}", API.GetCompanyRating).Methods("GET")
	// Discounted Cash Flow
	server.Router.HandleFunc("/api/financials/discounted-cash-flow/{symbol}", API.GetDiscountedCashFlow).Methods("GET")
	// Historical Discounted Cash Flow
	server.Router.HandleFunc("/api/financials/historical-discounted-cash-flow/{symbol}",
		API.GetHistoricalDiscountedCashFlow).Methods("GET")
	server.Router.HandleFunc("/api/financials/historical-discounted-cash-flow/{symbol}",
		API.GetQuarterlyHistoricalDiscountedCashFlow).Methods("GET")
	// Stock Symbols List
	server.Router.HandleFunc("/api/financials/stock/list", API.GetSymbolsList).Methods("GET")
	// All Real-time Stock Price
	server.Router.HandleFunc("/api/financials/real-time-stock-prices/", API.GetRealTimePrices).Methods("GET")
	// Real-time Stock Price for Company
	server.Router.HandleFunc("/api/financials/real-time-stock-price/{symbol}", API.GetRealTimePriceForCompany).Methods("GET")
	// Calls to Historical Stock Prices / Daily Historical Stock Prices
	server.Router.HandleFunc("/api/financials/historical-stock-prices/{time}/{symbol}",
		API.GetHistoricalStockPriceByMinute).Methods("GET")
	server.Router.HandleFunc("/api/financials/historical-stock-prices-hour/{symbol}",
		API.GetHistoricalStockPriceByHour).Methods("GET")
	server.Router.HandleFunc("/api/financials/full-historical-stock-prices/{symbol}",
		API.GetDailyHistoricalStockPrice).Methods("GET")
	server.Router.HandleFunc("/api/financials/major-indexes/",
		API.GetMajorIndexes).Methods("GET")

	// fmt.Printf("Spinning up server on  %s:%s", host, port)
	// log.Fatal(http.ListenAndServe(host+":"+port, server.Router))
}

