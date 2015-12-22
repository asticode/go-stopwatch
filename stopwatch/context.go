// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import "golang.org/x/net/context"

type contextKey int

const (
	stopwatchContextKey contextKey = 0
)

// NewContext adds a stopwatch to the context
func NewContext(ctx context.Context, s Stopwatch) context.Context {
	return context.WithValue(ctx, stopwatchContextKey, s)
}

// FromContext retrieve a stopwatch from the context
func FromContext(ctx context.Context) (Stopwatch, bool) {
	l, ok := ctx.Value(stopwatchContextKey).(Stopwatch)
	return l, ok
}
