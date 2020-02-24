package goroutine

import (
	"errors"
	"time"

	"github.com/cheggaaa/pb"
)

//Run to start execution of functions in tasks slice in N goroutines until maxErrCount errors
func Copy(from string, to string, limit int, offset int) error {
	count := 100000
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
	if checkErrCount(&errCount, maxErrCount, &mur) {
		return errors.New("Too much errors")
	}
	return nil
}
