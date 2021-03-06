package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	public := http.FileServer(http.Dir("../../../public"))
	http.Handle("/public/", http.StripPrefix("/public/", public))
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.HandleFunc("/test{{component}}", sayhello)
	http.ListenAndServe("0.0.0.0:8004", nil)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("I Hava Get Your Message:" + string(b)))
}
