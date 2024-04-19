package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ring0-rootkit/golang-staff-mngmnt/internal/common/logging"
	"github.com/ring0-rootkit/golang-staff-mngmnt/internal/middlware_service/repository"
)

var Log *log.Logger = logging.GetFor("middlware http handler")

func StartShiftHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		Log.Printf("Error while getting path value (id) from path: %s \n", r.URL.Path)
		return
	}

	err = repository.StartWorkShift(id)
	if err != nil {
		Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
		return
	}

	_, err = w.Write([]byte(`{"status":"success"}`))
	if err != nil {
		Log.Printf("Error during writing message to writer [start shift], err:%s \n", err.Error())
		return
	}
}

func EndShiftHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		Log.Printf("Error while getting path value (id) from path: %s \n", r.URL.Path)
		return
	}

	err = repository.EndWorkShift(id)
	if err != nil {
		Log.Printf("Error when trying to send grpc call. Err:%s\n", err.Error())
		return
	}

	_, err = w.Write([]byte(`{"status":"success"}`))
	if err != nil {
		Log.Printf("Error during writing message to writer [end shift], err:%s \n", err.Error())
		return
	}
}
