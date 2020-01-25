package cli

import (
	log "github.com/sirupsen/logrus"

	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/leonkaihao/nycab/pkg/config"
)

type count struct {
	conf *config.Cli
}

func (c *count) Handle(args []string) (resp string, err error) {

	if c.conf == nil {
		errMsg := "config error, failed to connect to api service"
		return errMsg, errors.New(errMsg)
	}
	url := "http://" + c.conf.APIHost + ":" + c.conf.APIPort + "/v1/cabs/pickups/count"
	url += reassembleArgs(args)
	log.Infoln("GET", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New(string(data))
	}
	return string(data), nil
}

func reassembleArgs(args []string) string {
	ret := ""
	for i, arg := range args {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			continue // ignore illegal args
		}
		switch kv[0] {
		case "medallions":
			fallthrough
		case "date":
			fallthrough
		case "refresh":
			if i == 0 {
				ret += ("?" + arg)
			} else {
				ret += ("&" + arg)
			}
		default:
			break
		}
	}
	return ret
}

// NewCountCmd ...
func NewCountCmd(conf *config.Cli) *Cmdline {
	return &Cmdline{
		Cmd:         "count",
		Example:     "count medallion=medallion1,medallion2,medallion... date=2013-01-12 refresh=true",
		Description: "count medallions pickup in a day ",
		Handler:     &count{conf},
	}
}
