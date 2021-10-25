package msgpack

import (
	"github.com/kmjayadeep/go-hexagonal-example/shortner"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortner.Redirect, error) {
  redirect := &shortner.Redirect{}

  if err := msgpack.Unmarshal(input, redirect); err != nil {
    return nil, errors.Wrap(err, "serializer.Redirect.Decode")
  }

  return redirect, nil
}

func (r *Redirect) Encode(input *shortner.Redirect) ([]byte, error) {
  rawMsg, err := msgpack.Marshal(input)
  if err != nil {
    return nil, errors.Wrap(err, "serializer.Redirect.Encode")
  }

  return rawMsg, nil
}
