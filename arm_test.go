package arm_wrestle

import (
	"math/bits"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	SinkByte      byte
	SinkUint16    uint16
	SinkUint8     uint8
	SinkUint      uint
	SinkUint2     uint
	SinkFloat32   float32
	SinkInt       int
	SinkInt64     int64
	SinkInt64_2   int64
	SinkInt64_3   int64
	SinkFloat64   float64
	SinkFloat64_2 float64
	SinkInt32     int32
	SinkInt32_2   int32
	SinkInt32_3   int32
	SinkUint32    uint32
	SinkUint32_2  uint32
	SinkUint32_3  uint32
	SinkUint64    uint64
	SinkUint64_2  uint64
	SinkUint64_3  uint64
	SinkUintptr   uintptr
	SinkUintptr_2 uintptr
	SinkUintptr_3 uintptr
	SinkBool      bool
	SinkBool_2    bool
	SinkBool_3    bool
)

func BenchmarkFloat32Arithmetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkFloat32 = (float32(i)*3.14159/2.0 + 4.6) * 20.463
	}
}

func BenchmarkInt32Arithmetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkInt32 = (-int32(i) * 3) / 2
	}
}

func BenchmarkFloat64Arithmetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkFloat64 = (float64(i)*3.14159/2.0 + 4.6) * 20.463
	}
}

func BenchmarkInt64Arithmetic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkInt64 = (-int64(i) * 3) / 2
	}
}

func BenchmarkANDconstBICconst(b *testing.B) {
	const mask uint64 = 0x00FF00FF00FF00FF
	for i := 0; i < b.N; i++ {
		SinkUint64 &= mask
		SinkUint64_2 &^= mask
	}
}

func BenchmarkUint64Move(b *testing.B) {
	src := uint64(0xAAAAAAAAAAAAAAAA)

	for i := 0; i < b.N; i++ {
		SinkUint64 = src
		SinkInt64 = int64(src)
	}
}

// Combined Benchmark for ADDD, NMULD, ADDF, and NMULF sequences
func BenchmarkADD(b *testing.B) {
	intA, intX, intY := int64(3), int64(7), int64(9)
	floatA, floatX, floatY := float64(3.1), float64(7.3), float64(9.5)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SinkInt64 = intA + intX*intY         // ADDD followed by MULD sequence
		SinkInt64 = intA - intX*intY         // ADDD followed by NMULD sequence
		SinkFloat64 = floatA + floatX*floatY // ADDF followed by MULF sequence
		SinkFloat64 = floatA - floatX*floatY // ADDF followed by NMULF sequence
	}
}

func BenchmarkADDBICconst(b *testing.B) {
	x := 10000
	constant := 0x10001 // A value greater than 0xffff, which will be caught by the rule
	x2 := int32(10000)
	constant2 := int32(0x10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SinkInt = x + constant
		SinkInt32_2 = x2 + constant2
		SinkInt32 = x2 &^ constant2
	}
}

func BenchmarkADDBICconstInt64(b *testing.B) {
	x := 10000
	constant := 0x10001 // A value greater than 0xffff, which will be caught by the rule
	x2 := int64(10000)
	constant2 := int64(0x10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SinkInt = x + constant
		SinkInt64_2 = x2 + constant2
		SinkInt64 = x2 &^ constant2
	}
}

func BenchmarkWithMulDAndMulF(b *testing.B) {
	a := float64(1.5)
	x := float64(2.5)
	y := float64(3.5)

	a2 := float64(1.5)
	x2 := float64(-2.5)
	y2 := float64(3.5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SinkFloat64 = a - x*y
		SinkFloat64_2 = a2 - x2*y2
	}
}

func BenchmarkBitwiseInt32(b *testing.B) {
	x := int32(0x12345678)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SinkInt32 = ((x & 0xff) << 24) | ((x & 0xff00) << 8) | ((x & 0xff0000) >> 8) | ((x >> 24) & 0xff)
	}
}

func BenchmarkBitwiseInt64(b *testing.B) {
	y := int64(0x87654321)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SinkInt64 = ((y & 0xff) << 24) | ((y & 0xff00) << 8) | ((y & 0xff0000) >> 8) | ((y >> 24) & 0xff)
	}
}

func BenchmarkTrailingZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkInt = bits.TrailingZeros64(uint64(i))
		SinkInt += bits.TrailingZeros32(uint32(i))
		SinkInt += bits.TrailingZeros16(uint16(i))
		SinkInt += bits.TrailingZeros8(uint8(i))
	}
}

func BenchmarkLeadingZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkInt = bits.LeadingZeros64(uint64(i))
		SinkInt += bits.LeadingZeros32(uint32(i))
		SinkInt += bits.LeadingZeros16(uint16(i))
		SinkInt += bits.LeadingZeros8(uint8(i))
	}
}

