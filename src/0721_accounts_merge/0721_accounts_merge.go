package problem0721

import (
	"leetcode_go/types"
	"sort"
	"fmt"
)

type UnionSet = types.UnionSet

func accountsMerge(accounts [][]string) [][]string {
	var u UnionSet
	u.Init(100000)


	emailToName := make(map[string]string)
	emailToId := make(map[string]int)

	id := 0
	for _, account := range accounts{
		name := account[0]
		emailFirst := account[1]
		for j := 1; j < len(account); j++{
			email := account[j]
			emailToName[email] = name

			if _,prs := emailToId[email]; !prs{
				emailToId[email] = id
				id++
			}
			u.Union(emailToId[emailFirst],emailToId[email])
		}
	}

	emailSet := make(map[int][]string)
	for email := range emailToName{
		emailIdParent := u.Find(emailToId[email])
		emailSet[emailIdParent] = append(emailSet[emailIdParent],email)
	}

	result := [][]string{}
	for _, emails := range emailSet{
		sort.Strings(emails)
		name := emailToName[emails[0]]
		currentList := []string{}
		currentList = append(currentList,name)
		currentList = append(currentList, emails...)
		result = append(result, currentList)
	}
	return result
}