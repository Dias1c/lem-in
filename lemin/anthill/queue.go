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

func (q *sortedQueue) Enqueue(r *room) {
	n := &node{
		Room: r,
	}
	if q.Back == nil {
		q.Front = n
		q.Back = n
		return
	}
	if q.Front.Room.Weight >= n.Room.Weight {
		n.Next = q.Front
		q.Front = n
		return
	} else if q.Back.Room.Weight < n.Room.Weight {
		q.Back.Next = n
		q.Back = n
		return
	}
	prev, cur := q.Front, q.Front.Next
	for cur != nil {
		if cur.Room.Weight >= n.Room.Weight {
			prev.Next = n
			n.Next = cur
			return
		}
		cur = cur.Next
	}
}

func (q *sortedQueue) Dequeue() *node {
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

func (q *sortedQueue) SortEnqueue(r *room) {
	if q.Front.Room == r {
		q.Front = q.Front.Next
		if q.Front == nil {
			q.Back = nil
		}
		q.Enqueue(r)
		return
	}
	prev, cur := q.Front, q.Front.Next
	for cur != nil {
		if cur.Room == r {
			prev.Next = cur.Next
			if cur == q.Back {
				q.Back = prev
			}
			q.Enqueue(r)
			return
		}
		prev, cur = cur, cur.Next
	}
	q.Enqueue(r)
}
