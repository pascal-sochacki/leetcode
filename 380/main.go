package main

import "math/rand"

func main() {
	test := Constructor()
	println(test.Remove(0))
	println(test.Remove(0))
	println(test.Insert(0))
	println(test.GetRandom())
	println(test.Remove(0))
	println(test.Insert(0))

}

type RandomizedSet struct {
	hashmap map[int]int
	list    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashmap: map[int]int{},
		list:    []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashmap[val]; ok {
		return false
	} else {

		this.hashmap[val] = len(this.list)
		this.list = append(this.list, val)

		return true
	}

}

func (this *RandomizedSet) Remove(val int) bool {
	if pos, ok := this.hashmap[val]; ok {

		if pos == len(this.list)-1 {
			this.list = this.list[0 : len(this.list)-1]
			delete(this.hashmap, val)
		} else {

			this.list[pos] = this.list[len(this.list)-1]
			this.list = this.list[0 : len(this.list)-1]

			delete(this.hashmap, val)
			this.hashmap[this.list[pos]] = pos
		}

		return true
	} else {
		return false

	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
