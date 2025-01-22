package exercises

import "fmt"

// Define the Clock type here.
type Clock struct {
	Hour   int
	Minute int
}

func (c Clock) validate() Clock {
	if c.Minute < 0 {
		multiplex := ((-c.Minute) / 60) + 1
		c.Minute += multiplex * 60
		c.Hour -= multiplex
	}

	if c.Minute >= 60 {
		c.Hour += c.Minute / 60
		c.Minute %= 60
	}

	if c.Hour >= 24 {
		c.Hour %= 24
	}

	if c.Hour < 0 {
		multiplex := ((-c.Hour) / 24) + 1
		c.Hour += multiplex * 24
	}

	c.Hour %= 24

	return c
}

func New(h, m int) Clock {
	clock := Clock{Hour: h, Minute: m}
	return clock.validate()
}

func (c Clock) Add(m int) Clock {
	c.Minute += m
	return c.validate()
}

func (c Clock) Subtract(m int) Clock {
	c.Minute -= m
	return c.validate()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
