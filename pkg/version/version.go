package version

var (
	version string
)

type Info struct {
	Version      string `json:"version"`
}


func Get() Info {
	return Info{
		Version:     version,
	}
}
