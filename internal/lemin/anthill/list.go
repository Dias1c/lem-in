package anthill

// pushing into front of list
func (l *list) PushFront(r *room) {
	newNode := &node{Room: r}
	if l.Front == nil {
		l.Len = 1
		l.Front = newNode
		l.Back = newNode
		return
	}
	l.Len++
	newNode.Next = l.Front
	l.Front = newNode
}

// pushing into back of list
func (l *list) PushBack(r *room) {
	newNode := &node{Room: r}
	if l.Front == nil {
		l.Len = 1
		l.Front = newNode
		l.Back = newNode
		return
	}
	l.Len++
	l.Back.Next = newNode
	l.Back = newNode
}

// removes front node of list
func (l *list) RemoveFront() {
	if l.Front == nil {
		return
	}
	l.Len--
	l.Front = l.Front.Next
	if l.Front == nil {
		l.Back = nil
	}
}

// returns array of rooms
func (l *list) ToArray(lenArr int) []*room {
	if l.Front == nil || lenArr < 1 {
		return nil
	}
	res := make([]*room, lenArr)
	cur := l.Front
	for i := range res {
		res[i] = cur.Room
		cur = cur.Next
		if cur == nil {
			break
		}
	}
	return res
}
