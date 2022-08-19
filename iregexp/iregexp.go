// Multi-pattern Match with iterative regular expression
package iregexp

import (
	"regexp"
)

type IRegexp struct {
	regexps    []*regexp.Regexp
	dictionary []string
}

func NewIRegexp(dictionary []string) *IRegexp {
	m := &IRegexp{
		regexps:    []*regexp.Regexp{},
		dictionary: dictionary,
	}
	for _, pattern := range dictionary {
		re, _ := regexp.Compile(pattern)
		m.regexps = append(m.regexps, re)
	}
	return m
}

func (m *IRegexp) FindAllIndex(s string) map[string][][]int {
	res := make(map[string][][]int)
	for i := 0; i < len(m.dictionary); i++ {
		indexes := m.regexps[i].FindAllStringIndex(s, -1)
		if len(indexes) > 0 {
			res[m.dictionary[i]] = indexes
		}
	}
	return res
}
