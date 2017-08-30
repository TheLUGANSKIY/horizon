package hal

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func TestLinkBuilder(t *testing.T) {

	Convey("Link Expansion", t, func() {

		check := func(href string, base string, expectedResult string) {
			lb := LinkBuilder{mustParseURL(base)}
			result := lb.expandLink(href)
			So(result, ShouldEqual, expectedResult)
		}

		check("/root", "", "/root")
		check("/root", "//TheLUGANSKIY.org", "//TheLUGANSKIY.org/root")
		check("/root", "https://TheLUGANSKIY.org", "https://TheLUGANSKIY.org/root")
		check("//else.org/root", "", "//else.org/root")
		check("//else.org/root", "//TheLUGANSKIY.org", "//else.org/root")
		check("//else.org/root", "https://TheLUGANSKIY.org", "//else.org/root")
		check("https://else.org/root", "", "https://else.org/root")
		check("https://else.org/root", "//TheLUGANSKIY.org", "https://else.org/root")
		check("https://else.org/root", "https://TheLUGANSKIY.org", "https://else.org/root")

		// Regression: ensure that parameters are not escaped
		check("/accounts/{id}", "https://TheLUGANSKIY.org", "https://TheLUGANSKIY.org/accounts/{id}")
	})

}

func mustParseURL(base string) *url.URL {
	if base == "" {
		return nil
	}

	u, err := url.Parse(base)
	if err != nil {
		panic(err)
	}
	return u
}
