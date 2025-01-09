func prefixCount(words []string, pref string) int {
    count := 0
    for i := 0; i < len(words); i++ {
        if _, found := strings.CutPrefix(words[i], pref); found {
            count++
        }
    }
    return count
}
