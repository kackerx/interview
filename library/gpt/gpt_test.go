package gpt

import (
	"context"
	"fmt"
	"testing"
)

func TestGpt(t *testing.T) {
	trans, err := Trans(context.Background(), "apple")
	if err != nil {
		panic(err)
	}

	fmt.Println(trans)
}
