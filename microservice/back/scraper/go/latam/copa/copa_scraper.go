package main

import (
	"fmt"
	"strings"

	"../../core/util"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

// TODO -> This Scraper Not Working. COPA Detect Boot
func main() {

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.186 Safari/537.36")

	util.CheckError(bow.Open("https://www.copaair.com/es/web/co"))
	authToken := getAuthToken(bow.Body())
	fmt.Println("AuthToken", authToken)

	// Outputs: "reddit: the front page of the internet"
	fm, _ := bow.Form("form.row-fluid")
	fm.Set("paymentType", "Money")
	fm.Set("originSearchString", "Bogotá, CO BOG")
	fm.Set("destinationSearchString", "Nueva York, US JFK")
	fm.Set("pos", "CMCO")
	fm.Set("lang", "es")
	fm.Set("outboundOption.originLocationCode:", "BOG")
	fm.Set("outboundOption.destinationLocationCode:", "JFK")
	fm.Set("inboundOption.originLocationCode", "JFK")
	fm.Set("inboundOption.destinationLocationCode", "BOG")
	fm.Set("dateRangeFrom", "mayo 9, 2018")
	fm.Set("dateRangeTo", "mayo 17, 2018")
	fm.Set("outboundOption.departureDay", "9")
	fm.Set("outboundOption.departureMonth", "5")
	fm.Set("outboundOption.departureYear", "2018")
	fm.Set("inboundOption.departureDay", "17")
	fm.Set("inboundOption.departureMonth", "5")
	fm.Set("inboundOption.departureYear", "2018")
	fm.Set("tripType", "RT")
	fm.Set("cabinClass", "Economy")
	fm.Set("flexibleSearch", "false")
	fm.Set("guestTypes[0].amount", "1")
	fm.Set("guestTypes[0].type", "ADT")
	fm.Set("guestTypes[1].amount", "0")
	fm.Set("guestTypes[1].type", "CNN")
	fm.Set("guestTypes[2].amount", "0")
	fm.Set("guestTypes[2].type", "INF")
	fm.Set("coupon", "")

	// bow.AddRequestHeader("Cache-Control", "private, no-cache, no-store, must-revalidate")
	// bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.airport/get-available-destinations/airport-origin-code/BOG/store-front/CO/culture/es-CO?p_auth=" + authToken)
	// bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.route/get-route-operation-date/origin-code/BOG/destination-code/JFK?p_auth=" + authToken)
	// bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.route/get-route-operation-date/origin-code/BOG/destination-code/JFK?p_auth=" + authToken)
	// bow.Open("https://www.copaair.com/api/jsonws/copa-flight-booking-portlet.route/get-route-operation-date/origin-code/BOG/destination-code/JFK?p_auth=" + authToken)

	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Cache-Control", "max-age=0")
	bow.AddRequestHeader("Connection", "keep-alive")
	bow.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	bow.AddRequestHeader("Host", "bookings.copaair.com")
	bow.AddRequestHeader("Origin", "https://www.copaair.com")
	bow.AddRequestHeader("Referer", "https://www.copaair.com/es/web/co")

	// TODO: Try delete this line.
	//	bow.AddRequestHeader("Cookie", "JSESSIONID=8E7B62D693DA2F25DC8200E860A366E0; TS012be448=01ec83691df6b17019c97b251f585e270d6544b0570acf2e2d187eb3c6720f096bf4a89bc7f0624f4d916096459a381409033438f1; s_nr=1521380170661; s_vi=[CS]v1|2D5735A5851D0FFC-40000136A000B720[CE]; _ga=GA1.2.1404192964.1521380173; D_SID=181.48.153.214:fdF900KVJqvbWmYgR1l1JxXBOtbcqvXJ9Mlb0FkoWMo; recordID=a79a7ff7-b54f-43f3-a883-edf5bc4dd8b6; cto_lwid=61c7814c-4d28-4dd7-9d66-77ab54b79d08; s_cc=true; s_sq=%5B%5BB%5D%5D; currentLang=es; currentPos=CMco; _gid=GA1.2.922741902.1523277767; current_PoS=CMCO; TS0125fb4e=01ec83691d03232aeab452a0e48c1acf547e2d936ad772a52e477be1a3e47d760bb79661e24828f04644b509b1f986342cca5b18fecd8c9e19d5b49d749d57c0c02a3eef89; D_IID=4DA46A68-0D72-33B7-88AB-D818E3586826; D_UID=81D40323-23D5-348C-8E09-D97F860AEC54; D_ZID=24CE4209-72EB-34CE-B71C-676A54089888; D_ZUID=CFD195F3-A5BE-3A97-81C5-0DC9B473597B; D_HID=5399A537-0D7F-3A71-8DBB-E1A48CBE81D5; dmSessionID=1bc70b11-f83d-4c7d-b394-884d28a88cd8; TS0192722c=01ec83691d68e32e00c2711960384bec54239b70d546e59850c00b221813a64cbcf2ee0ade02befdc06464e6c44b8692949856de1c608f19ee7d7c73ec18e9a42a920b7feb1600d7515b1eef2dcb3781613e51bfd7a99a7e0750bdaf947ddeb814c565d59be1fd4198476176ef599cb621dd0e1351; _uetsid=_uet2a6641a7; _dc_gtm_UA-26727407-1=1")
	bow.Post("https://bookings.copaair.com/CMGS/ValidateFlow.do", "application/x-www-form-urlencoded", strings.NewReader("validator=SHOPPING"))
	bow.Post("https://bookings.copaair.com/ga346409.js?PID=B6275175-83F1-35EA-95B7-8E08CB494E56", "text/plain;charset=UTF-8", strings.NewReader("p=%7B%22proof%22%3A%22da%3A1523284270289%3AetxhgbpbhBEyXvD7N2HY%22%2C%22fp2%22%3A%7B%22userAgent%22%3A%22Mozilla%2F5.0(WindowsNT6.1%3BWin64%3Bx64)AppleWebKit%2F537.36(KHTML%2ClikeGecko)Chrome%2F65.0.3325.181Safari%2F537.36%22%2C%22language%22%3A%22es-US%22%2C%22screen%22%3A%7B%22width%22%3A1920%2C%22height%22%3A1080%2C%22availHeight%22%3A1040%2C%22availWidth%22%3A1920%2C%22outerHeight%22%3A1040%2C%22outerWidth%22%3A1920%7D%2C%22timezone%22%3A-5%2C%22indexedDb%22%3Atrue%2C%22addBehavior%22%3Afalse%2C%22openDatabase%22%3Atrue%2C%22cpuClass%22%3A%22unknown%22%2C%22platform%22%3A%22Win32%22%2C%22doNotTrack%22%3A%22unknown%22%2C%22plugins%22%3A%22ChromePDFPlugin%3A%3APortableDocumentFormat%3A%3Aapplication%2Fx-google-chrome-pdf~pdf%3BChromePDFViewer%3A%3A%3A%3Aapplication%2Fpdf~pdf%3BNativeClient%3A%3A%3A%3Aapplication%2Fx-nacl~%2Capplication%2Fx-pnacl~%3BWidevineContentDecryptionModule%3A%3AEnablesWidevinelicensesforplaybackofHTMLaudio%2Fvideocontent.(version%3A1.4.9.1070)%3A%3Aapplication%2Fx-ppapi-widevine-cdm~%22%2C%22canvas%22%3A%7B%22winding%22%3A%22yes%22%2C%22towebp%22%3Atrue%2C%22blending%22%3Atrue%2C%22img%22%3A%2235329b5fe5e35a985b7836cc8121e7a79021460e%22%7D%2C%22webGL%22%3A%7B%22img%22%3A%22bd6549c125f67b18985a8c509803f4b883ff810c%22%2C%22extensions%22%3A%22ANGLE_instanced_arrays%3BEXT_blend_minmax%3BEXT_color_buffer_half_float%3BEXT_frag_depth%3BEXT_shader_texture_lod%3BEXT_texture_filter_anisotropic%3BWEBKIT_EXT_texture_filter_anisotropic%3BEXT_sRGB%3BOES_element_index_uint%3BOES_standard_derivatives%3BOES_texture_float%3BOES_texture_float_linear%3BOES_texture_half_float%3BOES_texture_half_float_linear%3BOES_vertex_array_object%3BWEBGL_color_buffer_float%3BWEBGL_compressed_texture_s3tc%3BWEBKIT_WEBGL_compressed_texture_s3tc%3BWEBGL_compressed_texture_s3tc_srgb%3BWEBGL_debug_renderer_info%3BWEBGL_debug_shaders%3BWEBGL_depth_texture%3BWEBKIT_WEBGL_depth_texture%3BWEBGL_draw_buffers%3BWEBGL_lose_context%3BWEBKIT_WEBGL_lose_context%22%2C%22aliasedlinewidthrange%22%3A%22%5B1%2C1%5D%22%2C%22aliasedpointsizerange%22%3A%22%5B1%2C1024%5D%22%2C%22alphabits%22%3A8%2C%22antialiasing%22%3A%22yes%22%2C%22bluebits%22%3A8%2C%22depthbits%22%3A24%2C%22greenbits%22%3A8%2C%22maxanisotropy%22%3A16%2C%22maxcombinedtextureimageunits%22%3A32%2C%22maxcubemaptexturesize%22%3A16384%2C%22maxfragmentuniformvectors%22%3A1024%2C%22maxrenderbuffersize%22%3A16384%2C%22maxtextureimageunits%22%3A16%2C%22maxtexturesize%22%3A16384%2C%22maxvaryingvectors%22%3A30%2C%22maxvertexattribs%22%3A16%2C%22maxvertextextureimageunits%22%3A16%2C%22maxvertexuniformvectors%22%3A4096%2C%22maxviewportdims%22%3A%22%5B16384%2C16384%5D%22%2C%22redbits%22%3A8%2C%22renderer%22%3A%22WebKitWebGL%22%2C%22shadinglanguageversion%22%3A%22WebGLGLSLES1.0(OpenGLESGLSLES1.0Chromium)%22%2C%22stencilbits%22%3A0%2C%22vendor%22%3A%22WebKit%22%2C%22version%22%3A%22WebGL1.0(OpenGLES2.0Chromium)%22%2C%22vertexshaderhighfloatprecision%22%3A23%2C%22vertexshaderhighfloatprecisionrangeMin%22%3A127%2C%22vertexshaderhighfloatprecisionrangeMax%22%3A127%2C%22vertexshadermediumfloatprecision%22%3A23%2C%22vertexshadermediumfloatprecisionrangeMin%22%3A127%2C%22vertexshadermediumfloatprecisionrangeMax%22%3A127%2C%22vertexshaderlowfloatprecision%22%3A23%2C%22vertexshaderlowfloatprecisionrangeMin%22%3A127%2C%22vertexshaderlowfloatprecisionrangeMax%22%3A127%2C%22fragmentshaderhighfloatprecision%22%3A23%2C%22fragmentshaderhighfloatprecisionrangeMin%22%3A127%2C%22fragmentshaderhighfloatprecisionrangeMax%22%3A127%2C%22fragmentshadermediumfloatprecision%22%3A23%2C%22fragmentshadermediumfloatprecisionrangeMin%22%3A127%2C%22fragmentshadermediumfloatprecisionrangeMax%22%3A127%2C%22fragmentshaderlowfloatprecision%22%3A23%2C%22fragmentshaderlowfloatprecisionrangeMin%22%3A127%2C%22fragmentshaderlowfloatprecisionrangeMax%22%3A127%2C%22vertexshaderhighintprecision%22%3A0%2C%22vertexshaderhighintprecisionrangeMin%22%3A31%2C%22vertexshaderhighintprecisionrangeMax%22%3A30%2C%22vertexshadermediumintprecision%22%3A0%2C%22vertexshadermediumintprecisionrangeMin%22%3A31%2C%22vertexshadermediumintprecisionrangeMax%22%3A30%2C%22vertexshaderlowintprecision%22%3A0%2C%22vertexshaderlowintprecisionrangeMin%22%3A31%2C%22vertexshaderlowintprecisionrangeMax%22%3A30%2C%22fragmentshaderhighintprecision%22%3A0%2C%22fragmentshaderhighintprecisionrangeMin%22%3A31%2C%22fragmentshaderhighintprecisionrangeMax%22%3A30%2C%22fragmentshadermediumintprecision%22%3A0%2C%22fragmentshadermediumintprecisionrangeMin%22%3A31%2C%22fragmentshadermediumintprecisionrangeMax%22%3A30%2C%22fragmentshaderlowintprecision%22%3A0%2C%22fragmentshaderlowintprecisionrangeMin%22%3A31%2C%22fragmentshaderlowintprecisionrangeMax%22%3A30%7D%2C%22touch%22%3A%7B%22maxTouchPoints%22%3A0%2C%22touchEvent%22%3Afalse%2C%22touchStart%22%3Afalse%7D%2C%22video%22%3A%7B%22ogg%22%3A%22probably%22%2C%22h264%22%3A%22probably%22%2C%22webm%22%3A%22probably%22%7D%2C%22audio%22%3A%7B%22ogg%22%3A%22probably%22%2C%22mp3%22%3A%22probably%22%2C%22wav%22%3A%22probably%22%2C%22m4a%22%3A%22maybe%22%7D%2C%22vendor%22%3A%22GoogleInc.%22%2C%22product%22%3A%22Gecko%22%2C%22productSub%22%3A%2220030107%22%2C%22browser%22%3A%7B%22ie%22%3Afalse%2C%22chrome%22%3Atrue%2C%22webdriver%22%3Afalse%7D%2C%22fonts%22%3A%22Batang%3BCalibri%3BCentury%3BLeelawadee%3BMarlett%3BPMingLiU%3BPristina%3BSimHei%3BVrinda%22%7D%2C%22cookies%22%3A1%2C%22setTimeout%22%3A0%2C%22setInterval%22%3A0%2C%22appName%22%3A%22Netscape%22%2C%22platform%22%3A%22Win32%22%2C%22syslang%22%3A%22es-US%22%2C%22userlang%22%3A%22es-US%22%2C%22cpu%22%3A%22%22%2C%22productSub%22%3A%2220030107%22%2C%22plugins%22%3A%7B%220%22%3A%22ChromePDFPlugin%22%2C%221%22%3A%22ChromePDFViewer%22%2C%222%22%3A%22NativeClient%22%2C%223%22%3A%22WidevineContentDecryptionModule%22%7D%2C%22mimeTypes%22%3A%7B%220%22%3A%22application%2Fpdf%22%2C%221%22%3A%22PortableDocumentFormatapplication%2Fx-google-chrome-pdf%22%2C%222%22%3A%22NativeClientExecutableapplication%2Fx-nacl%22%2C%223%22%3A%22PortableNativeClientExecutableapplication%2Fx-pnacl%22%2C%224%22%3A%22WidevineContentDecryptionModuleapplication%2Fx-ppapi-widevine-cdm%22%7D%2C%22screen%22%3A%7B%22width%22%3A1920%2C%22height%22%3A1080%2C%22colorDepth%22%3A24%7D%2C%22fonts%22%3A%7B%220%22%3A%22Calibri%22%2C%221%22%3A%22Cambria%22%2C%222%22%3A%22Constantia%22%2C%223%22%3A%22Georgia%22%2C%224%22%3A%22SegoeUI%22%2C%225%22%3A%22Candara%22%2C%226%22%3A%22TrebuchetMS%22%2C%227%22%3A%22Verdana%22%2C%228%22%3A%22Consolas%22%2C%229%22%3A%22LucidaConsole%22%2C%2210%22%3A%22CourierNew%22%2C%2211%22%3A%22Courier%22%7D%7D"))
	bow.Open("https://bookings.copaair.com/CMCO/html/empty.html?version=201803030746")

	util.CheckError(fm.Submit())
	bow.Post("https://bookings.copaair.com/CMGS/AirLowFareSearchExt.do", "text/plain;charset=UTF-8", strings.NewReader("ajaxAction=true"))
	bow.Open("https://bookings.copaair.com/CMCO/es/pages/tdp/interstitials/airLowFareSearch.html?version=201803030746")

	// form := url.Values{}
	// form.Add("ajaxAction", "true")
	// bow.Post("https://bookings.copaair.com/CMGS/AirLowFareSearchExt.do", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))

	bow.Open("https://bookings.copaair.com/CMGS/AirFareFamiliesForward.do")

	fmt.Println(bow.Body())

	bow.Dom().Find("tr.rowOdd").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Html())
	})
}

func getAuthToken(body string) string {
	position := strings.Index(body, "ng-init=\"p_au=")
	token := body[position+19 : position+27]
	fmt.Println(token)
	return strings.Replace(token, "\"", "", -1)
}
