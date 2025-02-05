package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func auth(login, password string) bool {
	return login == "lok" && password == "lok"
}

func authHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(res, form)
		return
	}

	if req.Method == http.MethodPost {
		if err := req.ParseForm(); err != nil {
			res.Write([]byte(err.Error()))
			return
		}
		if auth(req.FormValue("login"), req.FormValue("password")) {
			res.Write([]byte("Hello" + req.FormValue("login") + "!"))
			return
		}
	}

	http.Error(res, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func mainHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome buddy!\n"))
	body := fmt.Sprintln("Request Method:", req.Method)
	body += fmt.Sprintln("Request headers:")
	body += fmt.Sprintln()
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: \n", k)
		for _, v := range v {
			body += fmt.Sprintf("		%s\n", v)
		}
	}

	if err := req.ParseForm(); err != nil {
		res.Write([]byte(err.Error()))
		return
	}

	body += fmt.Sprintln()
	body += fmt.Sprintln("Request querry:")
	for k, v := range req.Form {
		body += fmt.Sprintf("   %s: %s\n", k, v)
	}

	res.Write([]byte(body))
}

func jsonHandler(res http.ResponseWriter, req *http.Request) {
	swaffard := User{Name: "Sergant", Age: 18}

	jsn, err := json.Marshal(swaffard)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusOK)

	res.Write(jsn)
}

func apiHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome to the api page!"))
}

func apiAuthHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to api/auth page..."))
}

func clientHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to client page!"))
}

func dummePrinter(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("1"))
	io.WriteString(res, "2")
	fmt.Fprintln(res, "3")
}

func middlewarelog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Side log from middleware....\n"))
		next.ServeHTTP(w, r)
	})
}

func middlewareprint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PRINT....\n"))
		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("dummy service is up!")

	mux := http.NewServeMux()

	mux.HandleFunc("/", authHandler)
	mux.HandleFunc("GET /main", mainHandler)
	http.Handle("GET /main1", middlewarelog(middlewareprint(http.HandlerFunc(mainHandler))))
	mux.HandleFunc("GET /dummy", dummePrinter)
	mux.HandleFunc("GET /client/", clientHandler)
	mux.HandleFunc("GET /api/", apiHandler)
	mux.HandleFunc("GET /api/auth", apiAuthHandler)
	mux.HandleFunc("GET /json", jsonHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
