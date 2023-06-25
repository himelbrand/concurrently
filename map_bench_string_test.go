package concurrently

import (
	"sync"
	"testing"
)

var goroutines = []int{1, 5, 10, 50, 100, 500, 1000, 5000, 10000, 25000}

var operations = []int{250000}
var numberOfKeys = []int{25000}
var numberRangeOfKeys = []int{1000}
var rangeOperations = []int{1000}
var shards = []uint32{16, 32, 64}

func BenchmarkConcurrentMapLoadOrStore_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				for _, shardsNum := range shards {
					cm := NewMap(shardsNum)
					benchmarkMapLoadOrStore_string_int64(b, cm, routines, total, initKeys, shardsNum)
				}
			}
		}
	}
}

func BenchmarkSyncMapLoadOrStore_string(b *testing.B) {

	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				sm := &sync.Map{}
				benchmarkMapLoadOrStore_string_int64(b, sm, routines, total, initKeys, 0)
			}
		}
	}
}

func BenchmarkConcurrentMapStore_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				for _, shardsNum := range shards {
					cm := NewMap(shardsNum)
					benchmarkMapStore_string_int64(b, cm, routines, total, initKeys, shardsNum)
				}
			}
		}
	}
}

func BenchmarkSyncMapStore_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				sm := &sync.Map{}
				benchmarkMapStore_string_int64(b, sm, routines, total, initKeys, 0)
			}
		}
	}
}

func BenchmarkConcurrentMapLoad_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				for _, shardsNum := range shards {
					cm := NewMap(shardsNum)
					benchmarkMapLoad_string_int64(b, cm, routines, total, initKeys, shardsNum)
				}
			}
		}
	}
}

func BenchmarkSyncMapLoad_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				sm := &sync.Map{}
				benchmarkMapLoad_string_int64(b, sm, routines, total, initKeys, 0)
			}
		}
	}
}

func BenchmarkConcurrentMapDelete_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				for _, shardsNum := range shards {
					cm := NewMap(shardsNum)
					benchmarkMapDelete_string_int64(b, cm, routines, total, initKeys, shardsNum)
				}
			}
		}
	}
}

func BenchmarkSyncMapDelete_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range operations {
			for _, initKeys := range numberOfKeys {
				sm := &sync.Map{}
				benchmarkMapDelete_string_int64(b, sm, routines, total, initKeys, 0)
			}
		}
	}
}

func BenchmarkConcurrentMapRange_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range rangeOperations {
			for _, initKeys := range numberRangeOfKeys {
				for _, shardsNum := range shards {
					cm := NewMap(shardsNum)
					benchmarkMapRange_string_int64(b, cm, routines, total, initKeys, shardsNum)
				}
			}
		}
	}
}

func BenchmarkSyncMapRange_string(b *testing.B) {
	for _, routines := range goroutines {
		for _, total := range rangeOperations {
			for _, initKeys := range numberRangeOfKeys {
				sm := &sync.Map{}
				benchmarkMapRange_string_int64(b, sm, routines, total, initKeys, 0)
			}
		}
	}
}
