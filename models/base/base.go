package base

type ListMapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ToMap(items []ListMapItem) map[string]string {
	dataMap := make(map[string]string, 0)
	for _, item := range items {
		dataMap[item.Key] = item.Value
	}
	return dataMap
}
func ToList(data map[string]string) []ListMapItem {
	list := make([]ListMapItem, 0)
	for k, v := range data {
		list = append(list, ListMapItem{
			Key:   k,
			Value: v,
		})
	}
	return list
}
func ToListWithMapByte(data map[string][]byte) []ListMapItem {
	list := make([]ListMapItem, 0)
	for k, v := range data {
		list = append(list, ListMapItem{
			Key:   k,
			Value: string(v),
		})
	}
	return list
}
