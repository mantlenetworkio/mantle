package rollup

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestRuntimeEnv(t *testing.T) {
	go func() {
		for true {
			value := os.Getenv("test")
			t.Log("value is: ", value)
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		count := 0
		ticker := time.NewTicker(2 * time.Second)
		for true {
			select {
			case <-ticker.C:
				count++
				os.Setenv("test", strconv.Itoa(count))
			}
		}
	}()

	errChan := make(chan struct{})
	<-errChan
}
