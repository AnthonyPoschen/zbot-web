// call 'goapp serve' in this folder.
package webserver

import (
	//"fmt"
	"net/http"
	//"strings"

	"appengine"
)

// appengineHandler wraps http.Handler to pass it a new `appengine.Context` and handle errors.
type appengineHandler func(c appengine.Context, w http.ResponseWriter, r *http.Request) error

func (h appengineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if err := h(c, w, r); err != nil {
		c.Errorf("%v", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//var	chttp = http.NewServeMux()
// TODO Setup a proper file handler to load a websites content.
// TODO setup a interface for api calls so that data can be passed around.
func init() {
	http.Handle("/", appengineHandler(handle))
}

func handle(c appengine.Context, w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte(r.URL.String() + ":D"))
	return nil
}
