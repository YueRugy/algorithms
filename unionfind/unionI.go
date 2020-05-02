package unionfind

type Union interface {
	IsSame(v1,v2 int) bool
	Find(v int) int
	Union(v1,v2 int)
}
