package _134_gas_station

import "math"

// Iterate the list and calculate the cumulative surplus, find the point with
// the lowest cumulative surplus. The next point would be the starting point because
// if there's guaranteed to be a unique answer, then the step which leads to the lowest point
// should be the last step, so that all potential increase in surplus have already been included (greedy).
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
