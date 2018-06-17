package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/sdaros/withmyvote/pkg/db"
)

var (
	createIssue = `
<div id="newBallot">
  <p>Looks like an Issue with that name doesn't exist yet</p>
  <p>Just enter markdown and you're all cool!</p>
  <textarea rows="6" cols="80">
# Which employee benefit would you most want?

[-] Extra 5 days of vacation per year
[-] Home office 8 days in a calendar month
[-] A brand new electric car!!1
  </textarea></br>
  <button>Create New Issue</button>
</div>
`
	viewIssue = `
  <header class="header">
    <div class="header-content header-content-ns">
      <h3 class="issue-name issue-name-ns">
        <span class="issue-name-span issue-name-span-ns">
          Which employee benefit would you most want?
        </span>
      </h3>
      <h4 class="issue-choose issue-choose-ns">Choose from one of the options below</h4>
    </div>
  </header>
  <main class="main-content main-content-ns">
    {{range .Candidates}}
    <div class="candidate">
      <div class="candidate-id candidate-id-ns">
        <span>{{.ID}}</span>
      </div>
	  <input class="candidate-input candidate-input-ns" type="radio" id="{{$.ID}}" name="{{$.ID}}" value="{{$.ID}}-{{.ID}}"/>
	  <label for="{{$.ID}}-{{.ID}}" class="candidate-label candidate-label-ns">
		{{.Text}}
      </label>
    </div>
    {{end}}
  </main>
`
	yourVoteHasBeenSubmitted = []byte(`
<p>Your vote has been submitted! Thanks :-)</p>
`)
	yourVoteHasBeenRetracted = []byte(`
<p>Your vote has been retracted!</p>
`)
)

// Ballot represents an issue by rendering a HTML form that Voters
// can use to cast their votes
type Ballot struct{}
type fingerprint string
type HTML []byte

func (b *Ballot) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		log.Println("POST")
		if err := b.CastVote("foo", "bar", 2); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write(yourVoteHasBeenSubmitted)
		return
	case "DELETE":
		log.Println("DELETE")
		if err := b.RetractVote("foo", "bar", 2); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(yourVoteHasBeenRetracted)
		return
	default:
		log.Println("DEFAULT")
		html, err := b.FetchOrCreate(req.URL.Path)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.Write(html)
		return
	}
}

// New creates a new Ballot
func (b *Ballot) New() error {
	return nil
}

// Fetch gets
func (b *Ballot) FetchOrCreate(key string) (content HTML, err error) {
	d := db.Open()
	issue, ok := db.Load(key)
	if !ok {
		// Issue does not exist, provide user with
		// a form to create a new issue
		return []byte(fmt.Sprintf(base, createIssue)), nil
	}
	tmpl, err := template.New("ballot").Parse(viewIssue)
	if err != nil {
		return nil, err
	}
	data := new(bytes.Buffer)
	if err := tmpl.Execute(data, issue); err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf(base, data)), nil
}

func (b *Ballot) CastVote(voterID fingerprint, issueID string, candidate int) error {

	/*
		get results from req.Formdata
		db.Store()
	*/
	return nil
}

func (b *Ballot) RetractVote(voterID fingerprint, issueID string, candidate int) error {

	/*
		get results from req.Formdata
		db.Store()
	*/
	return nil
}
