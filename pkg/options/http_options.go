package options

import "time"

var _ IOptions = (*HTTPOptions)(nil)

type HTTPOptions struct {
	Network string        `json:"network"`
	Addr    string        `json:"addr"`
	Timeout time.Duration `json:"timeout"`
}

func NewHTTPOptions() *HTTPOptions {
	return &HTTPOptions{
		Network: "tcp",
		Addr:    "0.0.0.0:38443",
		Timeout: 30 * time.Second,
	}
}

func (o *HTTPOptions) Validate() []error {
	errs := []error{}

	return errs
}
