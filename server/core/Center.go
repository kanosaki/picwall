package core

import (
	"fmt"
	"sync"
	"time"
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
	src, ok := c.nodes[key]
	return src, ok
}

func (c *Center) SetGlobalNode(key string, v Source) {
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
	cancel := make(chan struct{})
	go func() {
		wg.Wait()
		cancel <- struct{}{}
	}()
	select {
	case <-cancel:
		fmt.Println("Successfully shutdown")
	case <-time.After(timeout):
		fmt.Println("Shutdown timeout")
	}
}
