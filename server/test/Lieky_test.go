package test

import (
	"fmt"
	"github.com/malaschitz/plamienok/server/db"
	"testing"
)

func TestLiek(t *testing.T) {
	db.ImportCache("../../data")
	fmt.Println(db.Lieky.Data[100])
	fmt.Println(db.Diagnozy.Data[100])
	if len(db.Lieky.Data) < 100000 || len(db.Diagnozy.Data) < 10000 {
		t.Fatalf("Bolo nacitanych iba %d liekov a %d diagnoz.", len(db.Lieky.Data), len(db.Diagnozy.Data))
	}
}
