package main

import "fmt"

func main() {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	var stations = make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	stationKey := []string{"kone", "ktwo", "kthree", "kfour", "kfive"}

	fmt.Println(setCovering(statesNeeded, stationKey, stations))
}

func setCovering(statesNeeded, stationKey []string, stations map[string][]string) []string {
	var finalStations []string

	for len(statesNeeded) > 0 {
		var bestStation string
		var statesCovered []string

		for _, station := range stationKey {
			states := stations[station]
			var covered = equalData(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}
		statesNeeded = removeData(statesNeeded, statesCovered)
		finalStations = append(finalStations, bestStation)
	}

	return finalStations
}

func equalData(statesNeeded, states []string) []string {
	var covered []string

	for _, stateNeeded := range statesNeeded {
		for _, state := range states {
			if stateNeeded == state {
				covered = append(covered, stateNeeded)
			}
		}
	}

	return covered
}

func removeData(statesNeeded []string, statesCovered []string) []string {
	for _, stateCovered := range statesCovered {
		statesNeeded = remove(statesNeeded, stateCovered)
	}

	return statesNeeded
}

func remove(statesNeeded []string, stateCovered string) []string {
	for i, stateNeeded := range statesNeeded {
		if stateCovered == stateNeeded {
			return append(statesNeeded[:i], statesNeeded[i+1:]...)
		}
	}

	return statesNeeded
}
