package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var compileTimeTemplates *template.Template

var validPath = regexp.MustCompile(
	"^/()?$")

func init() {

	workdir, _ := os.Getwd()
	basedir, _ := os.Executable()

	workdirs := strings.Split(workdir, "\\")
	if len(workdirs) == 1 {
		workdirs = strings.Split(workdir, "/")
	}

	basedirs := strings.Split(basedir, "\\")
	if len(basedirs) == 1 {
		basedirs = strings.Split(basedir, "/")
	}

	currentworkdir := workdirs[len(workdirs)-1]
	currentbasedir := basedirs[len(basedirs)-1]

	if currentworkdir != currentbasedir {
		fmt.Printf("you are in %s, please cd to %s\n", currentworkdir, currentbasedir)
		println("wrong working directory please use 'go run .' AND NOT 'go run main.go' in root directory (forum)")
		os.Exit(0)
	}

	compileTimeTemplates = template.Must(
		template.ParseFiles(
			"pages/main.html",
			"pages/notfound.html",
			"pages/protected.html",
			"pages/internalerror.html",
		))
}

func HomeRedirector(w http.ResponseWriter, r *http.Request, title string) {
	println("redirected to /main")
	http.Redirect(w, r, "/main", http.StatusFound)
}

func MainHandler(w http.ResponseWriter, r *http.Request, title string) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusInternalServerError)
		renderTemplate(w, "notfound", &Page{Title: "404"})
		return
	}
	renderTemplate(w, "main", &Page{Title: "Forum"})
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := compileTimeTemplates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		renderTemplate(w, "internalerror", &Page{Title: "500"})
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request, title string) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusForbidden)
		renderTemplate(w, "protected", &Page{Title: "Protected"})
		return
	}
	fmt.Fprintf(w, "success")
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			w.WriteHeader(http.StatusNotFound)
			renderTemplate(w, "notfound", &Page{Title: "404"})
			return
		}
		fn(w, r, m[len(m)-1])
	}
}
