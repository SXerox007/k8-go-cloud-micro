package templates

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// get template dir
func getTemplatesDir() string {
	env := os.Getenv("ENV_NAME")
	if env == "" {
		env = "local"
	}
	return env
}

// parse templete
func ParseTemplate(path string) (*template.Template, error) {
	//reads all files
	log.Println("Path", path)
	p, er := template.New("mustache").Delims("<<", ">>").ParseGlob(getTemplatesDir() + "static" + "/templates/[a-z]*.mustache")
	log.Println("Error:----", er)
	return template.Must(p, er), er
}

// output html
func OutputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
