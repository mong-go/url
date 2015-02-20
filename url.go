package url

import (
	"fmt"
	"net/url"
	"strings"
)

// URL wraps url.URL providing extra methods for mongodb urls
type URL struct {
	*url.URL
}

// Parse parses a mongodb url provided this format:
//
//		mongodb://[user:pass@]host[:port]/database
//
func Parse(urlStr string) (*URL, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	return &URL{
		URL: u,
	}, nil
}

// Database name
func (u URL) Database() string {
	return u.URL.Path[1:]
}

// User returns the URL Userinfo
func (u URL) User() *url.Userinfo {
	return u.URL.User
}

// Scheme mongodb
func (u URL) Scheme() string {
	return u.URL.Scheme
}

// Host returns the ip/domain portion of the host only
func (u URL) Host() string {
	s := strings.Split(u.URL.Host, ":")
	return s[0]
}

// Port parses the port from the host, returning mongodb default if there is no
// port specified
func (u URL) Port() string {
	s := strings.Split(u.URL.Host, ":")
	if len(s) == 1 {
		return "27017"
	}

	return s[1]
}

// ShortString returns the compiled URL minus the database path
func (u URL) ShortString() string {
	f := "%s://%s"
	if us := u.User(); us != nil {
		f = fmt.Sprintf("%s://%s@%s", "%s", us.String(), "%s")
	}

	return fmt.Sprintf(f, u.URL.Scheme, u.URL.Host)
}

// String calls u.URL.String()
func (u URL) String() string {
	return u.URL.String()
}
