package global

type ArrayDiff[I any] struct {
	RemovedValues []I
	AddedValues   []I
}

type ArrayObjectDiff[I any, C any] struct {
	RemovedValues []I
	AddedValues   []C
}
