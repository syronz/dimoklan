package util

import (
	"context"
	"fmt"
	"reflect"
)

type contextKey struct{}

var parentKey = contextKey{}

func PrintContext(ctx context.Context, depth int) {
	fmt.Printf("%sType: %s\n", indentation(depth), reflect.TypeOf(ctx))
	fmt.Printf("%sDone: %v\n", indentation(depth), ctx.Done())
	fmt.Printf("%sErr: %v\n", indentation(depth), ctx.Err())
	fmt.Printf("%sValue: %v\n", indentation(depth), ctx.Value("key1")) // Replace with the key you want to inspect

	if parent := ctx.Value(parentKey); parent != nil {
		fmt.Printf("%sParent: %v\n", indentation(depth), parent)
	}

	fmt.Println("--------------")

	// Recursive call for each child in the context tree
	for _, child := range listChildren(ctx) {
		PrintContext(child, depth+1)
	}
}

func indentation(depth int) string {
	// return fmt.Sprintf("%s", "  |  ")[:depth*4]
	return " | "
}


func listChildren(ctx context.Context) []context.Context {
	// Extract all child contexts from the context tree
	children := make([]context.Context, 0)
	select {
	case <-ctx.Done():
		// Done channel closed, no children
	case <-make(chan struct{}):
		// This case is just for illustration, it won't be executed
	}
	return children
}
