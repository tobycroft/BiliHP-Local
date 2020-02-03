package Sort

import (
	"sort"
)

func Ksort(arr map[string]interface{}) map[string]interface{} {

	// To store the keys in slice in sorted order
	var strs []string
	for k := range arr {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	// To perform the opertion you want
	//for _, k := range strs {
	//	fmt.Printf("%s\t%d\n", k, arr[k])
	//}
	return arr
}
