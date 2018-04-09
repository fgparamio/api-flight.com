package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/headzoo/surf"
)

func main() {

	// create a new collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
		colly.AllowedDomains("www.aviorair.com"))

	// authenticate
	err := c.Post("https://www.aviorair.com/Booking/sendSSW", map[string]string{
		"journeySpan": "RT", "cityOrigin": "CCS", "cityDestination": "MDE",
		"departureDate": "04-30-2018", "returnDate": "05-10-2018",
		"numAdults": "1", "numChildren": "0", "numInfants": "0"})

	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("response received", string(r.Body))
	})

	// start scraping
	c.Visit("https://www.aviorair.com")

	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.SetUserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	bow.AddRequestHeader("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	bow.AddRequestHeader("Host", "www.aviorair.com")
	bow.AddRequestHeader("Upgrade-Insecure-Requests", "1")

	util.CheckError(bow.Open("https://www.aviorair.com"))

	fmt.Println(bow.Body())

	fm, _ := bow.Form("[name='formBookingFlight']")
	bow.Dom().Find("form").SetAttr("target", "_self")

	fm.Set("journeySpan", "RT")
	fm.Set("cityOrigin", "CCS")
	fm.Set("cityDestination", "MDE")
	fm.Set("departureDate", "04-30-2018")
	fm.Set("returnDate", "05-10-2018")
	fm.Set("numAdults", "1")
	fm.Set("numChildren", "0")
	fm.Set("numInfants", "0")

	util.CheckError(fm.Submit())

	bow.Open("https://dx.aviorair.com/dx/9VDX/")

	bow.Post("https://dx.aviorair.com/ga354671.js?PID=6744DD17-613F-3144-8E5A-421359AE340F", "text/plain;charset=UTF-8", strings.NewReader("p=%7B%22proof%22%3A%2282%3A1523221422384%3Aln4V7ogrI4nGhud1z4xH%22%2C%22fp2%22%3A%7B%22userAgent%22%3A%22Mozilla%2F5.0(WindowsNT6.1%3BWin64%3Bx64)AppleWebKit%2F537.36(KHTML%2ClikeGecko)Chrome%2F65.0.3325.181Safari%2F537.36%22%2C%22language%22%3A%22es-US%22%2C%22screen%22%3A%7B%22width%22%3A1920%2C%22height%22%3A1080%2C%22availHeight%22%3A1080%2C%22availWidth%22%3A1920%2C%22innerHeight%22%3A974%2C%22innerWidth%22%3A879%2C%22outerHeight%22%3A1040%2C%22outerWidth%22%3A1920%7D%2C%22timezone%22%3A-5%2C%22indexedDb%22%3Atrue%2C%22addBehavior%22%3Afalse%2C%22openDatabase%22%3Atrue%2C%22cpuClass%22%3A%22unknown%22%2C%22platform%22%3A%22Win32%22%2C%22doNotTrack%22%3A%22unknown%22%2C%22plugins%22%3A%22ChromePDFPlugin%3A%3APortableDocumentFormat%3A%3Aapplication%2Fx-google-chrome-pdf~pdf%3BChromePDFViewer%3A%3A%3A%3Aapplication%2Fpdf~pdf%3BNativeClient%3A%3A%3A%3Aapplication%2Fx-nacl~%2Capplication%2Fx-pnacl~%3BWidevineContentDecryptionModule%3A%3AEnablesWidevinelicensesforplaybackofHTMLaudio%2Fvideocontent.(version%3A1.4.9.1070)%3A%3Aapplication%2Fx-ppapi-widevine-cdm~%22%2C%22canvas%22%3A%7B%22winding%22%3A%22yes%22%2C%22towebp%22%3Atrue%2C%22blending%22%3Atrue%2C%22img%22%3A%2235329b5fe5e35a985b7836cc8121e7a79021460e%22%7D%2C%22webGL%22%3A%7B%22img%22%3A%22bd6549c125f67b18985a8c509803f4b883ff810c%22%2C%22extensions%22%3A%22ANGLE_instanced_arrays%3BEXT_blend_minmax%3BEXT_color_buffer_half_float%3BEXT_frag_depth%3BEXT_shader_texture_lod%3BEXT_texture_filter_anisotropic%3BWEBKIT_EXT_texture_filter_anisotropic%3BEXT_sRGB%3BOES_element_index_uint%3BOES_standard_derivatives%3BOES_texture_float%3BOES_texture_float_linear%3BOES_texture_half_float%3BOES_texture_half_float_linear%3BOES_vertex_array_object%3BWEBGL_color_buffer_float%3BWEBGL_compressed_texture_s3tc%3BWEBKIT_WEBGL_compressed_texture_s3tc%3BWEBGL_compressed_texture_s3tc_srgb%3BWEBGL_debug_renderer_info%3BWEBGL_debug_shaders%3BWEBGL_depth_texture%3BWEBKIT_WEBGL_depth_texture%3BWEBGL_draw_buffers%3BWEBGL_lose_context%3BWEBKIT_WEBGL_lose_context%22%2C%22aliasedlinewidthrange%22%3A%22%5B1%2C1%5D%22%2C%22aliasedpointsizerange%22%3A%22%5B1%2C1024%5D%22%2C%22alphabits%22%3A8%2C%22antialiasing%22%3A%22yes%22%2C%22bluebits%22%3A8%2C%22depthbits%22%3A24%2C%22greenbits%22%3A8%2C%22maxanisotropy%22%3A16%2C%22maxcombinedtextureimageunits%22%3A32%2C%22maxcubemaptexturesize%22%3A16384%2C%22maxfragmentuniformvectors%22%3A1024%2C%22maxrenderbuffersize%22%3A16384%2C%22maxtextureimageunits%22%3A16%2C%22maxtexturesize%22%3A16384%2C%22maxvaryingvectors%22%3A30%2C%22maxvertexattribs%22%3A16%2C%22maxvertextextureimageunits%22%3A16%2C%22maxvertexuniformvectors%22%3A4096%2C%22maxviewportdims%22%3A%22%5B16384%2C16384%5D%22%2C%22redbits%22%3A8%2C%22renderer%22%3A%22WebKitWebGL%22%2C%22shadinglanguageversion%22%3A%22WebGLGLSLES1.0(OpenGLESGLSLES1.0Chromium)%22%2C%22stencilbits%22%3A0%2C%22vendor%22%3A%22WebKit%22%2C%22version%22%3A%22WebGL1.0(OpenGLES2.0Chromium)%22%2C%22vertexshaderhighfloatprecision%22%3A23%2C%22vertexshaderhighfloatprecisionrangeMin%22%3A127%2C%22vertexshaderhighfloatprecisionrangeMax%22%3A127%2C%22vertexshadermediumfloatprecision%22%3A23%2C%22vertexshadermediumfloatprecisionrangeMin%22%3A127%2C%22vertexshadermediumfloatprecisionrangeMax%22%3A127%2C%22vertexshaderlowfloatprecision%22%3A23%2C%22vertexshaderlowfloatprecisionrangeMin%22%3A127%2C%22vertexshaderlowfloatprecisionrangeMax%22%3A127%2C%22fragmentshaderhighfloatprecision%22%3A23%2C%22fragmentshaderhighfloatprecisionrangeMin%22%3A127%2C%22fragmentshaderhighfloatprecisionrangeMax%22%3A127%2C%22fragmentshadermediumfloatprecision%22%3A23%2C%22fragmentshadermediumfloatprecisionrangeMin%22%3A127%2C%22fragmentshadermediumfloatprecisionrangeMax%22%3A127%2C%22fragmentshaderlowfloatprecision%22%3A23%2C%22fragmentshaderlowfloatprecisionrangeMin%22%3A127%2C%22fragmentshaderlowfloatprecisionrangeMax%22%3A127%2C%22vertexshaderhighintprecision%22%3A0%2C%22vertexshaderhighintprecisionrangeMin%22%3A31%2C%22vertexshaderhighintprecisionrangeMax%22%3A30%2C%22vertexshadermediumintprecision%22%3A0%2C%22vertexshadermediumintprecisionrangeMin%22%3A31%2C%22vertexshadermediumintprecisionrangeMax%22%3A30%2C%22vertexshaderlowintprecision%22%3A0%2C%22vertexshaderlowintprecisionrangeMin%22%3A31%2C%22vertexshaderlowintprecisionrangeMax%22%3A30%2C%22fragmentshaderhighintprecision%22%3A0%2C%22fragmentshaderhighintprecisionrangeMin%22%3A31%2C%22fragmentshaderhighintprecisionrangeMax%22%3A30%2C%22fragmentshadermediumintprecision%22%3A0%2C%22fragmentshadermediumintprecisionrangeMin%22%3A31%2C%22fragmentshadermediumintprecisionrangeMax%22%3A30%2C%22fragmentshaderlowintprecision%22%3A0%2C%22fragmentshaderlowintprecisionrangeMin%22%3A31%2C%22fragmentshaderlowintprecisionrangeMax%22%3A30%7D%2C%22touch%22%3A%7B%22maxTouchPoints%22%3A0%2C%22touchEvent%22%3Afalse%2C%22touchStart%22%3Afalse%7D%2C%22video%22%3A%7B%22ogg%22%3A%22probably%22%2C%22h264%22%3A%22probably%22%2C%22webm%22%3A%22probably%22%7D%2C%22audio%22%3A%7B%22ogg%22%3A%22probably%22%2C%22mp3%22%3A%22probably%22%2C%22wav%22%3A%22probably%22%2C%22m4a%22%3A%22maybe%22%7D%2C%22vendor%22%3A%22GoogleInc.%22%2C%22product%22%3A%22Gecko%22%2C%22productSub%22%3A%2220030107%22%2C%22browser%22%3A%7B%22ie%22%3Afalse%2C%22chrome%22%3Atrue%2C%22webdriver%22%3Afalse%7D%2C%22fonts%22%3A%22Batang%3BCalibri%3BCentury%3BLeelawadee%3BMarlett%3BPMingLiU%3BPristina%3BSimHei%3BVrinda%22%7D%2C%22cookies%22%3A1%2C%22setTimeout%22%3A0%2C%22setInterval%22%3A0%2C%22appName%22%3A%22Netscape%22%2C%22platform%22%3A%22Win32%22%2C%22syslang%22%3A%22es-US%22%2C%22userlang%22%3A%22es-US%22%2C%22cpu%22%3A%22%22%2C%22productSub%22%3A%2220030107%22%2C%22plugins%22%3A%7B%220%22%3A%22ChromePDFPlugin%22%2C%221%22%3A%22ChromePDFViewer%22%2C%222%22%3A%22NativeClient%22%2C%223%22%3A%22WidevineContentDecryptionModule%22%7D%2C%22mimeTypes%22%3A%7B%220%22%3A%22application%2Fpdf%22%2C%221%22%3A%22PortableDocumentFormatapplication%2Fx-google-chrome-pdf%22%2C%222%22%3A%22NativeClientExecutableapplication%2Fx-nacl%22%2C%223%22%3A%22PortableNativeClientExecutableapplication%2Fx-pnacl%22%2C%224%22%3A%22WidevineContentDecryptionModuleapplication%2Fx-ppapi-widevine-cdm%22%7D%2C%22screen%22%3A%7B%22width%22%3A1920%2C%22height%22%3A1080%2C%22colorDepth%22%3A24%7D%2C%22fonts%22%3A%7B%220%22%3A%22Calibri%22%2C%221%22%3A%22Cambria%22%2C%222%22%3A%22Constantia%22%2C%223%22%3A%22Georgia%22%2C%224%22%3A%22SegoeUI%22%2C%225%22%3A%22Candara%22%2C%226%22%3A%22TrebuchetMS%22%2C%227%22%3A%22Verdana%22%2C%228%22%3A%22Consolas%22%2C%229%22%3A%22LucidaConsole%22%2C%2210%22%3A%22CourierNew%22%2C%2211%22%3A%22Courier%22%7D%7D"))

	var jsonReq = []byte(`{"cabinClass":"Economy","promoCodes":[],"passengers":{"ADT":1,"CHD":0,"INF":0},"searchType":"BRANDED",
		"itineraryParts":[{"from":{"useNearbyLocations":false,"code":"CCS"},"to":{"useNearbyLocations":false,"code":"MDE"},
		"when":{"date":"2018-04-30"}},{"from":{"useNearbyLocations":false,"c`)

	bow.Post("https://dc.aviorair.com/v3.4/ssw/products/air/search?jipcc=9VDX", "application/json", strings.NewReader(string(jsonReq)))

	fmt.Println(bow.Body())

	jq := util.ToJSONQ(bow.Body())

	flights, _ := jq.ArrayOfObjects("brandedResults", "itineraryPartBrands")
	fmt.Println("Flights:", flights)

}
