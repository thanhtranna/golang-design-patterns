package main

import "fmt"

type Command interface {
	Execute()
}

type SomeCommand struct {
	message string
}

func (s *SomeCommand) Execute() {
	fmt.Println(s.message)
}

type SomeSpecialCommand struct {
	message string
}

func (s *SomeSpecialCommand) Execute() {
	fmt.Println("@" + s.message)
}

type CommandInvoker struct {
	queue          []Command
	processedItems int
}

func (c *CommandInvoker) processQueue() {
	for i := range c.queue {
		c.queue[i].Execute()
		c.processedItems++
	}
}

func (c *CommandInvoker) ProcessedItems() int {
	return c.processedItems
}

func (c *CommandInvoker) addToQueue(i Command) {
	fmt.Println("Appending command")
	c.queue = append(c.queue, i)
	if len(c.queue) == 3 {
		c.processQueue()
		c.queue = []Command{}
	}
}
