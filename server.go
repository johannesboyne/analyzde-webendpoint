package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/analyzde.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(`
(function (i,a) {
var x = new XMLHttpRequest(); 
x.open("GET","http://analyz.de/?id=` + r.FormValue("id") + `&event=pageload&uri="+encodeURIComponent(window.location.href)+"&action="+a,false);
x.send(null);
})("` + r.FormValue("id") + `", "pageview")`))
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
