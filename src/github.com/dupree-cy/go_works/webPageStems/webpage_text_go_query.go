package main

//
// by Dupree
//
//    Resource
//     goquery suggested
//       http://stackoverflow.com/questions/12883079/go-parse-html-table
//     removing white spaces
//       http://stackoverflow.com/questions/13737745/split-a-string-on-whitespace-in-go
//     regex on alpha numeric
//       http://stackoverflow.com/questions/336210/regular-expression-for-alphanumeric-and-underscores
//
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/reiver/go-porterstemmer"
	"log"
	"regexp"
	"strings"
)

func main() {
	buffer_no_word := 20000 // Max number of words on the page
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	url := "http://revolutionanalytics.com/why-revolution-analytics"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	//doc_nodes := doc.Selection.Size()
	doc_text := doc.Selection.Text()
	var doc_fields []string
	var doc_fields_stem []string
	doc_fields = make([]string, buffer_no_word)
	doc_fields_stem = make([]string, buffer_no_word)
	doc_fields = strings.Fields(doc_text)
	var j int
	j = 1
	for i := 0; i < len(doc_fields); i++ { // report words
		if re.MatchString(doc_fields[i]) {
			stem := porterstemmer.StemString(doc_fields[i])
			doc_fields_stem[j] = stem
			fmt.Println(j, "word=", doc_fields[i], "stem=", stem)
			j++
		}
	}
	return
}
