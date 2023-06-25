package concurrently

import (
	"fmt"
	"hash/fnv"
	"sync"
)

type Map interface {
	Store(key, value interface{})
	Load(key interface{}) (value interface{}, ok bool)
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	Range(f func(key, value interface{}) bool)
	Delete(key interface{})
}

type shard struct {
	*sync.Map
	activeIndex int32
	index       int
}

type concurrentMap struct {
	sync.RWMutex
	sn     uint32
	shards []*shard
	al     *sync.Mutex
}

func NewMap(shardsNum ...uint32) Map {
	numShards := uint32(64) // Default value

	if len(shardsNum) > 0 {
		numShards = shardsNum[0]
	}

	cm := &concurrentMap{
		sn:     numShards,
		shards: make([]*shard, numShards),
		al:     &sync.Mutex{},
	}

	var i uint32
	for i = 0; i < numShards; i++ {
		cm.shards[i] = &shard{
			Map:         &sync.Map{},
			index:       int(i),
			activeIndex: -1,
		}
	}

	return cm
}

func (cm *concurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	shard := cm.shardForKey(key)
	return shard.Load(key)
}

func (cm *concurrentMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	shard := cm.shardForKey(key)
	return shard.LoadOrStore(key, value)
}

func (cm *concurrentMap) shardForKey(key interface{}) *shard {
	index := cm.hash(key) % cm.sn
	return cm.shards[index]
}

func (cm *concurrentMap) hash(key interface{}) uint32 {
	var val uint32
	switch key := key.(type) {
	case string:
		index := fnv.New32a()
		index.Write([]byte(key))
		val = index.Sum32()
	default:
		index := fnv.New32a()
		index.Write([]byte(fmt.Sprintf("%v", key)))
		val = index.Sum32()
	}

	return val
}

func (cm *concurrentMap) Store(key, value interface{}) {
	shard := cm.shardForKey(key)
	shard.Store(key, value)
}

func (cm *concurrentMap) Delete(key interface{}) {
	shard := cm.shardForKey(key)
	shard.Delete(key)
}

func (cm *concurrentMap) Range(f func(key, value interface{}) bool) {
	for _, shard := range cm.shards {
		shard.Range(f)
	}
}
