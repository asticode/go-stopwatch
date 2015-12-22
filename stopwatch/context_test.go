// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestNewContext(t *testing.T) {
	// Initialize
	s := &stopwatch{}
	ctx := context.Background()

	// Create context
	ctx = NewContext(ctx, s)

	// Assert
	assert.Equal(t, nil, ctx.Value(0))

	// Get stopwatch
	sfc, ok := FromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, s, sfc)
}
