package shortner

import (
	"time"

	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not found")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

// Create a new RedirectService
func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo: redirectRepo,
	}
}

// Find a redirect object using code
// Returns ErrRedirectNotFound if it is not present
func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

// Takes a Redirect object, Generate a code and stores it
// The fields Code and CreatedAt are added to the Redirect object
func (r *redirectService) Store(redirect *Redirect) error {
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}
