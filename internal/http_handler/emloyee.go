package http_handler

import (
	"log"
	"net/http"

	"github.com/ring0-rootkit/golang-staff-mngmnt/internal/logging"
)

var Log *log.Logger = logging.GetFor("middlware")

func StartShiftHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`{"status":"success"}`))

	if err != nil {
		Log.Printf("Error during writing message to writer, err:%s \n", err.Error())
	}
}
