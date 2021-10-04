package anthill

type path struct {
	Len         int
	Front, Back *pNode
}

type pNode struct {
	Prev, Next *pNode
	Room       *room
}

func (p *path) PushFront(r *room) {
	newNode := &pNode{Room: r}
	p.Len++
	if p.Front == nil {
		p.Front = newNode
		p.Back = newNode
		return
	}
	newNode.Next = p.Front
	p.Front.Prev = newNode
	p.Front = newNode
}

func (p *path) PushBack(r *room) {
	newNode := &pNode{Room: r}
	p.Len++
	if p.Front == nil {
		p.Front = newNode
		p.Back = newNode
		return
	}
	p.Back.Next = newNode
	newNode.Prev = p.Back
	p.Back = newNode
}

func (p *path) RemoveFront() {
	if p.Front == nil {
		return
	}
	p.Front = p.Front.Next
	p.Front.Prev = nil
	if p.Front == nil {
		p.Back = nil
	}
}

func (p *path) RefreshData() {
	if p.Front == nil {
		p.Len = 0
		p.Back = nil
		return
	} else if p.Front.Next == nil {
		p.Len = 1
		p.Back = p.Front
		return
	}
	newLen := 1
	prev := p.Front
	cur := p.Front.Next
	for cur != nil {
		newLen++
		if cur.Prev != prev {
			cur.Prev = prev
		}
		prev = cur
		cur = cur.Next
	}
	p.Len = newLen
	p.Back = prev
}

func (p *path) GetCopyWithOrigRoom() *path {
	result := &path{}
	for n := p.Front; n != nil; n = n.Next {
		result.PushBack(n.Room)
	}
	return result
}
