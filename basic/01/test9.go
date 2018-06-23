//package Another
//
//import (
//	"sync"
//	"runtime"
//	"fmt"
//	"sync/atomic"
//)
//
//var(
//	count int32
//	wg sync.WaitGroup
//)
//func main()  {
//	wg.Add(2)
//	go intCount()
//	go intCount()
//	wg.Wait()
//	fmt.Println(count)
//}
//
//func intCount()  {
//	defer wg.Done()
//	for i:=0;i<2;i++{
//		value:=atomic.LoadInt32(&count)
//		runtime.Gosched()
//		value++
//		atomic.StoreInt32(&count,value)
//	}
//}
