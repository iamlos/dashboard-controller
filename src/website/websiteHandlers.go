package website

import (
	"github.com/gorilla/mux"
	"master/master"
	"net/http"
	"website/session"
	"sort"
)

func InitiateWebsiteHandlers(slaveMap map[string]master.Slave, router *mux.Router) {
	router.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir(IMAGES_PATH))))
	router.Handle("/assets/javascripts/", http.StripPrefix("/assets/javascripts/", http.FileServer(http.Dir(JAVASCRIPTS_PATH))))
	router.Handle("/assets/stylesheets/", http.StripPrefix("/assets/stylesheets/", http.FileServer(http.Dir(STYLESHEETS_PATH))))

	router.HandleFunc("/", IndexPageHandler)
	router.HandleFunc("/login", session.LoginHandler).Methods("POST")
	router.HandleFunc("/logout", session.LogoutHandler).Methods("POST")

	router.HandleFunc("/internal", func(w http.ResponseWriter, r *http.Request) {
		slaveNames := GetSlaveNamesFromMap(slaveMap)
		FormHandler(w, r, slaveNames)
	})
	router.HandleFunc("/form-submit", func(w http.ResponseWriter, r *http.Request) {
		SubmitHandler(w, r, slaveMap)
	})
}

func GetSlaveNamesFromMap(slaveMap map[string]master.Slave) (slaveNames []string) {
	for k := range slaveMap {
		slaveNames = append(slaveNames, k)
	}
	sort.Strings(slaveNames)
	return
}
