package fetcher

type ByStartTime []Session

func (a ByStartTime) Len() int {
	return len(a)
}
func (a ByStartTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStartTime) Less(i, j int) bool { return a[i].Start.Before(a[j].Start) }
