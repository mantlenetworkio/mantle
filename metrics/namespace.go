package metrics

type Namespace struct {
	name string

	labels map[string]string
}

func (n *Namespace) Init(name string) {
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
	for _, l := range labels {
		n.labels[l] = l
	}
}
