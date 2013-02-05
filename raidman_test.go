package raidman

import (
	"testing"
)

func TestTCP(t *testing.T) {
	c, err := Dial("tcp", "localhost:5555")
	if err != nil {
		t.Fatal(err.Error())
	}
	var event = &Event{
		State:   "success",
		Host:    "raidman",
		Service: "tcp",
		Int:     42,
		Ttl:     1,
	}

	err = c.Send(event)
	if err != nil {
		t.Error(err.Error())
	}

	events, err := c.Query("host = \"raidman\"")
	if err != nil {
		t.Error(err.Error())
	}

	if len(events) < 1 {
		t.Error("Submitted event not found")
	}

	c.Close()
}

func TestUDP(t *testing.T) {
	c, err := Dial("udp", "localhost:5555")
	if err != nil {
		t.Fatal(err.Error())
	}
	var event = &Event{
		State:   "warning",
		Host:    "raidman",
		Service: "udp",
		Double:  4.2,
		Ttl:     10.7,
	}

	err = c.Send(event)
	if err != nil {
		t.Error(err.Error())
	}
	c.Close()
}

func BenchmarkTCP(b *testing.B) {
	c, err := Dial("tcp", "localhost:5555")

	var event = &Event{
		State:   "good",
		Host:    "raidman",
		Service: "benchmark",
	}

	if err == nil {
		for i := 0; i < b.N; i++ {
			c.Send(event)
		}
	}
	c.Close()
}

func BenchmarkUDP(b *testing.B) {
	c, err := Dial("udp", "localhost:5555")

	var event = &Event{
		State:   "good",
		Host:    "raidman",
		Service: "benchmark",
	}

	if err == nil {
		for i := 0; i < b.N; i++ {
			c.Send(event)
		}
	}
	c.Close()
}
