// generated by wsp, DO NOT EDIT.

package main

import "net/http"
import "time"
import "github.com/simplejia/connsvr/test/logicsvr/controller"

func init() {
	http.HandleFunc("/Demo/Msgs", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(controller.Demo)
		defer func() {
			e = recover()
		}()
		c.Msgs(w, r)
	})

	http.HandleFunc("/Demo/Pub", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(controller.Demo)
		defer func() {
			e = recover()
		}()
		c.Pub(w, r)
	})

	http.HandleFunc("/Demo/Var", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		_ = t
		var e interface{}
		c := new(controller.Demo)
		defer func() {
			e = recover()
		}()
		c.Var(w, r)
	})

}