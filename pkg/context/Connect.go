package context

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"net/http"
)

func (c *Context) Connect(retry bool) error {
	if retry {
		err := backoff.Retry(func() error {
			resp, err := c.Client.Get(fmt.Sprintf("%s/healthz", c.ApiURL))

			if err != nil {
				return err
			} else {
				if resp.StatusCode == http.StatusOK {
					return nil
				} else {
					return errors.New("failed to authenticate against the smr-agent")
				}
			}

			return errors.New("context not saved")
		}, backoff.NewExponentialBackOff())

		if err != nil {
			return err
		}

		return nil
	} else {
		resp, err := c.Client.Get(fmt.Sprintf("%s/healthz", c.ApiURL))

		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusOK {
			return nil
		} else {
			return errors.New("failed to authenticate against the smr-agent")
		}
	}
}
