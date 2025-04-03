// Placeholder
package main

import "strings"

func main() {
	// Placeholder
}

// Relationship status
//
// Let's pretend that you are building a new app with social media functionality.
// Users can have relationships with other users.
//
// The two guidelines for describing relationships are:
// 1. Any user can follow any other user.
// 2. If two users follow each other, they are considered friends.
//
// This function describes the relationship that two users have with each other.
//
// Please see the sample data for examples of `socialGraph`.
//
// Params:
// - fromMember, the subject member
// - toMember, the object member
// - socialGraph, the relationship data
//
// Returns:
// - "follower" if fromMember follows toMember; "followed by" if fromMember is followed by toMember; "friends" if fromMember and toMember follow each other; "no relationship otherwise."
func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]string) string {
	fromFollowsTo := false
	toFollowsFrom := false

	// Check if fromMember follows toMember
	if following, exists := socialGraph[fromMember]["following"]; exists {
		followsList := strings.Split(following, ",")
		for _, user := range followsList {
			if user == toMember {
				fromFollowsTo = true
				break
			}
		}
	}

	// Check if toMember follows fromMember
	if following, exists := socialGraph[toMember]["following"]; exists {
		followsList := strings.Split(following, ",")
		for _, user := range followsList {
			if user == fromMember {
				toFollowsFrom = true
				break
			}
		}
	}

	// Determine the relationship status
	if fromFollowsTo && toFollowsFrom {
		return "friends"
	} else if fromFollowsTo {
		return "follower"
	} else if toFollowsFrom {
		return "followed by"
	}
	return "no relationship"
}

// Tic tac toe
//
// Tic Tac Toe is a common paper-and-pencil game.
// Players must attempt to draw a line of their symbol across a grid.
// The player that does this first is considered the winner.
//
// This function evaluates a Tic Tac Toe game board and returns the winner.
//
// Please see the sample data for examples of `board`.
//
// Params:
// - board, the representation of the Tic Tac Toe board as a square slice of slices of strings. The size of the slice will range between 3x3 to 6x6. The board will never have more than 1 winner. There will only ever be 2 unique symbols at the same time.
//
// Returns:
// - the symbol of the winner, or "NO WINNER" if there is no winner.
func ticTacToe(board [][]string) string {
	size := len(board)

	for i := 0; i < size; i++ {
		if allSame(board[i]) && board[i][0] != "" {
			return board[i][0]
		}
	}

	for i := 0; i < size; i++ {
		column := make([]string, size)
		for j := 0; j < size; j++ {
			column[j] = board[j][i]
		}
		if allSame(column) && column[0] != "" {
			return column[0]
		}
	}

	diagonal1 := make([]string, size)
	for i := 0; i < size; i++ {
		diagonal1[i] = board[i][i]
	}
	if allSame(diagonal1) && diagonal1[0] != "" {
		return diagonal1[0]
	}

	diagonal2 := make([]string, size)
	for i := 0; i < size; i++ {
		diagonal2[i] = board[i][size-1-i]
	}
	if allSame(diagonal2) && diagonal2[0] != "" {
		return diagonal2[0]
	}

	return "NO WINNER"
}

func allSame(arr []string) bool {
	if len(arr) == 0 {
		return true
	}

	first := arr[0]
	for _, item := range arr {
		if item != first {
			return false
		}
	}

	return true
}

// ETA
//
// A shuttle van service is tasked to travel one way along a predefined circular route.
// The route is divided into several legs between stops.
// The route is fully connected to itself.
//
// This function returns how long it will take the shuttle to arrive at a stop after leaving anothe rstop.
//
// Please see the sample data for examples of `routeMap`.
//
// Params:
// - firstStop, the stop that the shuttle will leave
// - secondStop, the stop that the shuttle will arrive at
// - routeMap, the data describing the routes
//
// Returns:
// - the time that it will take the shuttle to travel from firstStop to secondStop
func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	if firstStop == secondStop {
		return 0
	}

	graph := make(map[string]map[string]int)

	for route, details := range routeMap {
		stops := strings.Split(route, ",")
		if len(stops) != 2 {
			continue
		}

		from, to := stops[0], stops[1]

		if _, exists := graph[from]; !exists {
			graph[from] = make(map[string]int)
		}

		graph[from][to] = details["travel_time_mins"]
	}

	distances := make(map[string]int)
	visited := make(map[string]bool)

	for route := range routeMap {
		stops := strings.Split(route, ",")
		if len(stops) == 2 {
			from, to := stops[0], stops[1]
			if from != firstStop {
				distances[from] = int(1e9)
			}
			if to != firstStop {
				distances[to] = int(1e9)
			}
		}
	}
	distances[firstStop] = 0

	for {
		current := ""
		minDist := int(1e9)

		for stop, dist := range distances {
			if !visited[stop] && dist < minDist {
				current = stop
				minDist = dist
			}
		}

		if current == "" || minDist == int(1e9) {
			break
		}

		visited[current] = true

		if current == secondStop {
			return distances[current]
		}

		if neighbors, exists := graph[current]; exists {
			for neighbor, time := range neighbors {
				if !visited[neighbor] && distances[current]+time < distances[neighbor] {
					distances[neighbor] = distances[current] + time
				}
			}
		}
	}

	if dist, exists := distances[secondStop]; exists && dist < int(1e9) {
		return dist
	}

	return -1
}
