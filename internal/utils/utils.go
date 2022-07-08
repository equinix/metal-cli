package utils

import "strings"

func ChunkShort(s string) string {
	return chunk(s, 60)
}

func ChunkLong(s string) string {
	return chunk(s, 80)
}

func chunk(s string, limit int) string {

	if len(s) <= limit {
		return s
	}
	
	var charSlice []rune
	result := ""
	start  := "<"
	end    := ">"

	for _, char := range s {
		charSlice = append(charSlice, char)
	}
	for len(charSlice) >= 1 {
		str := string(charSlice)
		startIdx := strings.Index(str, start)
		if startIdx != -1 {
			endIdx := strings.Index(str, end)
			if endIdx != -1 {
				if len(charSlice[:endIdx]) > limit {
					limit = startIdx
					if len(charSlice[startIdx:endIdx]) > limit && startIdx == 0{
						limit = endIdx+1
					}
				}
			}
		}
		result = result + strings.Trim(string(charSlice[:limit]), " ") + "\r\n"
		charSlice = charSlice[limit:]
		if len(charSlice) < limit {
			limit = len(charSlice)
		}
	}
	return result
}