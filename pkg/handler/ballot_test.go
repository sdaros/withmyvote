package handler

import (
	"os"
	"testing"

	"github.com/sdaros/withmyvote/pkg/app"
	"github.com/sdaros/withmyvote/pkg/db"
)

var d db.DB

func TestMain(m *testing.M) {
	d = db.Open()
	d.Store(db.TestData.ID, db.TestData)
	os.Exit(m.Run())
}

// TODO
func TestFetch(t *testing.T) {
	val, ok := d.Load("employee-benefits")
	if !ok {
		t.Errorf("couldn't load issue")
	}
	if val.(*app.Issue).ID != "employee-benefits" {
		t.Errorf("expected id to be employee-benefits, got %#+v", val)
	}
}
