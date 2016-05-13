package common

import "net/url"

type Option struct {
	URL *url.URL
	Raw string
}

func ParseOption(rawUrl string) (Option, error) {
	url, err := url.Parse(rawUrl)

	if err != nil {
		return Option{Raw: rawUrl}, nil
	}

	return Option{URL: url, Raw: rawUrl}, nil
}
