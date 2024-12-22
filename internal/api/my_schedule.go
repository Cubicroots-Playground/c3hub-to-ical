package api

import (
	"net/http"

	ics "github.com/arran4/golang-ical"
)

func (service *service) MySchedule(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	if len(vals["token"]) == 0 || vals["token"][0] != service.conf.Token {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("unauthorized"))
		return
	}

	events, err := service.hubService.GetMySchedule(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	for _, ev := range events {
		calEv := cal.AddEvent(ev.ID)
		calEv.SetStartAt(ev.StartTime)
		calEv.SetEndAt(ev.EndTime)
		calEv.SetSummary(ev.Name)
		calEv.SetLocation(ev.Room)
	}

	respBody := cal.Serialize()
	w.Header().Set("Content-Disposition", "attachment; filename=my_schedule.ics")
	_, _ = w.Write([]byte(respBody))
}
