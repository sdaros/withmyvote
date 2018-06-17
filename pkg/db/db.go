package db

import (
	"sync"
	"time"

	"github.com/sdaros/withmyvote/pkg/app"
)

var (
	// TestData to simplify testing
	TestData = &app.Issue{
		ID:          "employee-benefits",
		Description: "Which employee benefit would you most want?",
		CreatedAt:   time.Now(),
		Candidates: []*app.Candidate{
			&app.Candidate{
				ID:   1,
				Text: "Extra 5 days of vacation per year",
			},
			&app.Candidate{
				ID:   2,
				Text: "Home office 8 days in a calendar month",
			},
			&app.Candidate{
				ID:   3,
				Text: "A brand new electric car!!1!",
			},
		},
	}
)

type DB interface {
	//Delete(key interface{})
	Load(key interface{}) (value interface{}, ok bool)
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	//Range(f func(key, value interface{}) bool)
	Store(key, value interface{})
}

type db struct {
	state *sync.Map
}

func Open() DB {
	return &db{state: &sync.Map{}}
}
func (d db) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	return d.state.LoadOrStore(key, value)
}

func (d db) Load(key interface{}) (value interface{}, ok bool) {
	return d.state.Load(key)
}

func (d db) Store(key, value interface{}) {
	d.state.Store(key, value)
}
