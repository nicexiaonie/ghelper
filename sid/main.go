package sid

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

var sidManage sync.Map

func Register() {
	sidManage.Store(getGID(), uniqueId())
}
func Destroy() {
	sidManage.Delete(getGID())
}
func Set(value string) {
	sidManage.Store(getGID(), value)
}
func Get() string {
	value, ok := sidManage.Load(getGID())
	if ok && value != nil {
		return value.(string)
	}
	return ""
}

//生成32位md5字串
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func uniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
