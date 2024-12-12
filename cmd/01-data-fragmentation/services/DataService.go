package services

import "githuhub.com/juanjoseluisgarcia/GenLayer/model"
import "githuhub.com/juanjoseluisgarcia/GenLayer/interfaces"

type DataService struct {
	handler interfaces.DataHandler
}

func NewDataService(handler interfaces.DataHandler) *DataService {
	return &DataService{handler: handler}
}

func (ds *DataService) ReconstructData(fragments []model.Fragment) string {
	return ds.handler.ReconstructData(fragments)
}
