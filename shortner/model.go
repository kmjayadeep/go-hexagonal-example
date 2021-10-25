package shortner

// Redirect type contains the code, URL and the timestamp of when the object was Created
type Redirect struct {
	Code      string `json:"code" bson:"code" msgpack:"code"`
	URL       string `json:"url" bson:"url" msgpack:"code" validate:"empty=false & format=url"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt" msgpack:"code"`
}
