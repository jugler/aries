package models

//ServerVars variables attached to the server
type ServerVars struct {
	ImagesPortrait []string
	ImagesLanscape []string
	ImageBatch     int
}

//Page variables attached to a page
type Page struct {
	Body         []byte
	ImageRefresh int
	PageRefresh  int
	Images       []string
	TypePage     string
	TypeOfImage  string
}

//Config variables attached to the config
type Config struct {
	NextImage   bool
	TypeOfImage string
	Images      []string
	Reload      bool
}
