package utils

import (
	"net/http"
)

type Htmx struct {
	Boosted               bool
	CurrentURL            string
	HistoryRestoreRequest bool
	Prompt                string
	Target                string
	TriggerName           string
	Trigger               string
}

func IsHtmx(r *http.Request) bool {
	// HX-Request always “true”
	if req := r.Header.Get("HX-Request"); req != "true" {
		return false
	}

	return true
}

func GetHtmxData(r *http.Request) Htmx {
	header := r.Header

	boosted := header.Get("HX-Boosted") == "true"
	historyRestoreRequest := header.Get("HX-History-Restore-Request") == "true"

	return Htmx{
		// HX-Boosted indicates that the request is via an element using hx-boost
		Boosted: boosted,

		// HX-Current-URL the current URL of the browser
		CurrentURL: header.Get("HX-Current-URL"),

		// HX-History-Restore-Request “true” if the request is for history restoration after a miss in the local history cache
		HistoryRestoreRequest: historyRestoreRequest,

		// HX-Prompt the user response to an hx-prompt
		Prompt: header.Get("HX-Prompt"),

		// HX-Target the id of the target element if it exists
		Target: header.Get("HX-Target"),

		// HX-Trigger-Name the name of the triggered element if it exists
		TriggerName: header.Get("HX-Trigger-Name"),

		// HX-Trigger the id of the triggered element if it exists
		Trigger: header.Get("HX-Trigger"),
	}
}

func RedirectHtmx(w http.ResponseWriter, r *http.Request, url string) error {
	htmx := IsHtmx(r)
	if !htmx {
		http.Redirect(w, r, url, http.StatusFound)
		return nil
	}

	w.Header().Set("HX-Location", url)
	w.WriteHeader(http.StatusOK)

	return nil
}
