package bsdatomtoicsappengine

import (
    "net/http"
    "github.com/elerch/bsd-atom-to-ics"
    "appengine"
    "appengine/urlfetch"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    school := r.URL.Query().Get("s");
    selectedSchool := school
    if selectedSchool == "" { selectedSchool = "district" }
    w.Header().Set("X-School", selectedSchool)
    c := appengine.NewContext(r)
    // GAE doesn't seem to like this cert
    client := createClient(c, true)
    atom, err := bsdatomtoics.FetchBytesWith(client, school)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    bsdatomtoics.AtomToICS(atom, w, false)
}

func createClient(context appengine.Context, allowInvalidServerCertificate bool) *http.Client {
    return &http.Client{
        Transport: &urlfetch.Transport{
            Context:  context,
            AllowInvalidServerCertificate: allowInvalidServerCertificate,
        },
    }
}

