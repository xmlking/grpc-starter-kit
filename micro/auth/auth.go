package auth

import "context"

// Access defines the type of access a rule grants
type Access int

const (
	// AccessGranted to a resource
	AccessGranted Access = iota
	// AccessDenied to a resource
	AccessDenied
)

// Account provided by an auth provider
type Account struct {
	// ID of the account e.g. email
	ID string `json:"id"`
	// Type of the account, e.g. service
	Type string `json:"type"`
	// Issuer of the account
	Issuer string `json:"issuer"`
	// Any other associated metadata
	Metadata map[string]string `json:"metadata"`
	// Scopes the account has access to
	Scopes []string `json:"scopes"`
	// Secret for the account, e.g. the password
	Secret string `json:"secret"`
}

// Resource is an entity such as a user or
type Resource struct {
	// Name of the resource, e.g. go.micro.service.notes
	Name string `json:"name"`
	// Type of resource, e.g. service
	Type string `json:"type"`
	// Endpoint resource e.g NotesService.Create
	Endpoint string `json:"endpoint"`
}

type Rule struct {
	// ID of the rule, e.g. "public"
	ID string
	// Scope the rule requires, a blank scope indicates open to the public and * indicates the rule
	// applies to any valid account
	Scope string
	// Resource the rule applies to
	Resource *Resource
	// Access determines if the rule grants or denies access to the resource
	Access Access
	// Priority the rule should take when verifying a request, the higher the value the sooner the
	// rule will be applied
	Priority int32
}

type accountKey struct{}

// AccountFromContext gets the account from the context, which
// is set by the auth wrapper at the start of a call. If the account
// is not set, a nil account will be returned. The error is only returned
// when there was a problem retrieving an account
func AccountFromContext(ctx context.Context) (*Account, bool) {
	acc, ok := ctx.Value(accountKey{}).(*Account)
	return acc, ok
}

// ContextWithAccount sets the account in the context
func ContextWithAccount(ctx context.Context, account *Account) context.Context {
	return context.WithValue(ctx, accountKey{}, account)
}
