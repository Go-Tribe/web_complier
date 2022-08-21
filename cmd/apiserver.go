package main

import (
	"math/rand"
	"time"
	"web_complier/internal/appserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	appserver.NewApp("api-server")
}
