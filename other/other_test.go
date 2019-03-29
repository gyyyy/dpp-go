package other

import (
	"sync"
	"testing"

	"github.com/gyyyy/dpp-go/other/x08nullobject"
	"github.com/gyyyy/dpp-go/other/x09objectpool"
)

func TestNullObject(t *testing.T) {
	operation1 := x08nullobject.NewOperation("")
	operation1.Request()
	operation2 := x08nullobject.NewOperation("test")
	operation2.Request()
}

func TestObjectPool(t *testing.T) {
	var (
		pool = x09objectpool.Instance()
		wg   sync.WaitGroup
	)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			reusable := pool.AcquireReusable()
			reusable.DoSomething()
			pool.ReleaseReusable(reusable)
		}()
	}
	wg.Wait()
	pool.Destroy()
}
