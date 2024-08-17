package Env

import "os"

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (e *Env) GetEnv(key string) string {
	return os.Getenv(key)
}
