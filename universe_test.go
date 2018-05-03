package artnet

import (
	"sync"
	"testing"
	"time"
)

func TestUniverse(t *testing.T) {

	_, n := NewNode("192.168.1.13:6454")

	u := NewUniverse(0, n, time.Millisecond*10)
	u.Set(25, 255)
	u.Send()

	wg := sync.WaitGroup{}
	wg.Add(1)

	// for i := 0; i <= 23; i++ {
	// 	go func(i int) {
	// 		for {
	// 			if i%3 == 0 {
	// 				u.MultiSet(i*3, []uint8{255, 0, 0})
	// 			} else if i%3 == 1 {
	// 				u.MultiSet(i*3, []uint8{255, 255, 0})
	// 			} else {
	// 				u.MultiSet(i*3, []uint8{255, 0, 255})
	// 			}

	// 			time.Sleep(100 * time.Millisecond)
	// 			u.MultiSet(i*3, []uint8{0, 0, 0})
	// 			time.Sleep(100 * time.Millisecond)
	// 		}
	// 	}(i)
	// }

	go func() {
		i := 0
		m := 75
		for {
			u.ClearMultiSet(i%m, []uint8{0, 255, 255})
			u.MultiSet((i+3)%m, []uint8{0, 80, 255})
			u.MultiSet((i+6)%m, []uint8{0, 0, 255})
			u.MultiSet((i+9)%m, []uint8{255, 0, 0})
			u.MultiSet((i+12)%m, []uint8{255, 255, 0})
			u.MultiSet((i+15)%m, []uint8{80, 255, 0})
			u.MultiSet((i+18)%m, []uint8{0, 255, 0})

			u.MultiSet((i+24)%m, []uint8{0, 255, 255})
			u.MultiSet((i+27)%m, []uint8{0, 80, 255})
			u.MultiSet((i+30)%m, []uint8{0, 0, 255})
			u.MultiSet((i+33)%m, []uint8{255, 0, 0})
			u.MultiSet((i+36)%m, []uint8{255, 255, 0})
			u.MultiSet((i+39)%m, []uint8{80, 255, 0})
			u.MultiSet((i+42)%m, []uint8{0, 255, 0})

			u.MultiSet((i+48)%m, []uint8{0, 255, 255})
			u.MultiSet((i+51)%m, []uint8{0, 80, 255})
			u.MultiSet((i+54)%m, []uint8{0, 0, 255})
			u.MultiSet((i+57)%m, []uint8{255, 0, 0})
			u.MultiSet((i+60)%m, []uint8{255, 255, 0})
			u.MultiSet((i+63)%m, []uint8{80, 255, 0})
			u.MultiSet((i+66)%m, []uint8{0, 255, 0})

			time.Sleep(200 * time.Millisecond)
			i = i + 3
		}
	}()

	// go func() {
	// 	i := 0
	// 	for {
	// 		u.ClearMultiSet(i+15, []uint8{0, 255, 255, 255, 0, 255, 0, 255, 0})
	// 		u.MultiSet(i, []uint8{0, 255, 255, 255, 0, 255, 255, 255, 0})
	// 		time.Sleep(20 * time.Millisecond)
	// 		if i == 69 {
	// 			i = 0
	// 		} else {
	// 			i = i + 3
	// 		}

	// 	}
	// 	wg.Done()
	// }()

	wg.Wait()
}
