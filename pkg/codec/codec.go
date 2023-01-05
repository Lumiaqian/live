package codec

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func CalcMD5(messages ...string) string {
	h := md5.New()
	for _, msg := range messages {
		io.WriteString(h, msg)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Base64Decode(str string) string {
	reader := strings.NewReader(str)
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 1024)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		dst += string(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return dst
}
