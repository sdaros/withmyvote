package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	viewIssue = `
<div id="ballot">
  <form method="POST" id="{{.Issue.Id}}">
    {{range .Issue.candidates -}}
	  <div class="candidate">
		<input name="{{.Issue.Id}}" type="radio" value=""/>
		<em>{{.Id}}: </em>{{.Text}}</br>
	  </div>
    {{- end}}
	<input class="submit" name="submit" type="submit" value="Submit"/>
  </form>
</div>
`
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
	foo = `
  <header class="header">
    <div class="header-content header-content-l">
      <h3 class="f2 f1-m f-subheadline-l mausre lh-title mv0">
        <span class="issue-name-span">
          Which employee benefit would you most want?
        </span>
      </h3>
      <h4 class="issue-choose">Choose from one of the options below</h4>
    </div>
  </header>
  <main class="main-content main-content-l">
    {{range .Candidates}}
    <div class="candidate">
      <div class="candidate-id">
        <span>{{.Id}}</span>
      </div>
	  <input class="candidate-input" type="radio" id="{{$.Id}}" name="{{$.Id}}" value="{{$.Id}}-{{.Id}}"/>
	  <label for="{{$.Id}}-{{.Id}}" class="candidate-label candidate-label-ns">
		{{.Text}}
      </label>
    </div>
    {{end}}
  </main>
  <footer class="footer">
    <div class="footer-content footer-content-l">
      <h5>withmy.vote from @sdaros with &lt;3</h4>
    </div>
  </footer>
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
		html, err := b.Fetch()
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
func (b *Ballot) Fetch() (content HTML, err error) {
	/*
		issue, ok := db.Load(issueId)
		if !ok {
	*/
	if true {
		tmpl, err := template.New("ballot").Parse(foo)
		if err != nil {
			return nil, err
		}
		data := new(bytes.Buffer)
		issue := struct {
			Id          string
			Description string
			CreatedAt   time.Time
			Candidates  []struct {
				Id   int
				Text string
			}
		}{
			Id:          "employee-benefits",
			Description: "Which employee benefit would you most want?",
			CreatedAt:   time.Now(),
			Candidates: []struct {
				Id   int
				Text string
			}{
				{
					Id:   1,
					Text: "Extra 5 days of vacation per year",
				},
				{
					Id:   2,
					Text: "Home office 8 days in a calendar month",
				},
				{
					Id:   3,
					Text: "A brand new electric car!!1!",
				},
			},
		}
		if err := tmpl.Execute(data, issue); err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(base, data)), nil
	}
	return []byte(""), nil
	/*
		t.execute(createIssue)
		}
		t.execute(viewIssue)

	*/
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
