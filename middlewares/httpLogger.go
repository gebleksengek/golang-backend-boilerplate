package middlewares

import (
	"log"
	"net/http"
	"os"
	"time"
)

//RequestLogger RequestLogger
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("logs"); os.IsNotExist(err) {
			os.Mkdir("logs", 0744)
		}
		f, err := os.OpenFile("logs/requests-"+time.Now().Format("01-02-2006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0744)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		logFormat := r.RemoteAddr + " - " + r.Header.Get("X-Forwarded-For") + " - " + r.Header.Get("X-Real-IP") + " - " + " " + r.Method + " " + r.RequestURI + " - " + r.Header.Get("User-Agent") + " - \"" + r.Header.Get("Authorization") + "\""
		logger := log.New(f, "", log.LstdFlags)
		logger.Println(logFormat)
		log.Println(logFormat)
		next.ServeHTTP(w, r)
	})
}
