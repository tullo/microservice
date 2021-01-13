package db

var migrations map[string]FS = map[string]FS{
	"haberdasher": haberdasher,
	"stats":       stats,
}
