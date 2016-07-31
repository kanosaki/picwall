package core

import (
	"fmt"
	"time"
	"sync"
)

type Space interface {
	Get(key string) (Source, bool)
}

type Center struct {
	nodes    map[string]Source
	sessions map[string]*Session
}

func (c *Center) NewSession(name string) (*Session, error) {
	if _, ok := c.sessions[name]; ok {
		return nil, fmt.Errorf("Session %s already exists! (duplicate named session is not allowed!)", name)
	}
	s := &Session{
		parent: c,
	}
	s.parent = c
	c.sessions[name] = s
	return s, nil
}

func (c *Center) Get(key string) (Source, bool) {
	return c.nodes[key]
}

func (c *Center) SetGlobalNode(key string, v Source) {
	if v == c {
		panic("Cannot set Center as its own node")
	}
	c.nodes[key] = v
}

func (c *Center) Shutdown(timeout time.Duration) {
	wg := sync.WaitGroup{}
	wg.Add(len(c.nodes) + len(c.sessions))
	for _, session := range c.sessions {
		go func() {
			session.Shutdown()
			wg.Done()
		}()
	}
	for _, node := range c.nodes {
		go func() {
			node.Shutdown()
			wg.Done()
		}()
	}
	c := make(chan struct{})
	go func() {
		wg.Wait()
		c <- struct{}{}
	}()
	select {
	case <-c:
		fmt.Println("Successfully shutdown")
	case <-time.After(timeout):
		fmt.Println("Shutdown timeout")
	}
}

