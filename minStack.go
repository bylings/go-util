package util

// MinStack 最小栈结构
type MinStack struct {
	data    []int
	minData []int // 保存最小元素
	len     int
}

// NewMinStack 返回最小栈
func NewMinStack() *MinStack {
	return new(MinStack)
}

// GetMin 获取最小元素
func (s *MinStack) GetMin() int {
	if s.len == 0 {
		return 0
	}
	return s.minData[len(s.minData)-1]
}

// Push 入栈
func (s *MinStack) Push(val int) {
	if s.len == 0 || val <= s.minData[len(s.minData)-1] {
		s.minData = append(s.minData, val)
	}
	s.data = append(s.data, val)
	s.len++
}

// Pop 出栈（后进先出LIFO）
func (s *MinStack) Pop() int {
	if s.len == 0 {
		return 0
	}

	val := s.data[s.len-1]

	if val == s.minData[len(s.minData)-1] {
		s.minData = s.minData[:len(s.minData)-1]
	}
	s.len--
	s.data = s.data[:len(s.data)-1] // 后入先出
	return val
}
