// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/tilt-dev/tilt-apiserver/blob/main/pkg/storage/filepath/watchset.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Kubernetes Authors.

package file

import (
	"fmt"
	"sync"

	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/storage"
)

// Keeps track of which watches need to be notified
type WatchSet struct {
	mu      sync.RWMutex
	nodes   map[int]*watchNode
	counter int
}

func NewWatchSet() *WatchSet {
	return &WatchSet{
		nodes:   make(map[int]*watchNode, 20),
		counter: 0,
	}
}

// Creates a new watch with a unique id, but
// does not start sending events to it until start() is called.
func (s *WatchSet) newWatch(requestedRV uint64) *watchNode {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++

	return &watchNode{
		requestedRV: requestedRV,
		id:          s.counter,
		s:           s,
		updateCh:    make(chan watch.Event, 10),
		outCh:       make(chan watch.Event),
	}
}

func (s *WatchSet) cleanupWatchers() {
	fmt.Println("Pre cleanup - lock get")
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Println("Looping on nodes for cleanup")
	for _, w := range s.nodes {
		fmt.Println("Stopping node")
		w.stop()
	}
}

func (s *WatchSet) notifyWatchers(ev watch.Event) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Println("notifyWatchers START")

	if len(s.nodes) == 0 {
		fmt.Println("No nodes")
		return
	}

	fmt.Println("NOTIFY: length of nodes=", len(s.nodes))

	fmt.Println("Looping on nodes for notify")
	for _, w := range s.nodes {
		fmt.Println("Updating channel with ev")
		w.updateCh <- ev
	}

	fmt.Println("notifyWatchers COMPLETE")
}

type watchNode struct {
	s           *WatchSet
	id          int
	updateCh    chan watch.Event
	outCh       chan watch.Event
	requestedRV uint64
}

// Start sending events to this watch.
func (w *watchNode) Start(p storage.SelectionPredicate, initEvents []watch.Event) {
	w.s.mu.Lock()
	w.s.nodes[w.id] = w
	w.s.mu.Unlock()

	fmt.Println("Start pre")

	go func() {
		for _, e := range initEvents {
			w.outCh <- e
		}

		for e := range w.updateCh {
			fmt.Println("From update channel", e)
			ok, err := p.Matches(e.Object)
			if err != nil {
				continue
			}

			if !ok {
				continue
			}
			fmt.Println("To out channel", e)
			w.outCh <- e
		}

		fmt.Println("Start post")
		close(w.outCh)
	}()
}

func (w *watchNode) Stop() {
	w.s.mu.RLock()
	defer w.s.mu.RUnlock()
	w.stop()
}

func (w *watchNode) stop() {
	delete(w.s.nodes, w.id)
	fmt.Println("Stop post")
	close(w.updateCh)
}

func (w *watchNode) ResultChan() <-chan watch.Event {
	fmt.Println("ResultChan")
	return w.outCh
}
