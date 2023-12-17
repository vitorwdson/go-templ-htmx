package utils

import "github.com/labstack/echo/v4"

type Htmx struct {
	Boosted               bool
	CurrentURL            string
	HistoryRestoreRequest bool
	Prompt                string
	Target                string
	TriggerName           string
	Trigger               string
}

func IsHtmx(c echo.Context) *Htmx {
	header := c.Request().Header

	// HX-Request always “true”
	if req := header.Get("HX-Request"); req != "true" {
		return nil
	}

	boosted := header.Get("HX-Boosted") == "true"
	historyRestoreRequest := header.Get("HX-History-Restore-Request") == "true"

	return &Htmx{
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
