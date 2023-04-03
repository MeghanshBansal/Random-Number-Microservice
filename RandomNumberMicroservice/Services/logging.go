package Service

import (
	
	"context"
	"fmt"
	"time"
)

type Logging struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &Logging{
		next: next,
	}
}

func (s *Logging) Fatal(ctx context.Context, err error, message ...string) {
	defer func(ctx context.Context, start time.Time) {
		p := fmt.Sprintf("Error=%v timestamp=%v", err.Error(), time.Since(start))
		if len(message) != 0 {
			p += fmt.Sprintf("%v\n", message)
		}
		fmt.Println(p)
	}(ctx, time.Now())
}

func (s *Logging) GetRandomNumber(ctx context.Context) (num int) {
	num = s.next.GetRandomNumber(ctx)
	defer func(ctx context.Context, start time.Time) {
		p := fmt.Sprintf("Number= %v took=%v\n", num, time.Since(start))
		fmt.Println(p)
	}(ctx, time.Now())
	return num
}

func (s *Logging) GetNRandomNumbers(ctx context.Context, n int) (nums []int) {
	nums = s.next.GetNRandomNumbers(ctx, n)
	defer func(ctx context.Context, start time.Time) {
		p := fmt.Sprintf("Number= %v took=%v\n", nums, time.Since(start))
		fmt.Println(p)
	}(ctx, time.Now())
	return nums
}
