package providers

import (
	"context"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/sessions"
)

const ncOIDCProviderName = "NC OIDC"

// NC custom provider based on OIDCProvider
type NCOIDCProvider struct {
	*OIDCProvider
}

func NewNCOIDCProvider(p *ProviderData, opts options.OIDCOptions) *NCOIDCProvider {
	p.setProviderDefaults(providerDefaults{
		name: ncOIDCProviderName,
	})

	provider := &NCOIDCProvider{
		OIDCProvider: &OIDCProvider{
			ProviderData: p,
		},
	}

	return provider
}

var auth_count = 10

func (p *NCOIDCProvider) Authorize(_ context.Context, s *sessions.SessionState) (bool, error) {

	auth_count -= 1
	if auth_count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
