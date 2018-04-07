package main

import (
	"fmt"

	"../../core/util"

	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")
	util.CheckError(bow.Open("https://www.americanairlines.es"))

	fm, _ := bow.Form("form#bookingModule")

	fm.Set("REFRESH", "0")
	fm.Set("DIRECT_NON_STOP", "FALSE")
	fm.Set("SEVEN_DAY_SEARCH", "TRUE")
	fm.Set("TRIP_FLOW", "YES")
	fm.Set("SO_SITE_UI_FLIFO_DISP_DOT", "TRUE")
	fm.Set("SO_SITE_AIRLINE_LOC_WAIT", "30000")
	fm.Set("SO_SITE_CORPORATE_ID", "OCG-MUCWW28AA")
	fm.Set("SO_SITE_ALLOW_PROMO", "FALSE")
	fm.Set("SO_SITE_ALLOW_LSA_INDICATOR", "TRUE")
	fm.Set("SO_SITE_DISPLAY_NBR_OF_LSA", "TRUE")
	fm.Set("SO_SITE_ALLOW_FLIFO_OTP_INFO", "TRUE")
	fm.Set("SO_SITE_REFARE_TO_1TICKET", "FALSE")
	fm.Set("BOOKING_FLOW", "REVENUE")
	fm.Set("DATE_RANGE_VALUE_1", "2")
	fm.Set("DATE_RANGE_QUALIFIER_1", "C")
	fm.Set("DATE_RANGE_VALUE_2", "2")
	fm.Set("DATE_RANGE_QUALIFIER_2", "C")
	fm.Set("SITE", "BFJOANEW")
	fm.Set("EXTERNAL_ID", "SPAIN")
	fm.Set("LANGUAGE", "ES")
	fm.Set("SO_SITE_POINT_OF_SALE", "MAD")
	fm.Set("SO_SITE_MARKET_ID", "SPAIN")
	fm.Set("MATRIX_CALENDAR", "FALSE")
	fm.Set("SO_SITE_MOP_PAY_LATER", "FALSE")
	fm.Set("SO_SITE_IS_INSURANCE_ENABLED", "TRUE")
	fm.Set("SO_SITE_MOP_CREDIT_CARD_INS", "TRUE")
	fm.Set("SO_SITE_OFFICE_ID", "MADAA18AB")
	fm.Set("SO_SITE_AMADEUS_OFFICE_ID", "MADAA18AB")
	fm.Set("EMBEDDED_TRANSACTION", "FlexPricerAvailability")
	fm.Set("PRICING_TYPE", "O")
	fm.Set("FP_ERT_ACTIVATED", "TRUE")
	fm.Set("SO_SITE_GET_FP_POSTING_LEVEL", "FALSE")
	fm.Set("B_LOCATION_1", "ATL")
	fm.Set("E_LOCATION_1", "BOS")
	fm.Set("B_LOCATION_2", "BOS")
	fm.Set("E_LOCATION_2", "ATL")
	fm.Set("B_DATE_1", "201803300000")
	fm.Set("B_ANY_TIME_1", "TRUE")
	fm.Set("B_DATE_2", "201804060000")
	fm.Set("B_ANY_TIME_2", "TRUE")
	fm.Set("TRIP_TYPE", "R")
	fm.Set("COMMERCIAL_FARE_FAMILY_1", "LOWEST")
	fm.Set("TRAVELLER_TYPE_1", "ADT")
	fm.Set("DISPLAY_TYPE", "2")
	fm.Set("ARRANGE_BY", "ND")

	util.CheckError(fm.Submit())

	fmt.Println(bow.Body())
}
