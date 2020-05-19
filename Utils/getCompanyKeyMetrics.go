package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type CompanyKeyMetrics struct {
	Symbol	string	`json:"symbol"`
	Metrics	[]struct {
		Date							string	`json:"date"`
		RevenuePerShare					string	`json:"Revenue per Share"`
		NetIncomePerShare				string	`json:"Net Income per Share"`
		OperatingCashFlowPerShare		string	`json:"Operating Cash Flow per Share"`
		FreeCashFlowPerShare			string	`json:"Free Cash Flow Per Share"`
		CashPerShare					string	`json:"Cash per Share"`
		BookValuePerShare				string	`json:"Book Value per Share"`
		TangibleBookValuePerShare		string	`json:"Tangible Book Value per Share"`
		ShareholdersEquityPerShare		string	`json:"Shareholders Equity per Share"`
		InterestDebtPerShare			string	`json:"Interest Debt per Share"`
		MarketCap						string	`json:"Market Cap"`
		EnterpriseValue					string	`json:"Enterprise Value"`
		PERatio							string	`json:"PE ratio"`
		PriceToSalesRatio				string	`json:"Price to Sales Ratio"`
		POCFRatio						string	`json:"POCF ratio"`
		PFCFRatio						string	`json:"PFCF ratio"`
		PBRatio							string	`json:"PB ratio"`
		PTBRatio						string	`json:"PTB ratio"`
		EVToSales						string	`json:"EV to Sales"`
		EnterpriseValueOverEBITDA		string	`json:"Enterprise Value over EBITDA"`
		EVToOperatingCashFlow			string	`json:"EV to Operating cash flow"`
		EVToFreeCashFlow				string	`json:"EV to Free cash flow"`
		EarningsYield					string	`json:"Earnings Yield"`
		FreeCashFlowYield				string	`json:"Free Cash Flow Yield"`
		DebtToEquity					string	`json:"Debt to Equity"`
		DebtToAssets					string	`json:"Debt to Assets"`
		NetDebtToEBITDA					string	`json:"Net Debt to EBITDA"`
		CurrentRatio					string	`json:"Current ratio"`
		InterestCoverage				string	`json:"Interest Coverage"`
		IncomeQuality					string	`json:"Income Quality"`
		DividendYield					string	`json:"Dividend Yield"`
		PayoutRatio						string	`json:"Payout Ratio"`
		SGnAToRevenue					string	`json:"SG&A to Revenue"`
		RnDToRevenue					string	`json:"R&D to Revenue"`
		IntangiblesToTotalAssets		string	`json:"Intangibles to Total Assets"`
		CapexToOperatingCashFlow		string	`json:"Capex to Operating Cash Flow"`
		CapexToRevenue					string	`json:"Capex to Revenue"`
		CapexToDepreciation				string	`json:"Capex to Depreciation"`
		StockBasedCompensationToRevenue	string	`json:"Stock-based compensation to Revenue"`
		GrahamNumber					string	`json:"Graham Number"`
		ROIC							string	`json:"ROIC"`
		ReturnOnTangibleAssets			string	`json:"Return on Tangible Assets"`
		GrahamNetNet					string	`json:"Graham Net-Net"`
		WorkingCapital					string	`json:"Working Capital"`
		TangibleAssetValue				string	`json:"Tangible Asset Value"`
		NetCurrentAssetValue			string	`json:"Net Current Asset Value"`
		InvestedCapital					string	`json:"Invested Capital"`
		AverageReceivables				string	`json:"Average Receivables"`
		AveragePayables					string	`json:"Average Payables"`
		AverageInventory				string	`json:"Average Inventory"`
		DaysSalesOutstanding			string	`json:"Days Sales Outstanding"`
		DaysPayablesOutstanding			string	`json:"Days Payables Outstanding"`
		DaysOfInventoryOnHand			string	`json:"Days of Inventory on Hand"`
		ReceivablesTurnover				string	`json:"Receivables Turnover"`
		PayablesTurnover				string	`json:"Payables Turnover"`
		InventoryTurnover				string	`json:"Inventory Turnover"`
		ROE								string	`json:"ROE"`
		CapexPerShare					string	`json:"Capex per Share"`
	}
}

var companyKeyMetricsUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/company-key-metrics/",
}

// Series of functions for getting Company Enterprise Value from financialmodelingprep.com
func GetCompanyKeyMetrics(symbol string) CompanyKeyMetrics {
	response, err := http.Get(companyKeyMetricsUrl.String() + symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyKeyMetrics

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func GetQuarterlyCompanyKeyMetrics(symbol string) CompanyKeyMetrics {
	response, err := http.Get(companyKeyMetricsUrl.String() + symbol + "?period=quarter")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data CompanyKeyMetrics

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}