func BenchmarkRotateLeft(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkUint64 = bits.RotateLeft64(uint64(i), i)
		SinkUint32 += bits.RotateLeft32(uint32(i), i)
		SinkUint16 += bits.RotateLeft16(uint16(i), i)
		SinkUint8 += bits.RotateLeft8(uint8(i), i)
	}
}

func BenchmarkOnesCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkInt += bits.OnesCount64(uint64(i))
		SinkInt += bits.OnesCount32(uint32(i))
		SinkInt += bits.OnesCount16(uint16(i))
		SinkInt += bits.OnesCount8(uint8(i))
	}
}

func BenchmarkProducerConsumerBufferedCh(b *testing.B) {
	ch := make(chan int, 1)
	// producer
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	// consumer
	for i := 0; i < b.N; i++ {
		val, ok := <-ch
		if !ok {
			b.Fatal("Channel closed prematurely")
		}
		if val != i {
			b.Errorf("Expected %d, got %d", i, val)
		}
	}
}

func BenchmarkProducerConsumerBufferedChInt64(b *testing.B) {
	ch := make(chan int64, 1)
	// producer
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- int64(i)
		}
		close(ch)
	}()

	// consumer
	for i := 0; i < b.N; i++ {
		val, ok := <-ch
		if !ok {
			b.Fatal("Channel closed prematurely")
		}
		if val != int64(i) {
			b.Errorf("Expected %d, got %d", i, val)
		}
	}
}

func BenchmarkProducerConsumerUnBufferedCh(b *testing.B) {
	ch := make(chan int)
	// producer
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	// consumer
	for i := 0; i < b.N; i++ {
		val, ok := <-ch
		if !ok {
			b.Fatal("Channel closed prematurely")
		}
		if val != i {
			b.Errorf("Expected %d, got %d", i, val)
		}
	}
}

func BenchmarkProducerConsumerUnBufferedChInt64(b *testing.B) {
	ch := make(chan int64)
	// producer
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- int64(i)
		}
		close(ch)
	}()

	// consumer
	for i := 0; i < b.N; i++ {
		val, ok := <-ch
		if !ok {
			b.Fatal("Channel closed prematurely")
		}
		if val != int64(i) {
			b.Errorf("Expected %d, got %d", i, val)
		}
	}
}

func BenchmarkGetCntxct(b *testing.B) {
	ch := make(chan int, 1) // Buffered channel

	// producer: keep reading count register
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	// consumer: check the read values
	var lastVal int
	for val := range ch {
		if val < lastVal {
			b.Errorf("Count register decreased: got %d after %d", val, lastVal)
		}
		lastVal = val
	}
}

func BenchmarkCASInt32(b *testing.B) {
	var x atomic.Int32
	for i := 0; i < b.N; i++ {
		x.CompareAndSwap(0, 1)
		x.CompareAndSwap(1, 0)
	}
}

func BenchmarkCASInt64(b *testing.B) {
	var x atomic.Int64
	for i := 0; i < b.N; i++ {
		x.CompareAndSwap(0, 1)
		x.CompareAndSwap(1, 0)
	}
}

func BenchmarkCASUint64(b *testing.B) {
	var x atomic.Uint64
	for i := 0; i < b.N; i++ {
		x.CompareAndSwap(0, 1)
		x.CompareAndSwap(1, 0)
	}
}

func BenchmarkCASUint32(b *testing.B) {
	var x atomic.Uint32
	for i := 0; i < b.N; i++ {
		x.CompareAndSwap(0, 1)
		x.CompareAndSwap(1, 0)
	}
}

func BenchmarkCASUintptr(b *testing.B) {
	var x atomic.Uintptr
	for i := 0; i < b.N; i++ {
		x.CompareAndSwap(0, 1)
		x.CompareAndSwap(1, 0)
	}
}

func BenchmarkAtomicOperationsInt64(b *testing.B) {
	var x atomic.Int64

	for i := 0; i < b.N; i++ {
		x.Store(int64(i))
		SinkInt64 = x.Load()
		SinkInt64_2 = x.Add(3)
		SinkInt64_3 = x.Swap(17)
	}
}

func BenchmarkAtomicOperationsInt32(b *testing.B) {
	var x atomic.Int32

	for i := 0; i < b.N; i++ {
		x.Store(int32(i))
		SinkInt32 = x.Load()
		SinkInt32_2 = x.Add(3)
		SinkInt32_3 = x.Swap(17)
	}
}

