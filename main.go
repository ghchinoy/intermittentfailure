package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	emails = [...]string{"gordon.gopher@fake.com", "polly.pixel@fake.com", "andy.android@fake.com"}
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", unreliableHandler(invoiceHandler))
	http.ListenAndServe(":8080", nil)
}

// unreliableHandler randomly returns error codes
func unreliableHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// failure chance
		fail := randomInt(0, 4)
		willFail := false
		if fail == 0 {
			willFail = true
		}
		log.Println(r.URL, fail, willFail)
		if willFail {
			expectedFailureHandler(w, r)
		} else {
			fn(w, r)
		}
	}
}

// expectedFailureHandler returns an expected failure
func expectedFailureHandler(w http.ResponseWriter, r *http.Request) {

	failtype := randomInt(0, 4)

	switch failtype {
	case 1:
		http.Error(w, "Planned outage", http.StatusInternalServerError)
	case 2:
		http.Error(w, "Planned outage", http.StatusBadGateway)
	case 3:
		http.Error(w, "Planned outage", http.StatusServiceUnavailable)
	default:
		http.Error(w, "Planned outage", http.StatusInternalServerError)

	}

	return
}

// invoiceHandler returns the e-mail and invoice count
func invoiceHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Email        string
		OpenInvoices int
	}{
		emails[randomInt(0, 3)],
		randomInt(0, 11),
	}

	databytes, err := json.Marshal(&data)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")

	w.Write(databytes)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
