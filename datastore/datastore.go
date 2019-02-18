package datastore

type DataStore interface {
	Search(string) (string, bool, error)
}

func NewDataStore(dsType, path string) (DataStore, error) {
	switch dsType {
	case "boltdb":
		return NewBoltDataStore(path)
	}
	return nil, nil
}
