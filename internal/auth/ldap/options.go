package ldap

import (
	"context"
	"crypto/x509"
	"fmt"

	"github.com/hashicorp/boundary/internal/errors"
)

type options struct {
	withName                 string
	withDescription          string
	withStartTls             bool
	withInsecureTls          bool
	withDiscoverDn           bool
	withAnonGroupSearch      bool
	withUpnDomain            string
	withUserDn               string
	withUserAttr             string
	withUserFilter           string
	withGroupDn              string
	withGroupAttr            string
	withGroupFilter          string
	withCertificates         []string
	withBindDn               string
	withBindPassword         string
	withClientCertificate    string
	withClientCertificateKey []byte
	withLimit                int
	withUnauthenticatedUser  bool
	withOrderByCreateTime    bool
	ascending                bool
	withOperationalState     AuthMethodState
}

// Option - how options are passed as args
type Option func(*options) error

func getDefaultOptions() options {
	return options{
		withOperationalState: InactiveState,
	}
}

func getOpts(opt ...Option) (options, error) {
	opts := getDefaultOptions()

	for _, o := range opt {
		if err := o(&opts); err != nil {
			return opts, err
		}
	}
	return opts, nil
}

// WithName provides an optional name.
func WithName(_ context.Context, n string) Option {
	return func(o *options) error {
		o.withName = n
		return nil
	}
}

// WithDescription provides an optional description.
func WithDescription(_ context.Context, desc string) Option {
	return func(o *options) error {
		o.withDescription = desc
		return nil
	}
}

// WithStartTLS optionally enables a StartTLS command after establishing an
// unencrypted connection.
func WithStartTLS(_ context.Context) Option {
	return func(o *options) error {
		o.withStartTls = true
		return nil
	}
}

// WithInsecureTLS optional specifies to skip LDAP server SSL certificate
// validation - insecure and use with caution
func WithInsecureTLS(_ context.Context) Option {
	return func(o *options) error {
		o.withInsecureTls = true
		return nil
	}
}

// WithDiscoverDn optionally specifies to use anon bind to discover the bind DN
// of a user.
func WithDiscoverDn(_ context.Context) Option {
	return func(o *options) error {
		o.withDiscoverDn = true
		return nil
	}
}

// WithAnonGroupSearch optionally specifies to use anon bind when performing LDAP
// group searches
func WithAnonGroupSearch(_ context.Context) Option {
	return func(o *options) error {
		o.withAnonGroupSearch = true
		return nil
	}
}

// WithUpnDomain optionally specifies the userPrincipalDomain used to construct
// the UPN string for the authenticating user. The constructed UPN will appear
// as [username]@UPNDomain  Example: example.com, which will cause Boundary to
// bind as username@example.com when authenticating the user.
func WithUpnDomain(_ context.Context, domain string) Option {
	return func(o *options) error {
		o.withUpnDomain = domain
		return nil
	}
}

// WithUserDn optionally specifies a user dn used to search for user entries.
func WithUserDn(_ context.Context, dn string) Option {
	return func(o *options) error {
		o.withUserDn = dn
		return nil
	}
}

// WithUserAttr optionally specifies a user attr used to search for user entries.
func WithUserAttr(_ context.Context, attr string) Option {
	return func(o *options) error {
		o.withUserAttr = attr
		return nil
	}
}

// WithUserFilter optionally specifies a user filter used to search for user entries.
func WithUserFilter(_ context.Context, filter string) Option {
	return func(o *options) error {
		o.withUserFilter = filter
		return nil
	}
}

// WithGroupDn optionally specifies a group dn used to search for group entries.
func WithGroupDn(_ context.Context, dn string) Option {
	return func(o *options) error {
		o.withGroupDn = dn
		return nil
	}
}

// WithGroupAttr optionally specifies a group attr used to search for group entries.
func WithGroupAttr(_ context.Context, attr string) Option {
	return func(o *options) error {
		o.withGroupAttr = attr
		return nil
	}
}

// WithGroupFilter optionally specifies a group filter used to search for group entries.
func WithGroupFilter(_ context.Context, filter string) Option {
	return func(o *options) error {
		o.withGroupFilter = filter
		return nil
	}
}

// WithBindCredential optionally specifies a set of optional configuration
// parameters which allow Boundary to bind (aka authenticate) using the
// credentials provided when searching for the user entry used to authenticate
// the end user.
func WithBindCredential(ctx context.Context, dn, password string) Option {
	const op = "ldap.WithBindCredential"
	return func(o *options) error {
		switch {
		case dn == "" && password == "":
			return errors.New(ctx, errors.InvalidParameter, op, "missing both dn and password")
		case dn != "" && password == "":
			return errors.New(ctx, errors.InvalidParameter, op, "missing password")
		case dn == "" && password != "":
			return errors.New(ctx, errors.InvalidParameter, op, "missing dn")
		}
		o.withBindDn = dn
		o.withBindPassword = password
		return nil
	}
}

// WithCertificates provides optional certificates.
func WithCertificates(ctx context.Context, certs ...*x509.Certificate) Option {
	const op = "ldap.WithCertificates"
	return func(o *options) error {
		if len(certs) > 0 {
			o.withCertificates = make([]string, 0, len(certs))
			pem, err := EncodeCertificates(ctx, certs...)
			if err != nil {
				return errors.Wrap(ctx, err, op)
			}
			o.withCertificates = append(o.withCertificates, pem...)
		}
		return nil
	}
}

// WithClientCertificate provides optional configuration fields used for
// specifying a mTLS client cert for LDAP connections.
func WithClientCertificate(ctx context.Context, privKey []byte, cert *x509.Certificate) Option {
	const op = "ldap.WithClientCertificate"
	return func(o *options) error {
		if privKey != nil || cert != nil {
			switch {
			case cert == nil:
				return errors.New(ctx, errors.InvalidParameter, op, "missing certificate")
			case len(privKey) == 0:
				return errors.New(ctx, errors.InvalidParameter, op, "missing private key")
			}
			if _, err := x509.ParsePKCS8PrivateKey(privKey); err != nil {
				return errors.Wrap(ctx, err, op)
			}
			o.withClientCertificateKey = privKey
			pem, err := EncodeCertificates(ctx, cert)
			if err != nil {
				return errors.Wrap(ctx, err, op)
			}
			if len(pem) != 1 {
				return errors.New(ctx, errors.InvalidParameter, op, fmt.Sprintf("too many client certificates (%d)", len(pem)))
			}
			o.withClientCertificate = pem[0]
		}
		return nil
	}
}

// WithLimit provides an option to provide a limit.  Intentionally allowing
// negative integers.   If WithLimit < 0, then unlimited results are returned.
// If WithLimit == 0, then default limits are used for results.
func WithLimit(l int) Option {
	return func(o *options) error {
		o.withLimit = l
		return nil
	}
}

// WithUnauthenticatedUser provides an option for filtering results for
// an unauthenticated users.
func WithUnauthenticatedUser(enabled bool) Option {
	return func(o *options) error {
		o.withUnauthenticatedUser = enabled
		return nil
	}
}

// WithOrderByCreateTime provides an option to specify ordering by the
// CreateTime field.
func WithOrderByCreateTime(ascending bool) Option {
	return func(o *options) error {
		o.withOrderByCreateTime = true
		o.ascending = ascending
		return nil
	}
}

// WithOperationalState provides an option for specifying the auth method's
// operational state
func WithOperationalState(state AuthMethodState) Option {
	return func(o *options) error {
		o.withOperationalState = state
		return nil
	}
}