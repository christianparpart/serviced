package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/christianparpart/serviced/marathon"
	"github.com/gorilla/mux"
	flag "github.com/ogier/pflag"
)

var (
	marathonHost string = "localhost"
	marathonPort int    = 8080
)

func GetDeployedApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	env := vars["env"]
	app := vars["app"]
	path := fmt.Sprintf("/%v/%v", env, app)

	mapp, err := marathon.GetApp(marathonHost, marathonPort, path)
	if err != nil {
		log.Printf("Failed to get app. %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(mapp)
	if err != nil {
		log.Printf("Failed to marshal to JSON. %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(out))
}

func DeployAppRelease(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	env := vars["env"]
	app := vars["app"]
	release := vars["release"]

	fmt.Fprintf(w, "Deploy App!\n %v-%v-%v\n", env, app, release)
}

func main() {
	marathonHostF := flag.StringP("marathon-host", "H", marathonHost, "Marathon Host")
	marathonPortF := flag.IntP("marathon-port", "P", marathonPort, "Marathon Port.")
	serveAddr := flag.IPP("listen-addr", "l", net.IPv4zero, "Listener IP address.")
	servePort := flag.IntP("listen-port", "p", 3000, "Listener port number.")
	flag.Parse()
	marathonHost = *marathonHostF
	marathonPort = *marathonPortF

	r := mux.NewRouter()

	r.Path("/deployments/{env}/{app}").
		HandlerFunc(GetDeployedApp).
		Methods("GET")
	r.Path("/deployments/{env}/{app}/{release}").
		HandlerFunc(DeployAppRelease).
		Methods("POST")

	http.ListenAndServe(fmt.Sprintf("%v:%v", *serveAddr, *servePort), r)
}
