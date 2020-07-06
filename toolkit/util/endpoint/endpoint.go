package endpoint

import (
    "fmt"
    "net"
    "net/url"
    "strings"

    "github.com/pkg/errors"
    "github.com/rs/zerolog/log"
)

// Helper Functions
func GetListener(endpoint string) (lis net.Listener, err error) {

    var target Target
    target, err = ParseTarget(endpoint)
    if err != nil {
        return
    }

    switch target.Scheme {
    case "unix":
        return net.Listen("unix", target.Path)
    case "tcp", "dns", "http", "https", "kubernetes":
        if target.Port == "" {
            target.Port = "0"
        }
        return net.Listen("tcp", fmt.Sprintf(":%s", target.Port))
    default:
        return nil, errors.New(fmt.Sprintf("unknown scheme: %s in endpoint: %s", target.Scheme, endpoint))
    }
}

type Target struct {
    Scheme string
    Host   string
    Port   string
    Path   string
}

// ParseTarget splits target into a Target struct containing scheme, host, port and path.
// If target is not a valid scheme://host:port/path, it returns error.
func ParseTarget(target string) (ret Target, err error) {
    if strings.HasPrefix(target, "unix://") {
        spl := strings.SplitN(target, "://", 2)
        ret.Scheme = spl[0]
        ret.Path = spl[1]
        return
    }
    target = strings.Replace(target, ":///", "://", 1)
    var u *url.URL
    if u, err = url.Parse(target); err == nil {
        ret.Scheme = u.Scheme
        ret.Path = u.Path
        if u.Host != "" {
            var portError error
            if ret.Host, ret.Port, portError = net.SplitHostPort(u.Host); portError != nil {
                log.Debug().Err(portError)
                ret.Host = u.Host
            }
        }
    }
    return
}

