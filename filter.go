package zerolog

import (
	"github.com/RoaringBitmap/roaring"
	"hash/fnv"
	"sync"
)

//log filter switch
var ylogFilter = roaring.New()
//is filter in LogFilters
var ylogFilterSaved = roaring.New()
//all filter string
var ylogFilterAll = make(map[string]bool)
var ylogFiltersLock = sync.Mutex{}

func ystrHash(str string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return h.Sum32()
}
func ybyteHash(b []byte) uint32 {
	h := fnv.New32a()
	h.Write(b)
	return h.Sum32()
}

func isYlogFilterOn(filter string) bool {
	h := ystrHash(filter)

	if ylogFilter.Contains(h) {
		return true
	}

	return false

}
func isYlogFilterOnAndSaved(filter string) (on bool, saved bool) {
	h := ystrHash(filter)

	on = ylogFilter.Contains(h)
	saved = ylogFilterSaved.Contains(h)

	return

}

func isYlogFilterSaved(filter string) bool {
	h := ystrHash(filter)

	if ylogFilterSaved.Contains(h) {
		return true
	}

	return false

}

func saveYlogFilter(filter string) {
	h := ystrHash(filter)

	ylogFilterSaved.Add(h)
	ylogFiltersLock.Lock()
	ylogFilterAll[filter] = true
	ylogFiltersLock.Unlock()
}

func YlogFilterOn(filter string) {
	h := ystrHash(filter)

	if ylogFilter.Contains(h) {
		return

	}

	if ylogFilterSaved.Contains(h) {
		return
	}

	ylogFilter.Add(h)
	saveYlogFilter(filter)
}
func YlogFilterOff(filter string) {
	h := ystrHash(filter)

	ylogFilter.Remove(h)
}
