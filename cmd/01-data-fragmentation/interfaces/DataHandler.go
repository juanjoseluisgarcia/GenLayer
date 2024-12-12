package interfaces

import "githuhub.com/juanjoseluisgarcia/GenLayer/model"

// DataHandler defines the interface for handling data reconstruction
type DataHandler interface {
	ReconstructData(fragments []model.Fragment) string
}
