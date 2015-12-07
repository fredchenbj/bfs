package main

import (
	log "github.com/golang/glog"
	"net/http"
	"os"
	"strconv"
	"time"
)

// StartApi start api http listen.
func StartApi(addr string, d *Directory) {
	go func() {
		var (
			err      error
			serveMux = http.NewServeMux()
		)
		serveMux.Handle("/get", httpGetHandler{d: d})
		serveMux.Handle("/upload", httpUploadHandler{d: d})
		serveMux.Handle("/del", httpUploadHandler{d: d})
		if err = http.ListenAndServe(addr, serveMux); err != nil {
			log.Errorf("http.ListenAndServe(\"%s\") error(%v)", addr, err)
			return
		}
	}()
	return
}

// httpGetHandler http upload a file.
type httpGetHandler struct {
	d   *Directory
}

func (h httpGetHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	var (
		err              error
		key              int64
		vid, cookie      int32
		stores           []string
		res              Response
		params           = r.URL.Query()
	)
	if r.Method != "GET" {
		ret = http.StatusMethodNotAllowed
		http.Error(wr, "method not allowed", ret)
		return
	}
	defer HttpWriter(r, wr, time.Now(), res, &ret)
	if key, err = strconv.ParseInt(params.Get("key"), 10, 64); err != nil {
		log.Errorf("strconv.ParseInt(\"%s\") error(%v)", params.Get("key"), err)
		ret = http.StatusBadRequest
		return
	}
	if cookie, err = strconv.ParseInt(params.Get("cookie"), 10, 32); err != nil {
		log.Errorf("strconv.ParseInt(\"%s\") error(%v)", params.Get("cookie"), err)
		ret = http.StatusBadRequest
		return
	}
	if stores, vid, ret, err = h.d.Rstores(key, cookie); err != nil {
		log.Errorf("Rstores() error(%v", err)
		ret = http.StatusInternalServerError
		return
	}
	res.Vid = vid
	res.Stores = stores
	return
}

// httpUploadHandler http upload a file.
type httpUploadHandler struct {
	d *Directory
}

func (h httpUploadHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	var (
		err              error
		keys             []int64
		vid, cookie,num  int32
		stores           []string
		ret              int
		res              Response
		params           = r.URL.Query()
	)
	if r.Method != "POST" {
		ret = http.StatusMethodNotAllowed
		http.Error(wr, "method not allowed", ret)
		return
	}
	defer HttpWriter(r, wr, time.Now(), res, &ret)
	if num, err = strconv.ParseInt(params.Get("num"), 10, 32); err != nil {
		log.Errorf("strconv.ParseInt(\"%s\") error(%v)", params.Get("key"), err)
		ret = http.StatusBadRequest
		return
	}
	if keys, vid, cookie, stores, ret, err = h.d.Wstores(num); err != nil {
		log.Errorf("Wstores() error(%v)", err)
		ret = http.StatusInternalServerError
		return
	}
	res.Keys = keys
	res.Vid = vid
	res.Cookie = cookie
	res.Stores = stores
	return
}

// httpDelHandler
type httpDelHandler struct {
	d *Directory
}

func (h httpDelHandler) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	var (
		err              error
		keys             []int64
		vid, cookie,num  int32
		stores           []string
		ret              int
		res              Response
		params           = r.URL.Query()
	)
	if r.Method != "POST" {
		ret = http.StatusMethodNotAllowed
		http.Error(wr, "method not allowed", ret)
		return
	}
	defer HttpWriter(r, wr, time.Now(), res, &ret)
	if key, err = strconv.ParseInt(params.Get("key"), 10, 64); err != nil {
		log.Errorf("strconv.ParseInt(\"%s\") error(%v)", params.Get("key"), err)
		ret = http.StatusBadRequest
		return
	}
	if cookie, err = strconv.ParseInt(params.Get("cookie"), 10, 32); err != nil {
		log.Errorf("strconv.ParseInt(\"%s\") error(%v)", params.Get("cookie"), err)
		ret = http.StatusBadRequest
		return
	}
	if stores, vid, ret, err = h.d.Dstores(key, cookie); err != nil {
		log.Errorf("Dstores() error(%v", err)
		ret = http.StatusInternalServerError
		return
	}
	res.Vid = vid
	res.Stores = stores
	return
}
