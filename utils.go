package main

import (
    "crypto/rand"
    "os/user"
)

var SimpleTimeFmt = "2006-01-02 15:04:05"

func RandString(n int) string {
    const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    var bytes = make([]byte, n)
    rand.Read(bytes)
    for i, b := range bytes {
        bytes[i] = alphanum[b % byte(len(alphanum))]
    }
    return string(bytes)
}

func GetHomeDir() string {
    usr, err := user.Current()
    if err != nil {
        return ""
    }
    return usr.HomeDir
}

