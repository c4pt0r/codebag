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
    ftype   string
    desc    string
    date    time.Time
    content []byte
}

func NewSnippet(ftype, desc string, date time.Time, content []byte) *Snippet {
    ret := &Snippet{
        -1,
        ftype,
        desc,
        date,
        content,
    }
    return ret
}

var createSnippetTblSql = `
create table snippet (id integer primary key, type text, desc text, date text, content blob)
`
var db *sql.DB

func InitDb() {
    log.Println("init db...")
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
	stmt, err := tx.Prepare("insert into snippet(type, desc, date, content) values(?, ?, ?, ?)")
	if err != nil {
        return -1, err
	}
	defer stmt.Close()
    res, err := stmt.Exec(s.ftype, s.desc, s.date.Format(time.RFC3339), s.content)
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
    rows, err := db.Query(fmt.Sprintf("select id, type, desc, date, content from snippet where id=%d", id))
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        s := &Snippet{}
        var date string
        rows.Scan(&s.id, &s.ftype, &s.desc, &date, &s.content)
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
	stmt, err := tx.Prepare(fmt.Sprintf("update snippet set type=?, desc=?, date=?, content=? where id=%d", id))
	if err != nil {
        return err
	}
	defer stmt.Close()
    _, err = stmt.Exec(s.ftype, s.desc, s.date.Format(time.RFC3339), s.content)
    if err != nil {
        return err
    }
	tx.Commit()
    return nil
}

func ListSnippets(offset, count int64) ([]*Snippet, error) {
    sql := fmt.Sprintf("select id, type, desc, date  from snippet order by id desc limit %d offset %d",
                    count,
                    offset)
    rows, err := db.Query(sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var ret []*Snippet
    for rows.Next() {
        s := &Snippet{}
        var date string
        rows.Scan(&s.id, &s.ftype, &s.desc, &date)
        s.date, _ = time.Parse(time.RFC3339, date)
        ret = append(ret, s)
    }
    return ret, nil
}

