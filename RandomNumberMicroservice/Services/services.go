package Service

import "context"

type Service interface {
	GetRandomNumber(ctx context.Context) (num int)
	GetNRandomNumbers(ctx context.Context, n int) (nums []int)
}
