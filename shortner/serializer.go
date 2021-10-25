package shortner

// RedirectSerializer is used to encode and decode a Redirect object as bytes array
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
