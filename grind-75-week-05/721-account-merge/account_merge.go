package account_merge

import "sort"

type people struct {
	name string
}

func accountsMerge(accounts [][]string) [][]string {
	mailMap := make(map[string]*people)
	peoples := make(map[*people][]string, 0)

	for i := 0; i < len(accounts); i++ {
		account := accounts[i]
		name := account[0]

		var p *people = nil
		for j := 1; j < len(account); j++ {
			mail := account[j]
			value, ok := mailMap[mail]
			if ok {
				p = value
			}
		}

		if p == nil {
			p = &people{name}
			peoples[p] = make([]string, 0)
			for j := 1; j < len(account); j++ {
				mail := account[j]
				peoples[p] = append(peoples[p], mail)
				mailMap[mail] = p
			}
		} else {
			for j := 1; j < len(account); j++ {
				mail := account[j]
				p2, ok := mailMap[mail]
				if ok {
					if p != p2 {
						for _, mail := range peoples[p2] {
							peoples[p] = append(peoples[p], mail)
							mailMap[mail] = p
						}
						delete(peoples, p2)
					}
				} else {
					mail := account[j]
					peoples[p] = append(peoples[p], mail)
					mailMap[mail] = p
				}
			}
		}
	}

	results := make([][]string, 0)
	for p, mails := range peoples {
		sort.Strings(mails)
		result := make([]string, 0)
		result = append(result, p.name)
		for index, mail := range mails {
			if index > 0 && mails[index] == mails[index-1] {
				continue
			}
			result = append(result, mail)
		}
		results = append(results, result)
	}

	return results
}
