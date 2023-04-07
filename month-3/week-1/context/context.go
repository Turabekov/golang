package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// root := context.Background()

	// root, _ = context.WithDeadline(root, time.Now().Add(time.Second*5))

	// now := time.Now()

	// for {
	// 	if root.Err() == context.DeadlineExceeded {
	// 		break
	// 	}
	// 	fmt.Println(time.Now())
	// 	fmt.Println("Time:", now)
	// }

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()

	ctx = context.WithValue(ctx, "x", "value")
	fn(ctx)
	fn(ctx)
	fn(ctx)
	fn(ctx)

	cancel()

	fn(ctx)
	fn(ctx)
	fn(ctx)
	fn(ctx)
}

func fn(ctx context.Context) {

	if ctx.Err() != context.Canceled {
		fmt.Println(ctx.Value("x"))
	}
}
