package main

type Ticket struct {
	Numbers []int
}

func main() {

}

func checkTickets(purchasedTickets []Ticket, n int, k int, l int) bool {
	// Generate all winning ticket combinations

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

}

// Given size of a ticket k, number of values N,
// and number of values required to win L
// Generate all winning tickets C(N,k)
// For set of purchased tickets, ensure that each winning
// ticket shares at least L values with one of the
// purchased tickets
