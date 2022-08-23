package common

type WordDictionary struct {
	Childs map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		Childs: make(map[byte]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		s := Constructor()
		this.Childs[' '] = &s
		return
	}

	v := word[0]
	if child, ok := this.Childs[v]; ok {
		child.AddWord(word[1:])
	} else {
		s := Constructor()
		s.AddWord(word[1:])
		this.Childs[v] = &s
	}
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		_, ok := this.Childs[' ']
		return ok
	}

	v := word[0]
	if v == '.' {
		for _, child := range this.Childs {
			if child.Search(word[1:]) {
				return true
			}
		}
	} else if child, ok := this.Childs[v]; ok {
		return child.Search(word[1:])
	} else {
		return false
	}

	return false
}
