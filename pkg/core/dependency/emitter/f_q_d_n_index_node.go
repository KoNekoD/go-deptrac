package emitter

type FQDNIndexNode struct {
	FQDN  bool
	nodes map[string]*FQDNIndexNode
}

func NewFQDNIndexNode() *FQDNIndexNode {
	return &FQDNIndexNode{
		FQDN:  false,
		nodes: make(map[string]*FQDNIndexNode),
	}
}

func (f *FQDNIndexNode) IsFQDN() bool {
	return f.FQDN
}

func (f *FQDNIndexNode) GetNestedNode(keys []string) *FQDNIndexNode {
	index := f
	for _, key := range keys {
		node, ok := index.nodes[key]
		if !ok {
			return nil
		}
		index = node
	}
	return index
}

func (f *FQDNIndexNode) SetNestedNode(keys []string) {
	if len(keys) == 0 {
		f.FQDN = true
		return
	}

	key := keys[0]
	node, ok := f.nodes[key]

	if !ok {
		node = &FQDNIndexNode{}
		f.nodes[key] = node
	}

	node.SetNestedNode(keys[1:])
}
