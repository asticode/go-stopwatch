// Copyright 2015, Quentin RENARD. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stopwatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertNanosecondsToMilliseconds(t *testing.T) {
	// Initialize
	i := 12345678
	e := float64(12.345678)

	// Assert
	assert.Equal(t, e, convertNanosecondsToMilliseconds(i))
}

func TestConvertBytesToMegabytes(t *testing.T) {
	// Initialize
	i := uint64(12345678)
	e := float64(i) / float64(1024*1024)

	// Assert
	assert.Equal(t, e, convertBytesToMegabytes(i))
}
