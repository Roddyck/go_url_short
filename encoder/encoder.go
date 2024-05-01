package encoder

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = int64(len(charset))

func Encode(n int64) string {
    var encoded string

    if n == 0 {
        return string(charset[0])
    }

    for n > 0 {
        encoded += string(charset[n%base])
        n /= base
    }

    return encoded
}
