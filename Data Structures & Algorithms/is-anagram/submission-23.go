func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    freq := make(map[rune]int, len(s))
    for _, ch := range s {
        freq[ch]++
    }
    for _, ch := range t {
        freq[ch]--
        if freq[ch] < 0 { // ранний выход: в t больше одинаковых букв, чем в s
            return false
        }
    }

    // проверка, что все счётчики обнулились (можно пропустить, если после вычитания нигде не ушли в минус)
    for _, count := range freq {
        if count != 0 {
            return false
        }
    }
    return true
}