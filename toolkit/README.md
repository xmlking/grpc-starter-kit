# toolkit

## TODO
 
WithConfig( interface{})

```go
package main

import (
	"context"
	"log"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	moron "github.com/spencer-p/moroncloudevents"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from HTTP!"))
}

func receive(ctx context.Context, event cloudevents.Event, r *cloudevents.EventResponse) error {
	databytes, _ := event.DataBytes()
	log.Printf("Received CloudEvent with data: %q\n", databytes)
	return nil
}

func main() {
    ctx := signals.NewContext()

	svr, err := moron.NewServer(&moron.ServerConfig{
		Port:                  "8080",
		CloudEventReceivePath: "/apis/receive",
	})
	if err != nil {
		log.Fatal("Could not create server: ", err)
	}

	svr.HandleCloudEvents(receive)

	svr.HandleFunc("/", index)

	log.Fatal(svr.ListenAndServe())
}
```

## 🔗 Credits
https://github.com/infobloxopen/atlas-app-toolkit/tree/master/server
https://github.com/spencer-p/moroncloudevents/tree/master
