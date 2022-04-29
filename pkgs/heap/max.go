package heap

type IMaxHeap interface {
	insert(data data)
	deleteMax()
	findMax() *data
}

// root - i
// left child - 2*i + 1
// right child - 2*i + 2

type maxHeap struct {
	dataArray []data
	size      int
}

func (he *maxHeap) insert(data data) {
	he.dataArray = append(he.dataArray, data)
	currentPosition := he.size
	he.size++

	for {
		parent := (currentPosition - 1) / 2

		if he.dataArray[parent].key < he.dataArray[currentPosition].key {
			he.dataArray[currentPosition] = he.dataArray[parent]
			he.dataArray[parent] = data
			currentPosition = parent
		} else {
			break
		}
	}
}

func (he *maxHeap) deleteMax() {
	if he.size == 0 {
		return
	}

	he.dataArray[0] = he.dataArray[he.size-1]

	parent := 0

	var leftChild, rightChild data

	for {
		leftChildIndex := 2*parent + 1
		rightChildIndex := 2*parent + 2

		if rightChildIndex <= he.size-1 {
			leftChild = he.dataArray[leftChildIndex]
			rightChild = he.dataArray[rightChildIndex]
		} else if leftChildIndex <= he.size-1 {
			leftChild = he.dataArray[leftChildIndex]
		} else {
			break
		}

		if leftChild.key >= rightChild.key {
			if leftChild.key > he.dataArray[parent].key {
				temp := he.dataArray[parent]
				he.dataArray[parent] = leftChild
				he.dataArray[leftChildIndex] = temp

				parent = leftChildIndex
			} else {
				break
			}
		}

		if rightChild.key > leftChild.key {
			if rightChild.key > he.dataArray[parent].key {
				temp := he.dataArray[parent]
				he.dataArray[parent] = rightChild
				he.dataArray[rightChildIndex] = temp

				parent = rightChildIndex
			} else {
				break
			}
		}
	}
}

func (he *maxHeap) findMax() *data {
	return &he.dataArray[0]
}

func NewMaxHeap() *maxHeap {
	return &maxHeap{}
}
