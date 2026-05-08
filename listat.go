package student

func ListAt(l *NodeL, pos int) *NodeL {
	i := 0
	current := l

	for current != nil {
		if i == pos {
			return current
		}
		current = current.Next
		i++
	}

	return nil
}
