package mongodb

import (
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {
	p := &Provider{}
	p.SessionInit(5, "mongodb://localhost:27017/test")
	if p.mgoSession == nil {
		t.Error("client err")
	}
	ss, err := p.SessionRead("foo-bar")
	if err != nil {
		t.Error(err)
	}
	if !p.SessionExist("foo-bar") {
		t.Error("SessionExist err")
	}
	err = ss.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}
	if ss.Get("foo") != "bar" {
		t.Error("Get err")
	}
	err = ss.Delete("foo")
	//err = ss.Flush()
	if err != nil {
		t.Error(err)
	}
	if ss.Get("key") == "value" {
		t.Error("Delete/Flush err")
	}
	if ss.SessionID() != "foo-bar" {
		t.Error("id err")
	}
	ss.Set("foo1", "bar1")
	var w http.ResponseWriter
	ss.SessionRelease(w)
	new, e := p.SessionRead("foo-bar")
	if new == nil || e != nil {
		t.Error(e)
	}
	if !p.SessionExist("foo-bar") {
		t.Error("SessionExist err")
	}
	newS, er := p.SessionRegenerate("foo-bar", "bar-foo")
	if er != nil || newS == nil {
		t.Error("SessionRegenerate err")
	}
	if p.SessionExist("foo-bar") {
		t.Error("SessionExist err")
	}
	if p.SessionAll() <= 0 {
		t.Error("SessionAll err")
	}

	t.Log("sleep 6s")
	time.Sleep(6 * time.Second)
	//p.SessionGC()
	p.SessionDestroy("bar-foo")
	if p.SessionExist("bar-foo") {
		t.Error("SessionExist err")
	}
}
