package datastructure

type Stack struct {
	Nodes []interface{}
}

func (s *Stack) Push(node interface{}) {
	s.Nodes = append(s.Nodes, node)
}

func (s *Stack) Pop() {
	size := s.Size()
	if size > 0 {
		s.Nodes = s.Nodes[:size-1]
	}
}

func (s *Stack) Size() int {
	return len(s.Nodes)
}

func (s *Stack) Last() interface{} {
	return s.Nodes[s.Size()-1]
}
