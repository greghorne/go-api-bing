package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"strings"
)




// ============================================================
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/v1/bing-isochrone/{lng}/{lat}/{time}/{key}", v1BingIsochrone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8002", router))

}
// ============================================================


// ============================================================
func v1BingIsochrone (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var jsonResult map[string]string

	if isochrone, msg := v1DoBingIsochrone(params["lng"], params["lat"], params["time"], params["key"]); msg == "" {
		jsonResult = map[string]string{"bing": isochrone}
	} else {
		jsonResult = map[string]string{"bing": msg}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(jsonResult)
}
// ============================================================


// ============================================================
func v1DoBingIsochrone(sxLng string, syLat string, sTime string, sKey string) (geojson string, msg string) {

	bing_url := "http://dev.virtualearth.net/REST/v1/Routes/Isochrones?waypoint=" +
		syLat + "," + sxLng + "&maxTime=" + sTime + "&timeUnit=Seconds&travelMode=Driving&key=" + sKey

	startSearchText := "\"polygons\":[{\"coordinates\":"
	endSearchText   := "]}]}],\"statusCode\""

	geojson = ""
	msg     = ""

	response, err := http.Get(bing_url)
	if err == nil {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			geojson = ""
			msg     = err.Error()
		} 

		jsonText := string(body)

		nStart   := strings.Index(jsonText, startSearchText) + len(startSearchText)
		nEnd     := strings.Index(jsonText, endSearchText)

		geojson = strings.Replace(jsonText[nStart:nEnd], "}", "", -1)
	} 

	return
}
// ============================================================


