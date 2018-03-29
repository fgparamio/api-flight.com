package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	checkError(bow.Open("https://www.spirit.com/Default.aspx"))

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8")
	bow.AddRequestHeader("Accept_Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Referer", "https://www.easyjet.com/es/")
	bow.AddRequestHeader("Host", "www.spirit.com")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")
	bow.AddRequestHeader("Origin", "https://www.spirit.com")
	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	bow.AddRequestHeader("Referer", "https://www.spirit.com/Default.aspx")

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form#book-travel-form")
	fm.Set("bypassHC", "False")
	fm.Set("bookingType", "F")
	fm.Set("carPickUpTime", "16")
	fm.Set("carDropOffTime", "16")
	fm.Set("tripType", "roundTrip")
	fm.Set("vacationPackageType", "on")
	fm.Set("from", "ATL")
	fm.Set("to", "BOS")
	fm.Set("departDate", "05/02/2018")
	fm.Set("departDateDisplay", "05/02/2018")
	fm.Set("returnDateDisplay", "05/07/2018")
	fm.Set("input_data_to", "03/24/2018")
	fm.Set("ADT", "1")
	fm.Set("CHD", "0")
	fm.Set("INF", "0")
	fm.Set("redeemMiles", "false")

	checkError(fm.Submit())

	bow.Open("https://36d71176.akstat.io/")

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xm…plication/xml;q=0.9,*/*;q=0.8")
	bow.AddRequestHeader("Host", "www.spirit.com")
	bow.AddRequestHeader("Accept-Language", "es-US,es-419;q=0.9,es;q=0.8")
	bow.AddRequestHeader("Referer", "https://www.spirit.com/Default.aspx")
	bow.AddRequestHeader("Cache-Control", "max-age=0")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Cookie", "_qubitTracker=1521334099939.545587; _ga=GA1.2.846257820.1521334100; _up=1.2.491268939.1521334100; _mibhv=anon-1521334100583-1371765294_7060; _ctz_variation=[{%22experiment%22:%22placebo-1:99-08104166402ef58add75d5439bea7728%22%2C%22publisherID%22:2486%2C%22variation%22:1%2C%22variationCacheKey%22:%226ecb5a21d721b4ca26b832b1ea3060d7%22%2C%22realIP%22:%22181.48.153.214%22%2C%22weight%22:0.99}%2C{%22experiment%22:%22optSetWidgetEndpoints-90:10-233f1ae31eade7741609637d0c8d489b%22%2C%22publisherID%22:2486%2C%22variation%22:0%2C%22variationCacheKey%22:%220b04fcabd4881c30cf4919cd900859dc%22%2C%22realIP%22:%22181.48.153.214%22%2C%22weight%22:0.1}]; ASP.NET_SessionId=r4e2ow0jwuysi4knw5grymw1; _gid=GA1.2.1154807514.1522333637; _qst_s=2; _qst=%5B2%2C0%5D; InsurancePopup=false; t011-pageView=2; __qca=P0-38095577-1522334111880; xyz_cr_348_et_100==NaN&cr=348&et=100&ap=; _gaexp=GAX1.2.TDjc2huxS-Gh-IqmPhQlcg.17684.0; _qsst_s=1522337451695; x_qtag_spiritairlinesus=EYDPPCalendarMarket.aspx*1522337205666*Vspirit.com*Idirect@*a*Qsc*Q*j2*C*B4*C*P13*5-@1-*4-es.wikipedia.org/*Y*9-*@2-/*kculture*+es-PR*Y*A1521334098664*b*C*a*I*Xgithub.com/fgparamio/api-flight.com*Y*9-*@2-/*Y*A1522333637920*b*C*a*@3-*Y*9-*@2-/*Y*A1522335222364*b*C*a*@3-*Y*9-*@2-/@0-*Y*A@1-*b*E*C*F*Q*@2-/@0-*Y*Q__v*z; SPIRIT_MKTGsessionId=39a5d03e-d1a5-4b2d-b2f3-501461a83469; _qPageNum_spirit_airlines_us=11; _qsst=1522337452494; _dc_gtm_UA-206943-1=1; _gat_UA-206943-1=1; t002=2; qb_permanent=1521334099939.545587:13:12:2:2:0::0:1:0:Barbdh:BavQau:::::181.48.153.214:medellin:10498:colombia:CO:6.29475:-75.553:unknown:unknown:antioquia:11641:53_C&4_C&5_C&6_C&7_C&8_C&CW_C&Ce_C&BV_C&Cf_C&CL!&BE_C&B!&CP_C&B5!&CQ_C&Be_C&CM!&Cy!&C4!&Bf_C&CZ_C&C0!&Bh_C&Bb_C&C/!&Bc_C&CN_C&C1!&Cz!&C5!&C6_C&C7_C&Bd_C&BI_C&CJ!&CK!&Bi_C&Bj_C&CO_C&B7!&B+)0&BZ_C&C!&DJ!&DK!&DA!&Cu)1&B8@N&T@N&B/@N&CR@N&Cl@N&B9@N&Ca@N&Dk@N&Ch@N&Cq@N&Cr@N&Cn@N&CT@N&CA@N&V@N&Cg@N&CY@N&Dj@N:CNm4=G2=G=91i=Ca&CiWT=Ls=B=BFqc=OW&Cq0n=Jt=C=BMHU=Ju&CwWX=CN=M=BMav=Nk::WJyYhp+:WJyJ+Ra:0:0:0::0:0:.spirit.com; qb_session=12:1:152:CNm4=F&CiWT=B&Cq0n=C&CwWX=M:0:WJyJ+Ra:0:0:0:0:.spirit.com; _gali=book-travel-form; FlightSearch=%7B%22BookingType%22%3A%22F%22%2C%22From%22%3A%22ATL%22%2C%22To%22%3A%22BOS%22%2C%22HotelOnlyAutoComplete%22%3A%22%22%2C%22HotelOnlyAutoCompleteHidden%22%3A%22%22%2C%22TripType%22%3A%22roundTrip%22%2C%22DepartureDate%22%3A%2204/06/2018%22%2C%22ReturnDate%22%3A%2204/13/2018%22%2C%22Adults%22%3A%221%22%2C%22Children%22%3A%220%22%2C%22PromoCode%22%3A%22%22%2C%22HotelRooms%22%3A%221%22%2C%22Language%22%3A%22en-US%22%2C%22RedeemMiles%22%3A%22false%22%7D; RT=\"sl=12&ss=1522333633575&tt=79574&obo=0&bcn=%2F%2F36d71176.akstat.io%2F&sh=1522337455207%3D12%3A0%3A79574%2C1522337208707%3D11%3A0%3A75151%2C1522336899013%3D10%3A0%3A67853%2C1522336871676%3D9%3A0%3A60008%2C1522336828280%3D8%3A0%3A53034&dm=spirit.com&si=b9831a0c-54fb-4742-b38b-6b2ee6727d3e&ld=1522337455208&nu=&cl=1522337502807&r=https%3A%2F%2Fwww.spirit.com%2FDefault.aspx&ul=1522337502876\"; ADRUM=s=1522337502911&r=https%3A%2F%2Fwww.spirit.com%2FDefault.aspx%3F0; ADRUM_BTa=R:35|g:a70b594d-9c7c-463a-ac95-f09b43dd20bc")
	bow.Open("https://www.spirit.com/DPPCalendarMarket.aspx")

	bow.Dom().Find("div.cal-date").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})

}

func checkError(err error) {
	if err != nil {
		panic(error.Error)
	}
}
