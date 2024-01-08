package common

// GameData represents the state needed to run and logically
// divide up a game
type GameData struct {
	Partitions []*Partition
}

type PartitionStatus int

const (
	None PartitionStatus = iota
	Ready
	Assigned
)

// Partition is a logical piece of a game
type Partition struct {
}
