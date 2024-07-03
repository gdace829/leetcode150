package NumAndSring

import "math/rand"

type RandomizedSet struct {
	nums     []int
	indexMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{nums: []int{}, indexMap: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indexMap[val]; ok {
		return false
	} else {
		this.nums = append(this.nums, val)
		this.indexMap[val] = len(this.nums) - 1
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.indexMap[val]; ok {
		index := this.indexMap[val]
		this.nums[index] = this.nums[len(this.nums)-1]
		this.indexMap[this.nums[len(this.nums)-1]] = index
		this.nums = this.nums[:len(this.nums)-1]
		delete(this.indexMap, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
