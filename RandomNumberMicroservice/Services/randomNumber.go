package Service

import (
	"context"
	"math/rand"
)

type RandomNumberService struct {
}

func NewRandomService(ctx context.Context) Service {
	return &RandomNumberService{}
}

func (s *RandomNumberService) GetRandomNumber(ctx context.Context) (num int) {
	num = rand.Intn(100)
	return num
}

func (s *RandomNumberService) GetNRandomNumbers(ctx context.Context, n int) (nums []int) {
	for n != 0 {
		n--
		nums = append(nums, rand.Intn(100))
	}
	return nums
}
