package idna

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToACE(t *testing.T) {
	cases := []struct {
		domainInUnicode     string
		expectedDomainInACE string
	}{
		{domainInUnicode: "com", expectedDomainInACE: "com"}, // top level domain
		{domainInUnicode: "go-proverbs.github.io", expectedDomainInACE: "go-proverbs.github.io"},
		{domainInUnicode: "lagom-är-bäst.com", expectedDomainInACE: "xn--lagom-r-bst-q8ad.com"},
		{domainInUnicode: "котомастер.рф", expectedDomainInACE: "xn--80aknijargfe.xn--p1ai"},
		{domainInUnicode: "en.википедия.org", expectedDomainInACE: "en.xn--b1aeclack5b4j.org"},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			v, err := ToACE(tt.domainInUnicode)
			require.NoError(t, err)
			require.Equal(t, tt.expectedDomainInACE, v)
		})
	}
}