func BenchmarkAtomicOperationsUint64(b *testing.B) {
	var x atomic.Uint64

	for i := 0; i < b.N; i++ {
		x.Store(uint64(i))
		SinkUint64 = x.Load()
		SinkUint64_2 = x.Add(3)
		SinkUint64_3 = x.Swap(17)
	}
}

func BenchmarkAtomicOperationsUint32(b *testing.B) {
	var x atomic.Uint32

	for i := 0; i < b.N; i++ {
		x.Store(uint32(i))
		SinkUint32 = x.Load()
		SinkUint32_2 = x.Add(3)
		SinkUint32_3 = x.Swap(17)
	}
}

func BenchmarkAtomicOperationsUintptr(b *testing.B) {
	var x atomic.Uintptr

	for i := 0; i < b.N; i++ {
		x.Store(uintptr(i))
		SinkUintptr = x.Load()
		SinkUintptr_2 = x.Add(3)
		SinkUintptr_3 = x.Swap(17)
	}
}

func BenchmarkAtomicOperationsBool(b *testing.B) {
	var x atomic.Bool

	for i := 0; i < b.N; i++ {
		x.Store(true)
		x.Store(false)
		SinkBool = x.Load()
		x.Swap(true)
		x.Swap(false)
		x.CompareAndSwap(false, true)
		x.CompareAndSwap(true, false)

	}
}

const iterations = 10_000

func BenchmarkMutex(b *testing.B) {
	var m sync.Mutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Lock()
		for j := 0; j < iterations; j++ {
		}
		m.Unlock()
	}
}

func BenchmarkRWMutex_Read(b *testing.B) {
	var m sync.RWMutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.RLock()
		for j := 0; j < iterations; j++ {
		}
		m.RUnlock()
	}
}

func BenchmarkRWMutex_Write(b *testing.B) {
	var m sync.RWMutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Lock()
		for j := 0; j < iterations; j++ {
		}
		m.Unlock()
	}
}

func BenchmarkWaitGroup(b *testing.B) {
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(iterations)
		for j := 0; j < iterations; j++ {
			go func() {
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkChannel(b *testing.B) {
	ch := make(chan int, iterations)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				ch <- j
			}
		}()
		for j := 0; j < iterations; j++ {
			<-ch
		}
	}
}

func BenchmarkAtomicAdd(b *testing.B) {
	var counter int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func BenchmarkOnce(b *testing.B) {
	var once sync.Once
	action := func() {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		once.Do(action)
	}
}

func BenchmarkCond(b *testing.B) {
	cond := sync.NewCond(&sync.Mutex{})
	ready := make(chan struct{})

	b.ResetTimer()
	go func() {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			cond.L.Lock()
			ready <- struct{}{} // Notify that we're about to Wait.
			cond.Wait()
			cond.L.Unlock()
		}
		b.StopTimer()
	}()

	for i := 0; i < b.N; i++ {
		<-ready // Wait until the other goroutine is about to Wait.
		cond.L.Lock()
		cond.Signal()
		cond.L.Unlock()
	}
}

func BenchmarkPool(b *testing.B) {
	pool := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.Put(i)
		pool.Get()
	}
}

func BenchmarkMutexContended(b *testing.B) {
	var mu sync.Mutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			mu.Unlock()
		}
	})
}

func BenchmarkRWMutexContendedRead(b *testing.B) {
	var mu sync.RWMutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.RLock()
			mu.RUnlock()
		}
	})
}

func BenchmarkRWMutexContendedWrite(b *testing.B) {
	var mu sync.RWMutex
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			mu.Unlock()
		}
	})
}

func BenchmarkSemaphore(b *testing.B) {
	sema := make(chan struct{}, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sema <- struct{}{}
		<-sema
	}
}

func BenchmarkMutex2(b *testing.B) {
	var mu sync.Mutex
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		mu.Unlock()
	}
}

func BenchmarkRWMutex(b *testing.B) {
	var rwmu sync.RWMutex
	for i := 0; i < b.N; i++ {
		rwmu.RLock()
		rwmu.RUnlock()
	}
}

func BenchmarkChannel2(b *testing.B) {
	ch := make(chan struct{}, 1)
	for i := 0; i < b.N; i++ {
		ch <- struct{}{}
		<-ch
	}
}

func BenchmarkMapRWMutex(b *testing.B) {
	var mu sync.RWMutex
	m := make(map[int]int)

	b.Run("Write", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.Lock()
				m[1] = 1
				mu.Unlock()
			}
		})
	})

	b.Run("Read", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mu.RLock()
				_ = m[1]
				mu.RUnlock()
			}
		})
	})
}

func BenchmarkMapMutex(b *testing.B) {
	var mu sync.Mutex
	m := make(map[int]int)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			m[1] = 1
			_ = m[1]
			mu.Unlock()
		}
	})
}
