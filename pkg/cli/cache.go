package cli

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/leonkaihao/nycab/pkg/config"
)

type cache struct {
	conf *config.Cli
}

func (c *cache) Handle(args []string) (resp string, err error) {
	if c.conf == nil {
		errMsg := "config error, failed to connect to api service"
		return errMsg, errors.New(errMsg)
	}
	url := "http://" + c.conf.APIHost + ":" + c.conf.APIPort + "/v1/cabs/pickups/count/cache"
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", errors.New(string(data))
	}
	return string(data), nil
}

// NewCacheCmd ...
func NewCacheCmd(conf *config.Cli) *Cmdline {
	return &Cmdline{
		Cmd:         "cache",
		Example:     "cache clean",
		Description: "remove all cache",
		Handler:     &cache{conf},
	}
}
