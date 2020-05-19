package Utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type FinancialRatios struct {
	Symbol	string	`json:"symbol"`
	Ratios	[]struct {
		Date						string	`json:"date"`
		InvestmentValuationRatios	[]struct {
			PriceBookValueRatio				string	`json:"priceBookValueRatio"`
			PriceToBookRatio				string	`json:"priceToBookRatio"`
			PriceToSalesRatio				string	`json:"priceToSalesRatio"`
			PriceEarningsRatio				string	`json:"priceEarningsRatio"`
			ReceivablesTurnover				string	`json:"receivablesTurnover"`
			PriceToFreeCashFlowRatio		string	`json:"priceToFreeCashFlowRatio"`
			PriceToOperatingCashFlowsRatio	string	`json:"priceToOperatingCashFlowsRatio"`
			PriceCashFlowRatio				string	`json:"priceCashFlowRatio"`
			PriceEarningsToGrowthRatio		string	`json:"priceEarningsToGrowthRatio"`
			PriceSalesRatio					string	`json:"priceSalesRatio"`
			DividendYield					string	`json:"dividendYield"`
			PriceFairValue					string	`json:"priceFairValue"`
		}
		ProfitabilityIndicatorRatios []struct {
			NiPerEBT				string	`json:"niperEBT"`
			EBTPerEBIT				string	`json:"ebtperEBIT"`
			EbitPerRevenue			string	`json:"ebitperRevenue"`
			GrossProfitRevenue		string	`json:"grossProfitRevenue"`
			OperatingProfitMargin	string	`json:"operatingProfitMargin"`
			PretaxProfitMargin		string	`json:"pretaxProfitMargin"`
			NetProfitMargin			string	`json:"netProfitMargin"`
			EffectiveTaxRate		string	`json:"effectiveTaxRate"`
			ReturnOnAssets			string	`json:"returnOnAssets"`
			ReturnOnEquity			string	`json:"returnOnEquity"`
			ReturnOnCapitalEmployed	string	`json:"returnOnCapitalEmployed"`
			NIPerEBT				string	`json:"nIperEBT"`
			EBTPerEBITFR				string	`json:"eBTperEBIT"`
			EBITPerRevenue			string	`json:"eBITperRevenue"`
		}
		OperatingPerformanceRatios	[]struct {
			ReceivablesTurnover		string	`json:"receivablesTurnover"`
			PayablesTurnover		string	`json:"payablesTurnover"`
			InventoryTurnover		string	`json:"inventoryTurnover"`
			FixedAssetTurnover		string	`json:"fixedAssetTurnover"`
			AssetTurnover			string	`json:"assetTurnover"`
		}
		LiquidityMeasurementRatios []struct {
			CurrentRatio					string	`json:"currentRatio"`
			QuickRatio						string	`json:"quickRatio"`
			CashRatio						string	`json:"cashRatio"`
			DaysOfSalesOutstanding			string	`json:"daysOfSalesOutstanding"`
			DaysOfInventoryOutstanding		string	`json:"daysOfInventoryOutstanding"`
			OperatingCycle					string	`json:"operatingCycle"`
			DaysOfPayablesOutstanding		string	`json:"daysOfPayablesOutstanding"`
			CashConversionCycle				string	`json:"cashConversionCycle"`
		}
		DebtRatios []struct {
			DebtRatio						string	`json:"debtRatio"`
			DebtEquityRatio					string	`json:"debtEquityRatio"`
			LongTermDebtToCapitalization	string	`json:"longtermDebtToCapitalization"`
			TotalDebtToCapitalization		string	`json:"totalDebtToCapitalization"`
			InterestCoverage				string	`json:"interestCoverage"`
			CashFlowToDebtRatio				string	`json:"cashFlowToDebtRatio"`
			CompanyEquityMultiplier			string	`json:"companyEquityMultiplier"`
		}
		CashFlowIndicatorRatios	[]struct {
			OperatingCashFlowPerShare			string	`json:"operatingCashFlowPerShare"`
			FreeCashFlowPerShare				string	`json:"freeCashFlowPerShare"`
			CashPerShare						string	`json:"cashPerShare"`
			PayoutRatio							string	`json:"payoutRatio"`
			ReceivableTurnover					string	`json:"receivableTurnover"`
			OperatingCashFlowSalesRatio			string	`json:"operatingCashFlowSalesRatio"`
			FreeCashFlowOperatingCashFlowRatio	string	`json:"freeCashFlowOperatingCashFlowRatio"`
			CashFlowCoverageRatios				string	`json:"cashFlowCoverageRatios"`
			ShortTermCoverageRatios				string	`json:"shortTermCoverageRatios"`
			CapitalExpenditureCoverageRatios	string	`json:"capitalExpenditureCoverageRatios"`
			DividendPaidAndCapexCoverageRatios	string	`json:"dividendPaidAndCapexCoverageRatios"`
			DividendPayoutRatio					string	`json:"dividendPayoutRatio"`
		}
	}
}

var financialRatiosUrl = url.URL {
	Scheme: "https",
	Host:   "financialmodelingprep.com",
	Path:   "/api/v3/financial-ratios/",
}

// Gets profile from financialmodelingprep.com
func GetFinancialRatios(symbol string) FinancialRatios {
	response, err := http.Get(financialRatiosUrl.String() + symbol)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data FinancialRatios

	// err = json.NewDecoder(response.Body).Decode(&data)
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	return data

	// return data
}
