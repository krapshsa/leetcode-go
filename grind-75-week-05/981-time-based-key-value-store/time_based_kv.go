package time_based_key_value_store

type Record struct {
	timestamp int
	value     string
}
type TimeMap struct {
	valuesMap map[string][]Record
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]Record)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	_, ok := (*this).valuesMap[key]
	if !ok {
		(*this).valuesMap[key] = make([]Record, 0)
	}
	(*this).valuesMap[key] = append((*this).valuesMap[key], Record{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	value := ""

	records, ok := (*this).valuesMap[key]
	if !ok {
		return value
	}

	left, right := 0, len(records)-1
	for left <= right {
		middle := (left + right) / 2
		if records[middle].timestamp < timestamp {
			value = records[middle].value
			left = middle + 1
		} else if records[middle].timestamp > timestamp {
			right = middle - 1
		} else {
			value = records[middle].value
			return value
		}
	}

	return value
}
