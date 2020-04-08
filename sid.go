package ghelper

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"runtime"
	"strconv"
	"sync"
)

type sidModel struct {
	sidManage sync.Map
}

var Sid = sidModel{}

func (current *sidModel) Register() {
	current.sidManage.Store(current.getGID(), current.uniqueId())
}
func (current *sidModel) Destroy() {
	current.sidManage.Delete(current.getGID())
}
func (current *sidModel) Set(value string) {
	current.sidManage.Store(current.getGID(), value)
}
func (current *sidModel) Get() string {
	value, ok := current.sidManage.Load(current.getGID())
	if ok && value != nil {
		return value.(string)
	}
	return ""
}

//生成32位md5字串
func (current *sidModel) getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func (current *sidModel) uniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return current.getMd5String(base64.URLEncoding.EncodeToString(b))
}

func (current *sidModel) getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
