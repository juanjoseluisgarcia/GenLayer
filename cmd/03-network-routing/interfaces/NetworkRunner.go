package interfaces

type NetworkRunner interface {
	FindMinimumLatencyPath(source string, target string) int
}
