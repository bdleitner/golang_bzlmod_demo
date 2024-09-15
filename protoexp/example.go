package example

import (
  "bdl.com/demos/protoexp/expb"
  "github.com/golang/protobuf/proto"
)

func value(p *expb.Example) string {
  return p.GetValue()
}

func NewExample() *expb.Example {
  p := &expb.Example{}
  _ = proto.Unmarshal([]byte(`value: "foo"`), p)
  return p
}
