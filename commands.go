package main

import (
	"errors"
)

type handlerFunc func(*state, Command) error

type Commands struct {
	handlers map[string]handlerFunc
}

type Command struct {
	Name string
	Args []string
}

func (c *Commands) run(s *state, cmd Command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return errors.New("command not found: " + cmd.Name)
	}

	return handler(s, cmd)
}

func (c *Commands) register(name string, f func(s *state, cmd Command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]handlerFunc)
	}

	c.handlers[name] = f
}
