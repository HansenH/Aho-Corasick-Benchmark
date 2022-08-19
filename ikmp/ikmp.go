// Multi-pattern Match with iterative KMP
package ikmp

import (
	"github.com/paddie/gokmp"
)

type IKMP struct {
	kmps       []*gokmp.KMP
	dictionary []string
}

func NewIKMP(dictionary []string) *IKMP {
	m := &IKMP{
		kmps:       []*gokmp.KMP{},
		dictionary: dictionary,
	}
	for _, pattern := range dictionary {
		kmp, _ := gokmp.NewKMP(pattern)
		m.kmps = append(m.kmps, kmp)
	}
	return m
}

func (m *IKMP) FindAllIndex(s string) map[string][]int {
	res := make(map[string][]int)
	for i := 0; i < len(m.dictionary); i++ {
		indexes := m.kmps[i].FindAllStringIndex(s)
		if len(indexes) > 0 {
			res[m.dictionary[i]] = indexes
		}
	}
	return res
}
