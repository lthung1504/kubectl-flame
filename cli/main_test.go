package main

import (
	"fmt"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
	"testing"
)

func TestTry(t *testing.T) {
	fmt.Printf("os.Args = %+v\n", os.Args)
	os.Args = os.Args[5:]
	fmt.Printf("os.Args = os.Args[5:] = %+v\n", os.Args)

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	execute(streams)
}
