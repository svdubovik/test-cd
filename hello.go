package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	name := "John"

	tmpl.Execute(w, name)

}

func buildSQL(email string) string {
	return fmt.Sprintf("SELECT * FROM users WHERE email='%s' AND Password='123456';", email)
}

func main() {
	buildSQL("test@test.com")
	cred := "KjasdlkjapoIKLlka98098sdf012U/rL2sLdBqOHQUlt5Z6kCgKGDyCFA=="
	secret := "12345"
	fmt.Println(cred)
	fmt.Println(secret)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
