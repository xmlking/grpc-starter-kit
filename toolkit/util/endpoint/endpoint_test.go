package endpoint

import (
    "testing"

)

func TestParseTargetString(t *testing.T) {
    for _, test := range []struct {
        targetStr string
        want      Target
    }{
        {targetStr: "", want: Target{Scheme: "", Host: "", Port: "", Path: ""}},
        {targetStr: "dns:///google.com:8080", want: Target{Scheme: "dns", Host: "google.com", Port: "8080", Path: ""}},
        {targetStr: "dns:///google.com", want: Target{Scheme: "dns", Host: "google.com", Port: "", Path: ""}},
        {targetStr: "dns:///google.com/?a=b", want: Target{Scheme: "dns", Host: "google.com", Port: "", Path: "/"}},
        {targetStr: "https://www.server.com:9999", want: Target{Scheme: "https", Host: "www.server.com", Port: "9999", Path: ""}},
        {targetStr: "/unix/socket/address", want: Target{Scheme: "", Host: "", Port: "", Path: "/unix/socket/address"}},
        {targetStr: "unix:///tmp/mysrv.sock", want: Target{Scheme: "unix", Host: "", Port: "", Path: "/tmp/mysrv.sock"}},
    } {
        got, err := ParseTarget(test.targetStr)
        if err != nil {
            t.Error(err)
        }
        if got != test.want {
            t.Errorf("ParseTarget(%q) = %+v, want %+v", test.targetStr, got, test.want)
        }
    }
}
