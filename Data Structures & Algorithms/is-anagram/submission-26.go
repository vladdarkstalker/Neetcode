func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    var bufferS = make(map[byte]int, len(s))
    var bufferT = make(map[byte]int, len(s))
    for i:=0; i<len(s); i++ {
        bufferS[s[i]]+=1
        bufferT[t[i]]+=1
    }

    for _, el := range t {
        if bufferS[byte(el)] != bufferT[byte(el)] {
            return false
        }
    }

    return true
}