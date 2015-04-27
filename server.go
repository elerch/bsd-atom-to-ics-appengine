package hello

import (
    "net/http",
    "github.com/elerch/bsd-atom-to-ics"
)

func init() {
    http.HandleFunc("/", handler);
}

func handler(w http.ResponseWriter, r *http.Request) {
    bsdatomtoics.AtomToICS(bsdatomtoics.FetchBytes(false), w, false);
}


