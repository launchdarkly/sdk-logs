package collections

import (
	"cmp"
	"sort"
)

func MapKeys[K comparable, V any](input map[K]V) []K {
	list := make([]K, len(input))
	index := 0
	for key, _ := range input {
		list[index] = key
		index++
	}
	return list
}

func MapForEachOrdered[K cmp.Ordered, V any](collection map[K]V, iter func(K, V)) {
	keys := MapKeys(collection)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, key := range keys {
		iter(key, collection[key])
	}
}

func MapFind[K cmp.Ordered, V any](collection map[K]V, search func(K, V) bool) (key K, value V, present bool) {
	present = false
	for k, v := range collection {
		if search(k, v) {
			key = k
			value = v
			present = true
			return
		}
	}
	return
}

func MapFilter[T any](dict map[string]T, predicate func(T) bool) (ret map[string]T) {
	ret = make(map[string]T)
	for key, value := range dict {
		if predicate(value) {
			ret[key] = value
		}
	}
	return
}
