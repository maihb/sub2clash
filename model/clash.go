package model

import "github.com/bestnite/sub2clash/parser"

type ClashType int

const (
	Clash ClashType = 1 + iota
	ClashMeta
)

func GetSupportProxyTypes(clashType ClashType) map[string]bool {
	supportProxyTypes := make(map[string]bool)

	for _, parser := range parser.GetAllParsers() {
		if clashType == Clash {
			if parser.SupportClash() {
				supportProxyTypes[parser.GetType()] = true
			}
		} else if clashType == ClashMeta {
			if parser.SupportMeta() {
				supportProxyTypes[parser.GetType()] = true
			}
		}
	}

	return supportProxyTypes
}
