package config

import (
    "os"
)

type Config struct {
    GrpcPort string
}

func NewConfig() Config {
    port, exists := os.LookupEnv("BITMEX_PROXY_GRPC_PORT")
    c := Config{}

    c.GrpcPort = ":" + "8080"
    if exists {
        c.GrpcPort = ":" + port
    }

    return c
}
