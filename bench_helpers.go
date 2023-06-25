package concurrently

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
)

func initMapWithStringKeys(m Map, key *int64, numberOfInitialKeys int) {
	for i := 0; i < numberOfInitialKeys; i++ {
		v := atomic.AddInt64(key, 1)
		k := fmt.Sprintf("string_key_%v", v)
		m.Store(k, v)
	}
}

func benchmarkMapStore_string_int64(b *testing.B, m Map, numGoroutines, total, initKeys int, shards uint32) {
	var key int64
	initMapWithStringKeys(m, &key, initKeys)
	b.ResetTimer()
	b.Run(fmt.Sprintf("%d goroutines,%d total,%d keys,%v shards", numGoroutines, total, initKeys, shards), func(b *testing.B) {
		var wg sync.WaitGroup
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				limit := total / numGoroutines
				if limit == 0 {
					limit = total
				}

				for j := 0; j < limit; j++ {
					v := int64(rand.Uint64())
					k := fmt.Sprintf("string_key_%v", v)
					m.Store(k, v)
				}
			}()
		}
		wg.Wait()
	})
}

func benchmarkMapLoadOrStore_string_int64(b *testing.B, m Map, numGoroutines, total, initKeys int, shards uint32) {
	var key int64
	initMapWithStringKeys(m, &key, initKeys)
	atomic.StoreInt64(&key, int64(initKeys)/2)
	b.ResetTimer()
	b.Run(fmt.Sprintf("%d goroutines,%d total,%d keys,%v shards", numGoroutines, total, initKeys, shards), func(b *testing.B) {
		var wg sync.WaitGroup
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				limit := total / numGoroutines
				if limit == 0 {
					limit = total
				}

				for j := 0; j < limit; j++ {
					v := atomic.AddInt64(&key, 1)
					k := fmt.Sprintf("string_key_%v", v)
					m.LoadOrStore(k, v)

				}
			}()
		}
		wg.Wait()
	})
}

func benchmarkMapLoad_string_int64(b *testing.B, m Map, numGoroutines, total, initKeys int, shards uint32) {
	var key int64
	initMapWithStringKeys(m, &key, initKeys)
	atomic.StoreInt64(&key, 0)
	b.ResetTimer()
	b.Run(fmt.Sprintf("%d goroutines,%d total,%d keys,%v shards", numGoroutines, total, initKeys, shards), func(b *testing.B) {
		var wg sync.WaitGroup
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				limit := total / numGoroutines
				if limit == 0 {
					limit = total
				}

				for j := 0; j < limit; j++ {
					v := atomic.AddInt64(&key, 1)
					k := fmt.Sprintf("string_key_%v", v)
					m.Load(k)

				}
			}()
		}
		wg.Wait()
	})
}

func benchmarkMapDelete_string_int64(b *testing.B, m Map, numGoroutines, total, initKeys int, shards uint32) {
	var key int64
	initMapWithStringKeys(m, &key, initKeys)
	atomic.StoreInt64(&key, 0)
	b.ResetTimer()
	b.Run(fmt.Sprintf("%d goroutines,%d total,%d keys,%v shards", numGoroutines, total, initKeys, shards), func(b *testing.B) {
		var wg sync.WaitGroup
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				limit := total / numGoroutines
				if limit == 0 {
					limit = total
				}

				for j := 0; j < limit; j++ {
					v := atomic.AddInt64(&key, 1)
					k := fmt.Sprintf("string_key_%v", v)
					m.Delete(k)
				}
			}()
		}
		wg.Wait()
	})
}

func benchmarkMapRange_string_int64(b *testing.B, m Map, numGoroutines, total, initKeys int, shards uint32) {
	var key int64
	initMapWithStringKeys(m, &key, initKeys)
	b.ResetTimer()
	b.Run(fmt.Sprintf("%d goroutines,%d total,%d keys,%v shards", numGoroutines, total, initKeys, shards), func(b *testing.B) {
		var wg sync.WaitGroup
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				limit := total / numGoroutines
				if limit == 0 {
					limit = total
				}

				for j := 0; j < limit; j++ {
					m.Range(func(key, value interface{}) bool {
						return true
					})
				}
			}()
		}
		wg.Wait()
	})
}
