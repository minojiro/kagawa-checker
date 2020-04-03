package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/ip2location/ip2location-go"
)

var db, dbErr = ip2location.OpenDB("./tmp/IP2LOCATION-LITE-DB3.BIN")

type apiResponse struct {
	IsKagawa bool   `json:"is_kagawa"`
	IP       string `json:"ip"`
}

func checkKagawa(ip string) bool {
	results, _ := db.Get_all(ip)
	return results.Region == "Kagawa"
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	if ip == "" {
		ip = r.Header.Get("X-Real-Ip")
	}
	return ip
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	var res = &apiResponse{
		IsKagawa: checkKagawa(getIP(r)),
		IP:       ip,
	}
	j, _ := json.Marshal(res)
	w.Write(j)
}

func main() {
	if dbErr != nil {
		log.Println("failed to open: ip2location")
		return
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", apiHandler)
	log.Println("Listening port " + port)
	http.ListenAndServe(":"+port, nil)
}
