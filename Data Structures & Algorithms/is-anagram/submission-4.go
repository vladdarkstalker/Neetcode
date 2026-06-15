func isAnagram(s string, t string) bool {
    var lenS = len(s)
    if lenS != len(t) {
        return false
    }

    var bufferS = make(map[rune]int, len(s))
    var bufferT = make(map[rune]int, len(s))
    for i:=0; i<lenS; i++ {
        bufferS[rune(s[i])]+=1
        bufferT[rune(t[i])]+=1
    }

    for _, el := range t {
        if bufferS[el] != bufferT[el] {
            return false
        }
    }

    return true
}
