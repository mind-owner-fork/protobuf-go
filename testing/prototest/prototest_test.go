// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototest_test

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/internal/flags"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/prototest"

	irregularpb "google.golang.org/protobuf/internal/testprotos/irregular"
	testpb "google.golang.org/protobuf/internal/testprotos/test"
	_ "google.golang.org/protobuf/internal/testprotos/test/weak1"
	_ "google.golang.org/protobuf/internal/testprotos/test/weak2"
	test3pb "google.golang.org/protobuf/internal/testprotos/test3"
)

func Test(t *testing.T) {
	ms := []proto.Message{
		(*testpb.TestAllTypes)(nil),
		(*test3pb.TestAllTypes)(nil),
		(*testpb.TestRequired)(nil),
		(*irregularpb.Message)(nil),
		(*testpb.TestAllExtensions)(nil),
	}
	if flags.Proto1Legacy {
		ms = append(ms, (*testpb.TestWeak)(nil))
	}

	for _, m := range ms {
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			prototest.TestMessage(t, m, prototest.MessageOptions{})
		})
	}
}
