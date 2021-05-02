package util

func SplitMessageIrc(content string) []string {
	if len(content) < 350 {
		return []string{content}
	}
	var parts []string
	var count int
	count = len(content) / 350
	if count == 0 {
		count = 1
	}
	for i := 0; i < count; i++ {
		parts = append(parts, content[i * 350:(i * 350) + 350])
	}
	parts = append(parts, content[(count * 350):])
	return parts
}
