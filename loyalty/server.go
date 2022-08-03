package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	//Mongo
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//html templates
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	//bcrypt hash and salt
	"golang.org/x/crypto/bcrypt"

	//prometheus
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Loyalist struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	LoyaltyTier string
}

type PageData struct {
	PageTitle  string
	Message    string
	SiteAction string
	ActionPage string
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("kefue-secret-198")
	store = sessions.NewCookieStore(key)
	//mongodb client declaration
	client      *mongo.Client
	clientError error
	certString string = ""
	mongoHost string = os.Getenv("MONGO_HOST")
	mongoUser string = os.Getenv("MONGO_USER")
	mongoPass string = os.Getenv("MONGO_PASS")
	mongoTLS string = os.Getenv("MONGO_TLS")
)

// Page Handlers
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method, "on URL:", r.URL)
	if r.URL.Path != "/healthz" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Ready")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method, "on URL:", r.URL)
	session, _ := store.Get(r, "cookie-name")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./static/register.html")
		t.Execute(w, nil)
	} else {
		found := checkLoyalist(r.FormValue("email"))
		if found == true {
			statusData := PageData{
				PageTitle: "Registration Status",
				Message:   "This Hard Topper has already been registered in the loyalty program!",
			}
			//Display Operation Status Page to User
			opstatus(w, r, statusData)
		} else {
			r.ParseForm()
			statusData := PageData{
				PageTitle: "Registration Status",
				Message:   "The email address " + r.FormValue("email") + " has been registered into the Loyalty Program!",
			}

			//write to mongo
			firstname := r.FormValue("firstname")
			lastname := r.FormValue("lastname")
			email := r.FormValue("email")
			password := r.FormValue("password")

			registerUser(firstname, lastname, email, password)

			//Display Operation Status Page to User
			session.Values["authenticated"] = true
			session.Values["email"] = r.FormValue("email")
			session.Save(r, w)
			opstatus(w, r, statusData)
		}
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method, "on URL:", r.URL)
	session, _ := store.Get(r, "cookie-name")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./static/login.html")
		t.Execute(w, nil)
	} else {
		found := checkLoyalist(r.FormValue("email"))
		if found == true {

			// Check Credentials
			access := checkCredentials(r.FormValue("email"), r.FormValue("password"))
			if access == true {
				fmt.Println("Access Granted")
				statusData := PageData{
					PageTitle: "Login Status",
					Message:   "You are now logged in to the Loyalty Program",
				}
				session.Values["authenticated"] = true
				session.Values["email"] = r.FormValue("email")
				session.Save(r, w)
				opstatus(w, r, statusData)
			} else {
				fmt.Println("Access Denied")
				statusData := PageData{
					PageTitle: "Login Status",
					Message:   "You could not be logged in with your account. Please try again!",
				}
				opstatus(w, r, statusData)
			}

		} else {
			fmt.Println("Account Not Found")
			statusData := PageData{
				PageTitle: "Login Status",
				Message:   "Your account could not be located. Please try another email address or register!",
			}
			opstatus(w, r, statusData)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method, "on URL:", r.URL)
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/", 302)
}

func loyaltystatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method, "on URL:", r.URL)
	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		fmt.Println("User not Authenticated - Forbidden!")
		statusData := PageData{
			PageTitle:  "",
			Message:    "Forbidden!",
			SiteAction: "Login",
			ActionPage: "/login",
		}
		if r.Method == "GET" {
			t, _ := template.ParseFiles("./static/loyaltystatus.html")
			t.Execute(w, statusData)
		}
		return
	}

	client, err := getMongoClient(mongoHost, mongoUser, mongoPass, mongoTLS)
	collection := client.Database("ht-loyalty").Collection("registrations")

	loginFilter := bson.D{{"email", session.Values["email"]}}

	var result Loyalist

	err = collection.FindOne(context.TODO(), loginFilter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	statusData := PageData{
		PageTitle:  "Loyalty Status",
		Message:    "Thank you for being a loyalty customer. You're current status level is " + fmt.Sprintf("%v", result.LoyaltyTier) + "!",
		SiteAction: "Logout",
		ActionPage: "/logout",
	}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./static/loyaltystatus.html")
		t.Execute(w, statusData)
	}

}

// end of page handlers

func mongoCheck(mongoHost string, mongoUser string, mongoPass string, mongoTLS string) {
	client, err := getMongoClient(mongoHost, mongoUser, mongoPass, mongoTLS)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func getMongoClient(mongoHost string, mongoUser string, mongoPass string, mongoTLS string) (*mongo.Client, error) {

	
	//mongoTLS string is required on DocumentDB. If running against DocumentDB ensure that the MONGO_TLS enviornment variable is not an empty string!
	if mongoTLS != "" {
		certString = "/?ssl=true&ssl_ca_certs=rds-combined-ca-bundle.pem&replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false"
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUser + ":" + mongoPass + "@" + mongoHost + ":27017" + certString)
	fmt.Println("Connection String is: " + "mongodb://" + mongoUser + ":" + mongoPass + "@" + mongoHost + ":27017" + certString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client, clientError
}

func checkLoyalist(email string) (found bool) {
	client, err := getMongoClient(mongoHost, mongoUser, mongoPass, mongoTLS)
	collection := client.Database("ht-loyalty").Collection("registrations")
	emailFilter := bson.D{{"email", email}}

	var result Loyalist

	found = false
	err = collection.FindOne(context.TODO(), emailFilter).Decode(&result)
	if err == nil {
		found = true
	}

	fmt.Println("Account with email address " + email + " was found " + strconv.FormatBool(found))
	return found
}

func checkCredentials(email string, password string) (access bool) {
	client, err := getMongoClient(mongoHost, mongoUser, mongoPass, mongoTLS)
	collection := client.Database("ht-loyalty").Collection("registrations")

	//convert plain text password to hash and compare with existing hash values
	passwordBytes := []byte(password)
	loginFilter := bson.D{{"email", email}}

	var result Loyalist

	access = false
	err = collection.FindOne(context.TODO(), loginFilter).Decode(&result)
	if err == nil {
		access = true
	}

	access = comparePasswords(result.Password, passwordBytes)

	return (access)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func registerUser(firstname string, lastname string, email string, password string) {
	client, err := getMongoClient(mongoHost, mongoUser, mongoPass, mongoTLS)
	collection := client.Database("ht-loyalty").Collection("registrations")

	password = hashAndSalt([]byte(password))
	//fmt.Println(password)
	entry := Loyalist{firstname, lastname, email, password, "silver"}

	insertResult, err := collection.InsertOne(context.TODO(), entry)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

}

func opstatus(w http.ResponseWriter, r *http.Request, messageData PageData) {
	fmt.Println("method:", r.Method, "on URL:", r.URL)
	session, _ := store.Get(r, "cookie-name")
	t, _ := template.ParseFiles("./static/internal_opstatus.html")
	if r.Method == "POST" {
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			t, _ = template.ParseFiles("./static/external_opstatus.html")
		}
		t.Execute(w, messageData)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 301)
}

func main() {
	// check DB Connection
	mongoCheck(mongoHost, mongoUser, mongoPass, mongoTLS)

	//web
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/loyalty", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path += ".html"
		fileServer.ServeHTTP(w, r)
	})
	http.HandleFunc("/healthz", healthHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/loyaltystatus", loyaltystatusHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
