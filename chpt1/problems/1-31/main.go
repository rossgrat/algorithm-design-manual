package main

import "fmt"

type Ticket struct {
	Numbers []int
}

func main() {

	fmt.Println(generateCombinations(4, 2))

}

func generateCombinations(N int, k int) [][]int {
	// Base Case
	if N < k {
		return [][]int{}
	}
	if k == 0 {
		return [][]int{{}}
	}

	// Recursive Case
	// Include N
	includeN := generateCombinations(N-1, k-1)
	for i, comb := range includeN {
		includeN[i] = append(comb, N)
	}

	// Exclude N
	excludeN := generateCombinations(N-1, k)

	return append(excludeN, includeN...)
}

func ticketsIsAWinner(winningNumbers []int, purchasedNumbers []int, L int) bool {
	winningTicketNumsMap := make(map[int]bool)
	for _, winningTicketNum := range winningNumbers {
		winningTicketNumsMap[winningTicketNum] = false
	}

	count := 0
	for _, purchasedTicketNum := range purchasedNumbers {
		// If a purchased number exists on the ticket and has not yet been found
		// set it as found and increment count
		if _, ok := winningTicketNumsMap[purchasedTicketNum]; ok {
			count++
		}
		// If we found L numbers on a purchased ticket, we have successfully matched
		// a purchased ticket to a winning ticket, and can move onto the
		// next winning ticket
		if count >= L {
			return true
		}
	}
	return false
}

func checkTickets(purchasedTickets []Ticket, n int, k int, L int) bool {
	// Generate all winning ticket combinations
	winningTickets := generateCombinations(n, k)

	for _, winningTicket := range winningTickets {
		matchedTicket := false
		for _, purchasedTicket := range purchasedTickets {
			if ticketsIsAWinner(winningTicket, purchasedTicket.Numbers, L) {
				matchedTicket = true
				break
			}
		}
		if !matchedTicket {
			return false
		}
	}
	return true

	// For each winning ticket
	//  matchedPurchasedTicket = false
	// 	For each purchased ticket
	//    Create winning value set
	// 	  For values of purchased ticket
	// 	    if purchased ticket value in winning value set
	// 		and not yet flagged
	// 			flag
	// 	  If found values in set > L
	// 	    matchedPurchasedTicket = true
	//      break
	//  If matchedPurchasedTicket = false
	// 	  return false
	//
	//

}

// Given size of a ticket k, number of values N,
// and number of values required to win L
// Generate all winning tickets C(N,k)
// For set of purchased tickets, ensure that each winning
// ticket shares at least L values with one of the
// purchased tickets
