/*
There's an algorithms tournament taking place in which teams of programmers compete against each other to solve algorithmic problems as fast as possible. Teams compete in a round robin, where each team faces off against all other teams. Only two teams compete against each other at a time, and for each competition, one team is designated the home team, while the other team is the away team. In each competition there's always one winner and one loser; there are no ties. A team receives 3 points if it wins and 0 points if it loses. The winner of the tournament is the team that receives the most amount of points.

Given an array of pairs representing the teams that have competed against each other and an array containing the results of each competition, write a function that returns the winner of the tournament. The input arrays are named competitions and results, respectively. The competitions array has elements in the form of [homeTeam, awayTeam], where each team is a string of at most 30 characters representing the name of the team. The results array contains information about the winner of each corresponding competition in the competitions array. Specifically, results[i] denotes the winner of competitions[i], where a 1 in the results array means that the home team in the corresponding competition won and a 0 means that the away team won.

It's guaranteed that exactly one team will win the tournament and that each team will compete against all other teams exactly once. It's also guaranteed that the tournament will always have at least two teams.

Sample Input
competitions = [
  ["HTML", "C#"],
  ["C#", "Python"],
  ["Python", "HTML"],
]
results = [0, 0, 1]
Sample Output
"Python"
// C# beats HTML, Python Beats C#, and Python Beats HTML.
// HTML - 0 points
// C# -  3 points
// Python -  6 points
*/
package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// getCompetitionWinner returns the winner of the competition
// if teams array consist of HOME TEAM at index 0 and AWAY TEAM  at index 1
// if HOME TEAM wins the res will be 1 vice versa
func getCompetitionWinner(teams []string, res int) string {
	return teams[1^res]
}

// findTournamentWinner returns the winner of the tournament
// Winner of the tournament is the team with hightest score
// o(n) time | o(k) space - n is total competitions and k is total teams
func findTournamentWinner(competitions [][]string, results []int) string {
	scoreBoard := make(map[string]int)

	tournamentWinner := ""
	maxScore := 0
	for indx := 0; indx < len(competitions); indx++ {
		winner := getCompetitionWinner(competitions[indx], results[indx])
		if _, ok := scoreBoard[winner]; ok {
			scoreBoard[winner]++
			continue
		}
		scoreBoard[winner] = 1
	}

	for team, score := range scoreBoard {
		if score > maxScore {
			maxScore = score
			tournamentWinner = team
		}
	}

	return tournamentWinner
}

func TestTournamentWinner(t *testing.T) {
	competitions := [][]string{
		[]string{"HTML", "C#"},
		[]string{"C#", "Python"},
		[]string{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	require.Equal(t, "Python", findTournamentWinner(competitions, results))
}
