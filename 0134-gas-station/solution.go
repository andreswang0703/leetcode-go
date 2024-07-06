package _134_gas_station

import "math"

// The solution is finding the lowest accumulative surplus point, and choose the index after that.
// The logic behind this solution is:
//  1. If there is a solution, the cumulative surplus of all points is guaranteed to be non-negative.
//  2. For choosing a starting point k, we need to choose one that the cumulative surplus of [k..n-1] is bigger
//     than [0..k-1], so that the surplus of [k..n-1] can cover the expense of [0..k-1].
//  3. Points that satisfy the previous rule don't automatically guarantee the non-existence of points that might dip
//     below 0. However, when condition 1 is satisfied, we can choose the point after the lowest cumulative surplus
//     because firstly it guarantees 2, and secondly it guarantees the non-existence of points that dips below 0 after k,
//     and that's because if there were, that point would be the actual lowest point, it also guarantees no point would dip below 0
//     before k, because the cumulative surplus of [k..n-1] is sufficient to cover all cost before k.
func canCompleteCircuit(gas []int, cost []int) int {
	minTotal := math.MaxInt8
	total := 0
	idx := 0

	for i := range gas {
		surplus := gas[i] - cost[i]
		total += surplus
		if total <= minTotal {
			minTotal = total
			idx = i
		}
	}
	if total >= 0 {
		return (idx + 1) % len(gas)
	} else {
		return -1
	}
}
