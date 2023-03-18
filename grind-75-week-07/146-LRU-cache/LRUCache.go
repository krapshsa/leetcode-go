package LRU_cache

type Slot struct {
	key   int
	value int
	prev  *Slot
	next  *Slot
}

type LRUCache struct {
	capacity    int
	length      int
	dummyRecent *Slot
	dummyLeast  *Slot
	SlotMap     map[int]*Slot
}

func Constructor(capacity int) LRUCache {
	dummyRecent := &Slot{-1, -1, nil, nil}
	dummyLeast := &Slot{-1, -1, nil, nil}
	dummyRecent.next = dummyLeast
	dummyLeast.prev = dummyRecent

	return LRUCache{
		capacity,
		0,
		dummyRecent,
		dummyLeast,
		make(map[int]*Slot),
	}
}

func (this *LRUCache) Get(key int) int {
	if slot, ok := this.SlotMap[key]; ok && slot != nil {
		if this.length > 1 {
			this.removeOneSlot(slot)
			this.putSlotToRecent(slot)
		}
		return slot.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	var s *Slot

	if slot, ok := this.SlotMap[key]; ok && slot != nil {
		slot.key = key
		slot.value = value
		s = slot
		this.removeOneSlot(s)
	} else {
		if this.length == this.capacity {
			slot := this.dummyLeast.prev
			this.removeOneSlot(slot)
			delete(this.SlotMap, slot.key)
			this.length--
		}

		// add new
		s = &Slot{key, value, nil, nil}
		this.SlotMap[key] = s
		this.length++
	}

	this.putSlotToRecent(s)
}

func (this *LRUCache) removeOneSlot(slot *Slot) {
	prev := slot.prev
	next := slot.next

	prev.next = next
	next.prev = prev

	slot.next = nil
	slot.prev = nil
}

func (this *LRUCache) putSlotToRecent(slot *Slot) {
	prevRecent := this.dummyRecent.next

	prevRecent.prev = slot
	this.dummyRecent.next = slot

	slot.prev = this.dummyRecent
	slot.next = prevRecent
}
