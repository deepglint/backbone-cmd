package controller

import (
	"testing"

	"github.com/deepglint/backbone-cmd/util"
)

func TestBoltDb(t *testing.T) {
	bc := util.NewBoltDBController("test.db")
	e := bc.InitDB()
	defer bc.CloseDB()
	if e != nil {
		t.Error(e.Error())
		return
	}
	e = bc.NewTable("test1")
	if e != nil {
		t.Error(e.Error())
		return
	}
	bc.SetValue([]byte("test1"), []byte("1"), []byte("hello"))
	bc.SetValue([]byte("test1"), []byte("2"), []byte("world"))
}
