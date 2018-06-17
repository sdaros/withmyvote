package app

import (
	"net/http"
	"time"

	"github.com/sdaros/withmyvote/pkg/handler"
)

var (
	// Routes to handle
	Routes = []Route{
		Route{"/", handler.Root{}},
	}
)

// Route is cool
type Route struct {
	Path    string
	Handler http.Handler
}

// Issue is voted on
type Issue struct {
	ID          string
	Description string
	CreatedAt   time.Time
	Candidates  []*Candidate
}

// Candidate in issue
type Candidate struct {
	ID   int
	Text string
}
