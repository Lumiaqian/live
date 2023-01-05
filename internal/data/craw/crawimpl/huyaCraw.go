package crawimpl

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"live/internal/model"
	"live/pkg/code/errorx"
	"live/pkg/codec"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"github.com/zeromicro/go-zero/core/logx"
)

type huya struct {
	httpClient *http.Client
}

// GetLiveUrl implements craw.HuYaCraw
func (h *huya) GetLiveUrl(ctx context.Context, roomId string) (*model.HuYaRoom, error) {
	roomUrl := "https://m.huya.com/" + roomId
	request, err := http.NewRequest("GET", roomUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36")
	response, err := h.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	reg := regexp.MustCompile("<script> window.HNF_GLOBAL_INIT = (.*)</script>")
	submatch := reg.FindStringSubmatch(string(result))
	if submatch == nil || len(submatch) < 2 {
		return nil, errors.Wrap(err, "查询失败！")
	}
	return extractInfo(submatch[1])
}

func extractInfo(content string) (*model.HuYaRoom, error) {
	parse := gjson.Parse(content)
	streamInfo := parse.Get("roomInfo.tLiveInfo.tLiveStreamInfo.vStreamInfo.value")
	nickName := parse.Get("roomInfo.tProfileInfo.sNick").String()
	var urls []string
	streamInfo.ForEach(func(key, value gjson.Result) bool {
		urlPart := value.Get("sStreamName").String() + "." + value.Get("sFlvUrlSuffix").String() + "?" + value.Get("sFlvAntiCode").String()
		urls = append(urls, value.Get("sFlvUrl").String()+"/"+urlPart)
		return true
	})
	liveLineUrl := parse.Get("roomProfile.liveLineUrl").String()
	logx.Infof("liveLineUrl:", liveLineUrl)
	liveUrlByteData, err := base64.StdEncoding.DecodeString(liveLineUrl)
	if err != nil {
		return nil, errorx.NewErrMsg("未开播或直播间不存在")
	}
	liveUrl, err := live(liveUrlByteData)
	if err != nil {
		return nil, errorx.NewErrMsg("未开播或直播间不存在")
	}
	liveUrl = strings.ReplaceAll("https:"+liveUrl, "hls", "flv")
	liveUrl = strings.ReplaceAll(liveUrl, "m3u8", "flv")
	liveUrl = strings.ReplaceAll(liveUrl, "&ctype=tars_mobile", "")
	return &model.HuYaRoom{
		Urls:    urls,
		Name:    nickName,
		LiveUrl: liveUrl,
	}, nil
}

func live(byteData []byte) (string, error) {
	liveUrl := string(byteData)
	strs := strings.Split(liveUrl, "?")
	if len(strs) <= 1 {
		return "", errors.New("未开播或直播间不存在")
	}
	r := strings.Split(strs[0], "/")
	reg := regexp.MustCompile(`.(flv|m3u8)`)
	s := reg.ReplaceAllString(r[len(r)-1], "")
	c := strings.SplitN(strs[1], "&", 4)
	c1 := []string{}
	for _, str := range c {
		if str != "" {
			c1 = append(c1, str)
		}
	}
	n := make(map[string]string)
	for _, str := range c1 {
		cs := strings.Split(str, "=")
		n[cs[0]] = cs[1]
	}
	fm, err := url.QueryUnescape(n["fm"])
	if err != nil {
		return "", err
	}
	u := codec.Base64Decode(fm)
	p := strings.Split(u, "_")[0]
	f := strconv.Itoa(int(time.Now().UnixNano()))
	l := n["wsTime"]
	t := "0"
	hs := []string{p, t, s, f, l}
	h := strings.Join(hs, "_")
	m := codec.CalcMD5(h)
	y := c1[len(c1)-1]
	url := fmt.Sprintf("%s?wsSecret=%s&wsTime=%s&u=%s&seqid=%s&%s", strs[0], m, l, t, f, y)
	return url, nil
}

func newHuya(dc *dataCraw) *huya {
	return &huya{dc.httpClient}
}
