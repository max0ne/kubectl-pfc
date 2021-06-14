package url

import (
	"fmt"
	"regexp"
)

var regex = regexp.MustCompile(`^((.+?)://)?([^:/]+)(:(\d+))?(.*?)$`)

func Parse(u string) (scheme, host, port, query string, err error) {
	matches := regex.FindStringSubmatch(u)
	if len(matches) != 7 {
		err = fmt.Errorf("Unable to parse url %s: unexpected submatches %v", u, matches)
		return
	}

	scheme = matches[2]
	host = matches[3]
	port = matches[5]
	query = matches[6]
	return
}
