package handlers

import (
	"github.com/palanti-r/go_final_project/utils"
	"net/http"
	"time"
)

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	originalDate, err := time.Parse("20060102", now)
	if err != nil {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}
	nextDate, err := utils.NextDate(originalDate, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(nextDate))

}
