package main

type Stack struct {
	items []rune
}

// Push добавляет элемент в стек
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop удаляет и возвращает элемент из стека
func (s *Stack) Pop() rune {
	if len(s.items) == 0 {
		return 0 // 0 используется для обозначения пустого символа
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek возвращает элемент с вершины стека
func (s *Stack) Peek() rune {
	if len(s.items) == 0 {
		return 0 // 0 используется для обозначения пустого символа
	}
	return s.items[len(s.items)-1]
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

type StackS struct {
	item []string
}

func (s *StackS) PushS(item string) {
	s.item = append(s.item, item)
}

func (s *StackS) PopS() (string, bool) {
	if len(s.item) == 0 {
		return "", false // Error: Empty stack
	}
	itemS := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return itemS, true
}

func (s *StackS) PeekS() string {
	if len(s.item) == 0 {
		return "" // Error: Empty stack
	}
	return s.item[len(s.item)-1]
}

func (s *StackS) IsEmptyS() bool {
	return len(s.item) == 0
}
