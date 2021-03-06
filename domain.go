package whois

import (
	"fmt"
	"strings"

	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

type Domain struct {
	TLD     string // gTLD or ccTLD or newTLD
	eTLD    string // effective top-level domains
	isICANN bool   // is managed by the ICANN
}

func Parse(domain string) (*Domain, error) {
	var err error
	if domain, err = idna.ToASCII(domain); err != nil {
		return nil, err
	}

	if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") || strings.Contains(domain, "..") {
		return nil, fmt.Errorf("empty label in domain %q", domain)
	}

	var d Domain
	if d.TLD, d.isICANN = publicsuffix.PublicSuffix(domain); len(domain) <= len(d.TLD) {
		return nil, fmt.Errorf("cannot derive eTLD+1 for domain %q", domain)
	}

	i := len(domain) - len(d.TLD) - 1
	if domain[i] != '.' {
		return nil, fmt.Errorf("invalid public TLD %q for domain %q", d.TLD, domain)
	}
	d.eTLD = domain[1+strings.LastIndex(domain[:i], "."):]
	return &d, nil
}

func (d Domain) Query() string {
	arr, ok := whoisServerData[d.TLD]
	if !ok {
		arr = []string{"whois.nic." + d.TLD, "whois." + d.TLD, "nic." + d.TLD}
	}

	return <-AsyncQueryWithTimeout(d.eTLD, arr)
}
