package kubecost

import (
	"reflect"
	"testing"
	"time"
)

func TestDiff(t *testing.T) {

	start := time.Now().AddDate(0, 0, -1)
	end := time.Now()
	window1 := NewWindow(&start, &end)

	node1 := NewNode("node1", "cluster1", "123abc", start, end, window1)
	node1.CPUCost = 10
	node1b := node1.Clone().(*Node)
	node1b.CPUCost = 20
	node1Key, _ := key(node1, nil)
	node2 := NewNode("node2", "cluster1", "123abc", start, end, window1)
	node2Key, _ := key(node2, nil)
	node3 := NewNode("node3", "cluster1", "123abc", start, end, window1)
	node3Key, _ := key(node3, nil)
	node4 := NewNode("node4", "cluster1", "123abc", start, end, window1)
	node4Key, _ := key(node4, nil)
	disk1 := NewDisk("disk1", "cluster1", "123abc", start, end, window1)
	disk1Key, _ := key(disk1, nil)
	disk2 := NewDisk("disk2", "cluster1", "123abc", start, end, window1)
	disk2Key, _ := key(disk2, nil)

	cases := map[string]struct {
		inputAssetsBefore []Asset
		inputAssetsAfter  []Asset
		test              []Node
		costChange        []int
		expected          map[string]Diff[Asset]
	}{
		"added node": {
			inputAssetsBefore: []Asset{node1, node2},
			inputAssetsAfter:  []Asset{node1, node2, node3},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{node3Key: {node3, DiffAdded}},
		},
		"multiple adds": {
			inputAssetsBefore: []Asset{node1, node2},
			inputAssetsAfter:  []Asset{node1, node2, node3, node4},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{node3Key: {node3, DiffAdded}, node4Key: {node4, DiffAdded}},
		},
		"removed node": {
			inputAssetsBefore: []Asset{node1, node2},
			inputAssetsAfter:  []Asset{node2},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{node1Key: {node1, DiffRemoved}},
		},
		"multiple removes": {
			inputAssetsBefore: []Asset{node1, node2, node3},
			inputAssetsAfter:  []Asset{node2},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{node1Key: {node1, DiffRemoved}, node3Key: {node3, DiffRemoved}},
		},
		"remove all": {
			inputAssetsBefore: []Asset{node1, node2},
			inputAssetsAfter:  []Asset{},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{node1Key: {node1, DiffRemoved}, node2Key: {node2, DiffRemoved}},
		},
		"add and remove": {
			inputAssetsBefore: []Asset{node1, node2},
			inputAssetsAfter:  []Asset{node2, node3},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{node1Key: {node1, DiffRemoved}, node3Key: {node3, DiffAdded}},
		},
		"no change": {
			inputAssetsBefore: []Asset{node1, node2},
			inputAssetsAfter:  []Asset{node1, node2},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{},
		},
		"order switch": {
			inputAssetsBefore: []Asset{node2, node1},
			inputAssetsAfter:  []Asset{node1, node2},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{},
		},
		"disk add": {
			inputAssetsBefore: []Asset{disk1, node1},
			inputAssetsAfter:  []Asset{disk1, node1, disk2},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{disk2Key: {disk2, DiffAdded}},
		},
		"disk and node add": {
			inputAssetsBefore: []Asset{disk1, node1},
			inputAssetsAfter:  []Asset{disk1, node1, disk2, node2},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{disk2Key: {disk2, DiffAdded}, node2Key: {node2, DiffAdded}},
		},
		"disk and node removed": {
			inputAssetsBefore: []Asset{disk1, node1, disk2, node2},
			inputAssetsAfter:  []Asset{disk2, node2},
			costChange:        []int{},
			expected:          map[string]Diff[Asset]{disk1Key: {disk1, DiffRemoved}, node1Key: {node1, DiffRemoved}},
		},
		"cost change": {
			inputAssetsBefore: []Asset{node1},
			inputAssetsAfter:  []Asset{node1b},
			expected:          map[string]Diff[Asset]{node1Key: {node1, DiffChanged}},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			as1 := NewAssetSet(start, end, tc.inputAssetsBefore...)

			as2 := NewAssetSet(start, end, tc.inputAssetsAfter...)

			result := DiffAsset(as1.Clone(), as2.Clone())

			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("expected %+v; got %+v", tc.expected, result)
			}

		})
	}

}
