package profile

// Album represents data about a record Album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Collection represents a set of Albums.
type Collection struct {
	Albums []Album
}

func New() *Collection {
	return &Collection{}
}

func (collection *Collection) Add(album Album) {
	collection.Albums = append(collection.Albums, album)
}

func (collection *Collection) GetAll() []Album {
	return collection.Albums
}
