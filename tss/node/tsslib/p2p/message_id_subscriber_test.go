package p2p

import (
	. "gopkg.in/check.v1"
)

type MessageIDSubscriberTestSuit struct{}

var _ = Suite(&MessageIDSubscriberTestSuit{})

func (MessageIDSubscriberTestSuit) TestMessageIDSubscriber(c *C) {
	ms := NewMessageIDSubscriber()
	ms.Subscribe("hello", make(chan *Message))
	channel := ms.GetSubscriber("helloworld")
	c.Assert(channel, IsNil)
	channel1 := ms.GetSubscriber("hello")
	c.Assert(channel1, NotNil)
	c.Assert(ms.IsEmpty(), Equals, false)
	ms.UnSubscribe("hello")
	channel2 := ms.GetSubscriber("hello")
	c.Assert(channel2, IsNil)
	c.Assert(ms.IsEmpty(), Equals, true)
}
