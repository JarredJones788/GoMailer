package router

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mailer"
	"net/http"
	"templates"
	"types"

	"github.com/gorilla/mux"
)

//Router type
type Router struct {
	Host        string
	Mailer      *mailer.Emailer
	EmailSecret string
}

//Init - inits all routes.
func (router Router) Init(mailer *mailer.Emailer, config *config.Config) error {

	router.Host = config.Host
	router.Mailer = mailer

	//Get email secret. Incoming request need to match this key
	data, err := ioutil.ReadFile(config.EmailKey)
	if err != nil {
		return err
	}

	//Set key from the file
	router.EmailSecret = string(data)

	//Setup mux router
	r := mux.NewRouter()
	router.setUpRoutes(r)
	fmt.Println("Server Started")
	http.ListenAndServe(config.Port, r)

	return nil
}

//setUpRoutes - sets up all endpoints for the service
func (router Router) setUpRoutes(r *mux.Router) {
	r.HandleFunc("/api/email/send", router.sendEmail)
}

//-----------------HELPERS BELOW-----------------\\

//badRequest - returns a generic bad response
func (router Router) badRequest(w http.ResponseWriter) {
	failed, err := json.Marshal(types.GenericResponse{Response: false})
	if err != nil {
		w.Write([]byte("BACKEND ERROR"))
		return
	}
	w.Write(failed)
}

//goodRequest - returns a generic good response
func (router Router) goodRequest(w http.ResponseWriter) {
	good, err := json.Marshal(types.GenericResponse{Response: true})
	if err != nil {
		w.Write([]byte("BACKEND ERROR"))
		return
	}
	w.Write(good)
}

//reasonRequest - returns a response with a reason
func (router Router) reasonRequest(w http.ResponseWriter, response bool, reason string) {
	good, err := json.Marshal(types.ReasonResponse{Response: response, Reason: reason})
	if err != nil {
		w.Write([]byte("BACKEND ERROR"))
		return
	}
	w.Write(good)
}

//setUpHeaders - sets the desired headers for an http response
func (router Router) setUpHeaders(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", router.Host)
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Max-Age", "120")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
		return false
	}
	return true
}

//-----------------ROUTES BELOW-----------------\\

//sendEmail - sends requested email
func (router Router) sendEmail(w http.ResponseWriter, r *http.Request) {
	if !router.setUpHeaders(w, r) {
		return //request was an OPTIONS which was handled.
	}

	var data types.EmailData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("Test Error: " + err.Error())
		router.badRequest(w)
		return
	}

	//Make sure key provided is the correct one.
	//Only services authorized should have the key
	if data.Key != router.EmailSecret {
		fmt.Println("Email key does not match")
		router.reasonRequest(w, false, "Email key does not match")
		return
	}

	var template mailer.Template

	//Get the requested template
	switch data.Template {
	case "Test":
		template = templates.TestTemplate{}.Create(data.Contacts)
		break
	default:
		fmt.Println("No template " + data.Template + " found")
		router.reasonRequest(w, false, "No template "+data.Template+" found")
		return
	}

	//Send email
	//go router.Mailer.SendEmail(template)
	fmt.Println(template.GetFirstRecipient().Name)

	router.goodRequest(w)
}
