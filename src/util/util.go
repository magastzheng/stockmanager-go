package util

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func IsStringNotEmpty(s string) bool {
    // s != ""
    return len(s) > 0
}


