package Model

type Network struct {
	Graph                  map[string][][2]interface{}
	CompressionNodes       []string
	CompressionAlreadyUsed bool
}
