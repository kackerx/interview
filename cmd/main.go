package main

import (
	"context"

	"github.com/kackerx/interview/common/log"
)

func main() {
	log.New(context.Background()).Debug("hehe", "k", "v")
}
