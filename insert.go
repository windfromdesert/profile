package main

import (
    "os"
    "strings"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

type Rec struct {
    Id_         bson.ObjectId       `bson:"_id"`
    Riqi        string              `bson:"riqi"`
    Shijian     string
    Tag         string
    Title       string
    Beizhu      string
}

func main() {
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()
    c := session.DB("profile").C("profile")
    mdfile := "md"
    fin, err := os.Open(mdfile)
    defer fin.Close()
    if err != nil {
        panic(err)
    }
    buf := make([]byte,1024)
    text := ""
    for {
        n, _ := fin.Read(buf)
        if 0 == n {  break }
        text = text + string(buf[:n])
    }
    text2 := strings.Split(text,"\r\n\r\n")
    text3 := strings.Split(text2[0],"\r\n")
    sj := ""
    for k, v := range text2 {
        switch k {
            case 0:
            case len(text2)-1:
                sj = sj + v
            default:
                sj = sj + v + "\n\n"
        }
    }
    er := c.Insert(&Rec{
        Id_:        bson.NewObjectId(),
        Riqi:       text3[0],
        Shijian:    sj,
        Tag:        text3[2],
        Title:      text3[1],
        Beizhu:     text3[3],
    })
    if er != nil {
        panic(er)
    }
}
