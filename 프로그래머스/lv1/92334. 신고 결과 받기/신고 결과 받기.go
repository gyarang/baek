import (
	"strings"
)

type User struct {
	id          string
	reports     []string
	reportedCnt int
}

func (user *User) report(reportedUser *User) {
	for _, report := range user.reports {
		if report == reportedUser.id {
			return
		}
	}
	user.reports = append(user.reports, reportedUser.id)
	reportedUser.reportedCnt += 1
}

func solution(id_list []string, report []string, k int) []int {
	idMap := make(map[string]*User)

	for _, id := range id_list {
		_, exists := idMap[id]
		if !exists {
			idMap[id] = &User{
				id:          id,
				reports:     []string{},
				reportedCnt: 0,
			}
		}
	}

	for _, rep := range report {
		split := strings.Split(rep, " ")
		idMap[split[0]].report(idMap[split[1]])
	}

	result := make([]int, int(len(id_list)))
	for i, id := range id_list {
		blockedCnt := 0
		user := idMap[id]
		for _, report := range user.reports {
			if idMap[report].reportedCnt >= k {
				blockedCnt += 1
			}
		}
		result[i] = blockedCnt
	}

	return result
}