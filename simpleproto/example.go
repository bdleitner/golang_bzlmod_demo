package example

import (
  "bdl.com/demos/simpleproto/expb"
)

func value(p *expb.Example) string {
  return p.GetValue()
}
