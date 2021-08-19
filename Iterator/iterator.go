package Iterator

type Object struct {
	val interface{}
}

type Iterator interface {
	HasNext() bool
	Next() *Object
}

//type Container interface {
//	getIterator() Iterator
//}

type NumArray struct {
	nums []int
	cur  int
}

func NewNumArray(arr []int) *NumArray {
	return &NumArray{arr, 0}
}

func (a NumArray) HasNext() bool {
	return a.cur < len(a.nums)
}

func (a *NumArray) Next() *Object {
	defer func() { a.cur++ }()
	if a.HasNext() {
		return &Object{a.nums[a.cur]}
	}
	return nil
}

//func (a *NumArray) getInterator() Iterator {
//
//}
