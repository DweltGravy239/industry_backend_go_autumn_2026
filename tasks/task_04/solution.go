package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}
	s := Stats{}
	result := nums[1] - nums[0]
	s.Min = result
	s.Max = result
	for i := 1; i < len(nums); i++ {
		a := nums[i] - nums[i-1]
		s.Count += 1
		s.Sum += a
		if s.Min > a {
			s.Min = a
		}
		if s.Max < a {
			s.Max = a
		}
	}
	return s
}
