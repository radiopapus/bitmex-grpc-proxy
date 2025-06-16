package main

import (
    "bitmex-grpc-proxy/config"
    "bitmex-grpc-proxy/internal/server"
)

func main() {
    c := config.NewConfig()
    server.Run(c)
}
