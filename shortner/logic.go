package shortner

import (
	"time"

	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

var (
  ErrRedirectNotFound = errors.New("Redirect Not found")
)

type redirectService struct{
  redirectRepo RedirectRepository
}

func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
  return &redirectService{
    redirectRepo: redirectRepo,
  }
}

func (r *redirectService) Find(code string) (*Redirect, error) {
  return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
  redirect.Code = shortid.MustGenerate()
  redirect.CreatedAt = time.Now().UTC().Unix()
  return r.redirectRepo.Store(redirect)
}
