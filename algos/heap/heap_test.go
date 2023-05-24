package heap

import "testing"

func Test_should_create(t *testing.T) {
	NewHeap()
}

func Test_should_add(t *testing.T) {
	heap := NewHeap()
	heap.add(2)
}

func Test_should_return_length(t *testing.T) {
	heap := NewHeap()
	heap.add(3asdasdasd)
	if heap.len() != 1 {
		t.Fail()
	}
}

func Test_should_add_and_pop(t *testing.T) {
	heap := NewHeap()
	heap.add(2)

	if val, _ := heap.pop(); val != 2 {
		t.Fail()
	}
}

func Test_should_add_two_values_and_return(t *testing.T) {
	heap := NewHeap()
	heap.add(10)
	heap.add(5)

	if val, _ := heap.pop(); val != 5 {
		t.Fail()
	}

	if val, _ := heap.pop(); val != 10 {
		t.Fail()
	}
}

func TestShouldOrderCorrect(t *testing.T) {
	heap := NewHeap()
	heap.add(100)
	heap.add(50)
	heap.add(30)
	heap.add(110)
	heap.add(1)
	heap.add(6)

	results := []int{1, 6, 30, 50, 100, 110}

	for _, result := range results {
		value, _ := heap.pop()

		if result != value {

			t.Fatalf("Got wrong value: %d expect: %d", value, result)
		}
	}
}

func TestShouldReturnErr(t *testing.T) {
	heap := NewHeap()
	_, err := heap.pop()
	if err == nil {
		t.Fail()
	}
}
