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
	marathonHost net.IP = net.ParseIP("127.0.0.1")
	marathonPort uint   = 8080
)

func GetDeployedApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	env := vars["env"]
	app := vars["app"]
	path := fmt.Sprintf("/%v/%v", env, app)

	ms, err := marathon.NewService(marathonHost, marathonPort)
	if err != nil {
		log.Printf("Failed to connect to Marathon endpoint. %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	mapp, err := ms.GetApp(path)
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
	marathonHostF := flag.IPP("marathon-ip", "H", marathonHost, "Marathon endpoint IP address")
	marathonPortF := flag.UintP("marathon-port", "P", marathonPort, "Marathon endpoint Port number.")
	serveAddr := flag.IPP("listen-addr", "l", net.IPv4zero, "Listener IP address.")
	servePort := flag.IntP("listen-port", "p", 3000, "Listener port number.")
	dbHostname := flag.StringP("db-host", "", "localhost", "mySQL server hostname.")
	dbPort := flag.IntP("db-port", "", 3306, "mySQL server port number.")
	dbUsername := flag.StringP("db-username", "", "root", "database username.")
	dbPassword := flag.StringP("db-password", "", "", "database password.")
	dbName := flag.StringP("db-name", "", "serviced", "database name.")
	flag.Parse()
	marathonHost = *marathonHostF
	marathonPort = *marathonPortF

	r := mux.NewRouter()

	db, err := OpenDB(*dbHostname, *dbPort, *dbUsername, *dbPassword, *dbName)
	if err != nil {
		log.Fatalf("NewDB error. %+v\n", err)
	}

	log.Printf("db: %+v\n", db)

	r.Path("/deployments/{env}/{app}").
		HandlerFunc(GetDeployedApp).
		Methods("GET")
	r.Path("/deployments/{env}/{app}/{release}").
		HandlerFunc(DeployAppRelease).
		Methods("POST")

	http.ListenAndServe(fmt.Sprintf("%v:%v", *serveAddr, *servePort), r)
}
