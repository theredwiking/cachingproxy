package origin

type Origin struct {
	Url string
}

var validCodes = [3]int{200, 201, 202}

// Creates an new Origin struct from provided url
func NewOrigin(url string) *Origin {
	return &Origin{Url: url}
}
