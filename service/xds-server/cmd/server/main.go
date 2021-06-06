//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"context"
	"flag"

	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
    "github.com/rs/zerolog"
    "github.com/xmlking/grpc-starter-kit/service/xds-server/internal/processor"
    "github.com/xmlking/grpc-starter-kit/service/xds-server/internal/server"
    "github.com/xmlking/grpc-starter-kit/service/xds-server/internal/watcher"
    "github.com/xmlking/grpc-starter-kit/service/xds-server/logger"
)

var (


	watchDirectoryFileName string
	port                   uint
	basePort               uint
	mode                   string

	nodeID string
)

func init() {

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 9002, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")

	// Define the directory to watch for Envoy configuration files
	flag.StringVar(&watchDirectoryFileName, "watchDirectoryFileName", "config/xds-config.yml", "full path to directory to watch for files")
}

func main() {
	flag.Parse()

	appCtx := context.Background()
    zerolog.DefaultLogger().WithFields("component","xds-server")
    logr := zerolog.Ctx(appCtx)


	// Create a cache
	cache := cache.NewSnapshotCache(false, cache.IDHash{},  logger.NewLogger(logr))

	// Create a processor
	proc := processor.NewProcessor(cache, nodeID)

	// Create initial snapshot from file
	proc.ProcessFile(watcher.NotifyMessage{
		Operation: watcher.Create,
		FilePath:  watchDirectoryFileName,
	})

	// Notify channel for file system events
	notifyCh := make(chan watcher.NotifyMessage)

	go func() {
		// Watch for file changes
		watcher.Watch(watchDirectoryFileName, notifyCh)
	}()

	go func() {
		// Run the xDS server
		ctx := context.Background()
		srv := serverv3.NewServer(ctx, cache, nil)
		server.RunServer(ctx, srv, port)
	}()

	for {
		select {
		case msg := <-notifyCh:
			proc.ProcessFile(msg)
		}
	}
}
