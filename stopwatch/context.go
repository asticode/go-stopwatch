// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import (
	"errors"

	"golang.org/x/net/context"
)

type contextKey int

const (
	stopwatchContextKey contextKey = 0
)

var (
	errContextInvalid = errors.New("Invalid context")
)

// NewContext adds a stopwatch to the context
func NewContext(ctx context.Context, s Stopwatch) context.Context {
	return context.WithValue(ctx, stopwatchContextKey, s)
}

// FromContext retrieve a stopwatch from the context
func FromContext(ctx context.Context) (Stopwatch, error) {
	if l, ok := ctx.Value(stopwatchContextKey).(Stopwatch); ok {
		return l, nil
	}
	return nil, errContextInvalid
}
