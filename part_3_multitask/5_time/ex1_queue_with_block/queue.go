package ex1queuewithblock

import ("errors")


var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

type Queue struct {
	elems chan int
	n int
}

func (q Queue) Get(block bool) (int, error) {
	select{
	case v:= <- q.elems:
		return v, nil
	default:
		if block {
			return <-q.elems, nil
		} else {			
			return 0, ErrEmpty
		}
	}
}

func (q Queue) Put(val int, block bool) error {
	select {
	case q.elems<- val:
		return nil
	default:
		if block {
			q.elems <-val
			return nil
		} else {
			return ErrFull
		}
	}
}

func MakeQueue(n int) Queue {
	elems := make(chan int, n)
	q := Queue{elems, n}
	return q
}