// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 Google Inc.  All rights reserved.
// http://code.google.com/p/goprotobuf/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package proto_test

import (
	"bytes"
	"testing"

	"goprotobuf.googlecode.com/hg/proto"

	pb "./testdata/_obj/test_proto"
)

func newTestMessage() *pb.MyMessage {
	msg := &pb.MyMessage{
		Count: proto.Int32(42),
		Name:  proto.String("Dave"),
		Quote: proto.String(`"I didn't want to go."`),
		Pet:   []string{"bunny", "kitty", "horsey"},
		Inner: &pb.InnerMessage{
			Host:      proto.String("footrest.syd"),
			Port:      proto.Int32(7001),
			Connected: proto.Bool(true),
		},
		Others: []*pb.OtherMessage{
			&pb.OtherMessage{
				Key:   proto.Int64(0xdeadbeef),
				Value: []byte{1, 65, 7, 12},
			},
			&pb.OtherMessage{
				Weight: proto.Float32(6.022),
				Inner: &pb.InnerMessage{
					Host: proto.String("lesha.mtv"),
					Port: proto.Int32(8002),
				},
			},
		},
		Bikeshed: pb.NewMyMessage_Color(pb.MyMessage_BLUE),
		Somegroup: &pb.MyMessage_SomeGroup{
			GroupField: proto.Int32(8),
		},
	}
	ext := &pb.Ext{
		Data: proto.String("Big gobs for big rats"),
	}
	if err := proto.SetExtension(msg, pb.E_Ext_More, ext); err != nil {
		panic(err)
	}
	return msg
}

const text = `count: 42
name: "Dave"
quote: "\"I didn't want to go.\""
pet: "bunny"
pet: "kitty"
pet: "horsey"
inner: <
  host: "footrest.syd"
  port: 7001
  connected: true
>
others: <
  key: 3735928559
  value: "\x01A\a\f"
>
others: <
  weight: 6.022
  inner: <
    host: "lesha.mtv"
    port: 8002
  >
>
bikeshed: BLUE
SomeGroup {
  group_field: 8
}
[test_proto.more]: <
  data: "Big gobs for big rats"
>
`

func TestMarshalTextFull(t *testing.T) {
	buf := new(bytes.Buffer)
	proto.MarshalText(buf, newTestMessage())
	s := buf.String()
	if s != text {
		t.Errorf("Got:\n===\n%v===\nExpected:\n===\n%v===\n", s, text)
	}
}

func compact(src string) string {
	// s/[ \n]+/ /g; s/ $//;
	dst := make([]byte, len(src))
	space := false
	j := 0
	for i := 0; i < len(src); i++ {
		c := src[i]
		if c == ' ' || c == '\n' {
			space = true
			continue
		}
		if j > 0 && (dst[j-1] == ':' || dst[j-1] == '<' || dst[j-1] == '{') {
			space = false
		}
		if c == '{' {
			space = false
		}
		if space {
			dst[j] = ' '
			j++
			space = false
		}
		dst[j] = c
		j++
	}
	if space {
		dst[j] = ' '
		j++
	}
	return string(dst[0:j])
}

var compactText = compact(text)

func TestCompactText(t *testing.T) {
	s := proto.CompactTextString(newTestMessage())
	if s != compactText {
		t.Errorf("Got:\n===\n%v===\nExpected:\n===\n%v===\n", s, compactText)
	}
}
