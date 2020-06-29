package main

import (
    "os"
)

func init() {
    os.Setenv("GOOGLE_CLOUD_PROJECT_ID", "my-project-id")
    os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")
    // os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "~/my-json.json")
}
