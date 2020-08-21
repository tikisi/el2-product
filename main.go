package main
import (
	"net/http"
	"html/template"
	"os/exec"
	"io"
	"fmt"
	"strings"
	"os"
	"log"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("index.tpl"))
		tpl.Execute(w, nil)
	} else if r.Method == "POST" {
		tpl := template.Must(template.ParseFiles("index.tpl"))

		r.ParseForm()
		var inputText =  r.Form["comment"][0]
		escapedText := strings.NewReplacer(
			"\r\n", " ",
			"\r", " ",
			"\n", " ",
		).Replace(inputText)

		//cmd := exec.Command("python3" ,"exec.py")
		cmd := exec.Command("./exec")
		stdin, _ := cmd.StdinPipe()
		io.WriteString(stdin, escapedText)
		stdin.Close()
		ColorizedText, _ := cmd.Output()

		fmt.Print(ColorizedText)

		m := map[string]template.HTML {
			"InputText": template.HTML(inputText),
			"ColorizedText": template.HTML(ColorizedText),
		}
		tpl.Execute(w, m)
	}
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":" + port, nil)
}
