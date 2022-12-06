package utils

// see https://www.folkstalk.com/2022/09/golang-stack-with-code-examples.html

type Stack[T any] struct {
	list []T
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

func NewStackFilled[T any](values []T) Stack[T] {
	return Stack[T]{list: values}
}

func (stack *Stack[T]) Peek() T {
	if stack.Empty() {
		panic("Stack is empty")
	}
	return stack.list[len(stack.list)-1]
}

func (stack *Stack[T]) Push(elm T) {
	stack.list = append(stack.list, elm)
}

func (stack *Stack[T]) Pop() (elm T) {
	if stack.Empty() {
		panic("Stack is empty")
	}

	elm, stack.list = stack.Peek(), stack.list[0:len(stack.list)-1]
	return elm
}

func (stack *Stack[T]) Empty() bool {
	return len(stack.list) == 0
}

func (stack *Stack[T]) Length() int {
	return len(stack.list)
}

func (stack *Stack[T]) Clone() Stack[T] {
	return NewStackFilled(ShallowCopy(stack.list))
}
