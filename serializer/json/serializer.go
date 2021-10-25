package json

import (
  "encoding/json"
  "github.com/pkg/errors"
  "github.com/kmjayadeep/go-hexagonal-example/shortner"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortner.Redirect, error) {
  redirect := &shortner.Redirect{}

  if err := json.Unmarshal(input, redirect); err != nil {
    return nil, errors.Wrap(err, "serializer.Redirect.Decode")
  }

  return redirect, nil
}

func (r *Redirect) Encode(input *shortner.Redirect) ([]byte, error) {
  rawMsg, err := json.Marshal(input)
  if err != nil {
    return nil, errors.Wrap(err, "serializer.Redirect.Encode")
  }

  return rawMsg, nil
}
