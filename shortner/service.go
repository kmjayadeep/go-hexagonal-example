package shortner

// RedirectService contains business logic to find and store Redirect objects
type RedirectService interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
