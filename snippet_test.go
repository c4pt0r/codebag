package main

import (
    "testing"
    "time"
    "log"
)

func init() {
    InitDb()
}

func Test_AddSnippet(t *testing.T) {
    s := NewSnippet("go", "go example", time.Now(), []byte("\xff\x33"))
    if _, err := AddSnippet(s); err == nil {
        log.Println(s.id)
    } else {
        t.Error(err)
    }

    s.desc = "go rocks"
    err := UpdateSnippet(s.id, s)
    if err != nil {
        t.Error(err)
    }

    id := s.id
    s1, err := FetchSnippet(id)
    if err != nil {
        t.Error(err)
    } else {
        log.Println(s1.id, s1.desc)
    }

    err = RemoveSnippet(id)
    if err != nil {
        t.Error(err)
    }

    lists, err := ListSnippets(0, 10)
    for _, s := range lists {
        log.Println(s.id, s.desc, s.date)
    }
}
