package metrics

import "sync"

type Namespace struct {
	sync.Mutex

	name   string
	labels map[string]string
}

func (n *Namespace) Init(name string) {
	n.Mutex = sync.Mutex{}
	n.name = name
	n.labels = make(map[string]string, 0)
}

func (n *Namespace) Name() string {
	return n.name
}

func (n *Namespace) Label(label string) string {
	l, ok := n.labels[label]
	if !ok {
		n.setLabel(label)
		l = label
	}
	return l
}

func (n *Namespace) setLabel(labels ...string) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	for _, l := range labels {
		n.labels[l] = l
	}
}
