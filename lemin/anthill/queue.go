package anthill

import (
	"fmt"
	"log"
)

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
	if q.Front == q.Back {
		log.Fatal("Enqueue #1 front == back")
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

func (q *sortedQueue) SortEnqueue(r *room, weight int, mark bool) {
	if q.Front == nil {
		q.Enqueue(r, weight, mark)
		return
	} else if q.Front.Room == r {
		if mark {
			if q.Front.Mark && q.Front.Weight <= weight {
				return
			}
			if q.Front.Weight > weight {
				q.Front = q.Front.Next
				if q.Front == nil {
					q.Back = nil
				}
			}
		} else {
			if q.Front.Mark || q.Front.Weight <= weight {
				if q.Front.Mark && q.Front.Weight > weight {
					q.Enqueue(r, weight, mark)
				}
				return
			}
			q.Front = q.Front.Next
			if q.Front == nil {
				q.Back = nil
			}
		}
		q.Enqueue(r, weight, mark)
		return
	}
	prev := q.Front
	cur := prev.Next
	for cur != nil {
		if cur.Room == r {
			if mark {
				if cur.Mark && cur.Weight <= weight {
					return
				}
				if cur.Weight > weight {
					prev.Next = cur.Next
					if prev.Next == nil {
						q.Back = prev
					}
				}
			} else {
				if cur.Mark || cur.Weight <= weight {
					if cur.Mark && cur.Weight > weight {
						q.Enqueue(r, weight, mark)
					}
					return
				}
				prev.Next = cur.Next
				if prev.Next == nil {
					q.Back = prev
				}
			}
			q.Enqueue(r, weight, mark)
			return
		}
		prev = cur
		cur = cur.Next
	}
	q.Enqueue(r, weight, mark)
}

func (q *sortedQueue) DebugPrint() {
	if q.Front == nil {
		fmt.Println("queue nil")
		return
	}
	cur := q.Front
	for cur != q.Back {
		fmt.Printf("%s|%d --> ", cur.Room.Name, cur.Room.Weight)
		cur = cur.Next
	}
	fmt.Printf("%s|%d\n", cur.Room.Name, cur.Room.Weight)
}

func (q *sortedQueue) DebugLen() int {
	if q.Front == nil {
		return 0
	}
	res := 1
	cur := q.Front
	for cur != q.Back {
		cur = cur.Next
		res++
	}
	return res
}
