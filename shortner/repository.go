package shortner

// RedirectRepository interface is used to interact with store
// It allows finding Redirect by code and saving Redirect onto the store
type RedirectRepository interface {
  Find(code string) (*Redirect, error)
  Store(redirect *Redirect) (error)
}

