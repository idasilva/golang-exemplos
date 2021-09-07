package ngnix

import "github.com/idasilva/golang-exemplos/pattern/proxy/app"

type Ngnix struct {
	app    *app.Application
	maxAllowedRequest int
	rateLimiter     map[string]int
}

func (n *Ngnix)HandlerRequest(url, method string) (int,string)  {
	allowed := n.checkReteLimiting(url)
	if !allowed{
		return 403, "Not Allowed"
	}

	return n.app.HandlerRequest(url,method)
}

func (n *Ngnix) checkReteLimiting(url string) bool{
	if n.rateLimiter[url] == 0{
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url]  > n.maxAllowedRequest{
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func NewNginx() *Ngnix{
	return &Ngnix{
		app:               &app.Application{},
		maxAllowedRequest: 2,
		rateLimiter: make(map[string]int),
	}
}
