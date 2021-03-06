package yaml

import "errors"

type (
	// Secret is a resource that provides encrypted data
	// and pointers to external data (i.e. from vault).
	Secret struct {
		Version string `json:"version,omitempty"`
		Kind    string `json:"kind,omitempty"`
		Type    string `json:"type,omitempty"`

		Data map[string]string `json:"data,omitempty"`
	}
)

// GetVersion returns the resource version.
func (s *Secret) GetVersion() string { return s.Version }

// GetKind returns the resource kind.
func (s *Secret) GetKind() string { return s.Kind }

// Validate returns an error if the secret is invalid.
func (s *Secret) Validate() error {
	if len(s.Data) == 0 {
		return errors.New("yaml: invalid secret resource")
	}
	return nil
}
