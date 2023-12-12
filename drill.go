package drill

import (
	"net/http"
	"net/url"
)

type drill struct {
	url *url.URL

	method string
}

func new(urls string, mthd string) *drill {
	parseurls, err := url.Parse(urls)
	if err != nil {
		panic(err)
	}
	return &drill{
		url:    parseurls,
		method: mthd,
	}
}

func Get(url string) {
	client := new(url, "GET")
	http.Get(client.url.String())
}

func Post(url string) {
	client := new(url, "POST")
	http.Post(client.url.String(), "", nil)
}

func Put(url string) {
	client := new(url, "PUT")
	r, err := http.NewRequest(client.method, client.url.String(), nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Do(r)
}

func Delete(url string) {
	client := new(url, "DELETE")
	r, err := http.NewRequest(client.method, client.url.String(), nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Do(r)
}

func Head(url string) {
	client := new(url, "HEAD")
	r, err := http.NewRequest(client.method, client.url.String(), nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Do(r)
}
