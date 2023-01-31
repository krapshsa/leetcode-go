package implement_trie

type Trie struct {
	child map[byte]*Trie
	end   bool
}

func Constructor() Trie {
	return Trie{make(map[byte]*Trie), false}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		b := word[i]
		_, ok := (*cur).child[b]
		if false == ok {
			(*cur).child[b] = &Trie{make(map[byte]*Trie), false}
		}
		cur = (*cur).child[b]
	}
	(*cur).end = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		b := word[i]
		_, ok := (*cur).child[b]
		if false == ok {
			return false
		}
		cur = (*cur).child[b]
	}

	return true == (*cur).end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		b := prefix[i]
		_, ok := (*cur).child[b]
		if false == ok {
			return false
		}
		cur = (*cur).child[b]
	}

	return true
}
