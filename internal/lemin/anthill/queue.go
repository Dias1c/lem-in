package anthill

func (q *antQueue) Enqueue(num, path, pos int) {
	ant := &antStruct{
		Num:  num,
		Path: path,
		Pos:  pos,
	}
	if q.Back == nil {
		q.Front = ant
		q.Back = ant
		return
	}
	q.Back.Next = ant
	q.Back = ant
}

func (q *antQueue) Dequeue() *antStruct {
	if q.Front == nil {
		return nil
	}
	res := q.Front
	if q.Front == q.Back {
		q.Front = nil
		q.Back = nil
	} else {
		q.Front = q.Front.Next
	}
	return res
}

func (q *antQueue) EnqueueAnt(ant *antStruct) {
	ant.Next = nil
	if q.Back == nil {
		q.Front = ant
		q.Back = ant
		return
	}
	q.Back.Next = ant
	q.Back = ant
}

func (q *sortedQueue) Enqueue(r *room, weight int, mark bool) {
	node := &weightNode{
		Room:   r,
		Weight: weight,
		Mark:   mark,
	}
	if q.Front == nil {
		q.Front = node
		q.Back = node
		return
	}
	if q.Front.Weight > weight {
		node.Next = q.Front
		q.Front = node
		return
	} else if q.Back.Weight <= weight {
		q.Back.Next = node
		q.Back = node
		return
	}
	prev := q.Front
	cur := prev.Next
	for cur != nil {
		if cur.Weight > weight {
			prev.Next = node
			node.Next = cur
			return
		}
		prev = cur
		cur = cur.Next
	}
}

func (q *sortedQueue) Dequeue() *weightNode {
	if q.Front == nil {
		return nil
	}
	res := q.Front
	if q.Front == q.Back {
		q.Front = nil
		q.Back = nil
	} else {
		q.Front = q.Front.Next
	}
	return res
}
