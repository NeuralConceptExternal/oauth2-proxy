package providers

import (
	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
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
