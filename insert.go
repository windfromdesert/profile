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
    var recs []Rec
    c := session.DB("profile").C("profile")
    c.Find(nil).All(&recs)
    list := ""
    for _, v := range recs {
        r := strings.Replace(v.Riqi,"\n","",-1)
        t := strings.Replace(v.Tag,"\n","",-1)
        list = list + r + "-" + t + "-" + v.Title
    }
    indexfile := "index"
    f, err := os.Create(indexfile)
    defer f.Close()
    if err != nil {
        panic(err)
    }
    f.WriteString(list)
}
