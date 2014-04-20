package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "fmt"
    "time"
    "errors"
)

type Snippet struct {
    id      int64 // id maybe -1, when dirty
    tags    string
    desc    string
    date    time.Time
    content []byte // json format content, [{"filename": xxx, "source":xxxxxxxxxxx}]
}

func NewSnippet(tags, desc string, date time.Time, content []byte) *Snippet {
    ret := &Snippet{
        -1,
        tags,
        desc,
        date,
        content,
    }
    return ret
}

var createSnippetTblSql = `
create table snippet (id integer primary key, tags text, desc text, date text, content blob)
`
var db *sql.DB

func InitDb() {
    var err error
    dbFile := GlobalCfg.DbFile
    db, err = sql.Open("sqlite3", dbFile)
    if err != nil {
        log.Fatal(err)
    }
    // ignore table exists error
    db.Exec(createSnippetTblSql)
}

func AddSnippet(s *Snippet) (int64, error) {
    tx, err := db.Begin()
	if err != nil {
        return -1, err
	}
	stmt, err := tx.Prepare("insert into snippet(tags, desc, date, content) values(?, ?, ?, ?)")
	if err != nil {
        return -1, err
	}
	defer stmt.Close()
    res, err := stmt.Exec(s.tags, s.desc, s.date.Format(time.RFC3339), s.content)
    if err != nil {
        return -1, err
    }
	tx.Commit()
    id, err := res.LastInsertId()
    s.id = id
    return id, err
}

var ErrNoSuchSnippet error = errors.New("no such snippet")

func FetchSnippet(id int64) (*Snippet, error) {
    rows, err := db.Query(fmt.Sprintf("select id, tags, desc, date, content from snippet where id=%d", id))
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        s := &Snippet{}
        var date string
        rows.Scan(&s.id, &s.tags, &s.desc, &date, &s.content)
        s.date, _ = time.Parse(time.RFC3339, date)
        return s, nil
    }
    return nil, ErrNoSuchSnippet
}

func RemoveSnippet(id int64) error {
    _, err := db.Exec(fmt.Sprintf("delete from snippet where id=%d", id))
    return err
}

func UpdateSnippet(id int64, s *Snippet) error {
    tx, err := db.Begin()
	if err != nil {
        return err
	}
	stmt, err := tx.Prepare(fmt.Sprintf("update snippet set tags=?, desc=?, date=?, content=? where id=%d", id))
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(s.tags, s.desc, s.date.Format(time.RFC3339), s.content)
    if err != nil {
        return err
    }
	tx.Commit()
    return nil
}

func ListSnippets(offset, count int64) ([]*Snippet, error) {
    var sql string
    if count == 0 {
        sql = fmt.Sprintf("select id, tags, desc, date  from snippet where id > %d order by id desc ", offset)
    } else {
        sql = fmt.Sprintf("select id, tags, desc, date  from snippet order by id desc where id > %d limit %d",
                    offset,
                    count)
    }
    rows, err := db.Query(sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var ret []*Snippet
    for rows.Next() {
        s := &Snippet{}
        var date string
        rows.Scan(&s.id, &s.tags, &s.desc, &date)
        s.date, _ = time.Parse(time.RFC3339, date)
        ret = append(ret, s)
    }
    return ret, nil
}

