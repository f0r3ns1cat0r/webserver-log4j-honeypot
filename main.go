package main

import (
	"log"
	"fmt"
	"regexp"
	"strings"
        "net/http"
	"io/ioutil"
	"github.com/schadom/webserver-log4j-honeypot/extractor"
)

func analyzeReq(text string) {
	log.Printf("Testing text: %s\n", text)

	pattern := regexp.MustCompile(`\${jndi:(.*)}`)
	finder := extractor.NewFinder(pattern)
	injections := finder.FindInjections(text)

	for _, url := range injections {
		log.Printf("Fetching payload for: jndi:%s", url.String())

		files, err := extractor.FetchFromLdap(url)
		if err != nil {
			log.Printf("Failed to fetch class from %s", url)
			continue
		}

		for _, filename := range files {
			log.Printf("Saved payload to file %s\n", filename)
		}
	}
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	reqURL := r.URL.Path;
	log.Printf("New request received to %s\n", reqURL)
	analyzeReq(string(reqURL));

	for k,v := range r.Header {
		_ = k
		str := strings.Join(v," ")
		analyzeReq(string(str))
	}
	body, err := ioutil.ReadAll(r.Body);

	if err != nil {
		log.Fatal(err)
	}
	analyzeReq(string(body));
	fmt.Fprintf(w, "<p>Hello world</p>")
}

func main() {
	httpPort := 8888
	fmt.Println("Starting Webserver on port", httpPort)

	http.HandleFunc("/", reqHandler)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Println(err)
	}
}
