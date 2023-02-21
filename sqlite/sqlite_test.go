package sqlite

import (
	"log"
	"os"
	"testing"

	"github.com/louisbranch/drake"
)

func TestDBInterface(t *testing.T) {
	var _ drake.Database = &DB{}

}

func testDB() (*DB, string) {
	tmpfile, err := os.CreateTemp("", "drake.db")
	if err != nil {
		log.Fatal(err)
	}
	name := tmpfile.Name()
	db, err := New(name)
	if err != nil {
		log.Fatal(err)
	}
	return db, name
}

func TestCreateSession(t *testing.T) {
	db, path := testDB()
	defer os.Remove(path)

	session := &drake.Session{}
	err := db.CreateSession(session)

	if err == nil {
		t.Errorf("wants error, got none")
	}
}
