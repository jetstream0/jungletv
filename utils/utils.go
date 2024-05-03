package utils

import (
	"context"
	"net/netip"
	"regexp"
	"strings"

	"github.com/JohannesKaufmann/html-to-markdown/escape"
	"github.com/palantir/stacktrace"
	"github.com/sethvargo/go-limiter"
	"github.com/sethvargo/go-limiter/memorystore"
)

// GetUniquifiedIP returns a uniquified version of an IP address
// (sets the lower bits of a IPv6 to zero, leaves IPv4 untouched)
func GetUniquifiedIP(remoteAddress string) string {
	addr, err := netip.ParseAddr(remoteAddress)
	if err != nil {
		return remoteAddress
	}
	keepTopBits := 64
	if keepTopBits > addr.BitLen() {
		keepTopBits = addr.BitLen()
	}
	prefix, err := addr.Prefix(keepTopBits)
	if err != nil {
		return remoteAddress
	}
	return prefix.Addr().Unmap().WithZone("").String()
}

// ReplaceAllStringSubmatchFunc is a version of func (*regexp.Regexp) ReplaceAllStringFunc
// that passes submatches to the callback.
// It follows the "semantic naming" convention for functions in the regexp package.
// Based on the implementation found at
// https://elliotchance.medium.com/go-replace-string-with-regular-expression-callback-f89948bad0bb
func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0
	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			group := ""
			if v[i] >= 0 {
				group = str[v[i]:v[i+1]]
			}
			groups = append(groups, group)
		}
		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}
	return result + str[lastIndex:]
}

// ReplaceAllStringSubmatchFuncExcludingInside works like ReplaceAllStringSubmatchFunc,
// but it does not perform any replacements in regions that match the regexp passed as the second argument
func ReplaceAllStringSubmatchFuncExcludingInside(re, excludeInside *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0
	for _, v := range excludeInside.FindAllIndex([]byte(str), -1) {
		// pass all from the last match up to here excluding what's in the excludeInside match
		result += ReplaceAllStringSubmatchFunc(re, str[lastIndex:v[0]], repl)
		// what's inside the match goes as-is
		result += str[v[0]:v[1]]
		lastIndex = v[1]
	}
	// append what's outside of any match. we need to pass that through as well
	result += ReplaceAllStringSubmatchFunc(re, str[lastIndex:], repl)
	return result
}

// CastStringLikeSlice converts between slices of string-like types
func CastStringLikeSlice[T ~string, V ~string](in []T) []V {
	result := make([]V, len(in))
	for i, t := range in {
		result[i] = V(t)
	}
	return result
}

// SliceToSet returns a map-based set with the elements of the specified slice
func SliceToSet[T comparable](s []T) map[T]struct{} {
	set := make(map[T]struct{})
	for _, item := range s {
		set[item] = struct{}{}
	}
	return set
}

var replacer = strings.NewReplacer(
	`~`, `\~`,
)

// EscapeMarkdown is like escape.MarkdownCharacters but aware of JungleTV-specific chat markdown extensions
func EscapeMarkdownCharacters(s string) string {
	s = escape.MarkdownCharacters(s)

	s = replacer.Replace(s)

	// TODO do something about emotes and timestamps

	return s
}

// NewRateLimiterMemoryStoreWithContext creates a new in-memory rate limiter with the given config and takes care of its closure when the context is done
func NewRateLimiterMemoryStoreWithContext(ctx context.Context, config *memorystore.Config) (limiter.Store, error) {
	s, err := memorystore.New(config)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	go func(ctx context.Context, s limiter.Store) {
		<-ctx.Done()
		s.Close(ctx)
	}(ctx, s)
	return s, nil
}
