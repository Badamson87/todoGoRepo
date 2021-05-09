package todo
import (
    "fmt"
    "database/sql"
    "strings"
    _ "github.com/go-sql-driver/mysql"
)

type Item struct {
    Id int `json:"id"`
    Checked bool  `json:"checked"`
    Title string  `json:"title"`
}

type Todo struct {
    Id int `json:"id"`
    Checked int `json:"checked"`
    Title string `json:"title"`
}

func Add(item Item, db *sql.DB) int64 {
    var ins *sql.Stmt
    var err error
    ins,err = db.Prepare ("INSERT INTO `tododb`.`todo` (`checked`, `title`) VALUES (?, ?);")
     if err != nil {
        panic(err.Error())
     }
     defer ins.Close()
     res,err := ins.Exec(item.Checked, item.Title)
     if err != nil {
             panic(err.Error())
          }
      lid, err := res.LastInsertId()
      return lid
}

func Update(item Item, db *sql.DB) sql.Result {
    var ins *sql.Stmt
    var err error
    ins,err = db.Prepare ("UPDATE `tododb`.`todo` SET `checked` = ?, `title` = ? WHERE (`id` = ?);")
     if err != nil {
        panic(err.Error())
     }
     defer ins.Close()
     res,err := ins.Exec(item.Checked, item.Title, item.Id)
     if err != nil {
             panic(err.Error())
          }
      return res;
}

func Delete(ids string, db *sql.DB) bool {
        idArray := strings.Split(ids, ",")
        var ins *sql.Stmt
        var err error

        for i := 0;  i  < len(idArray); i++  {
           ins,err = db.Prepare ("DELETE FROM `tododb`.`todo` WHERE (`id` = ?);")
            if err != nil {
                panic(err.Error())
             }
             defer ins.Close()
             res,err := ins.Exec(idArray[i])
             if err != nil {
                     panic(err.Error())
                  }
                  fmt.Println(res)
              }
              return err == nil
}

func GetAll(db *sql.DB) []Item {
    var retVal []Item
    results,err := db.Query("SELECT * FROM todo")
    if err != nil {
            panic(err.Error())
        }
        defer results.Close()
        for results.Next() {
            var item Item
            err = results.Scan(&item.Id, &item.Checked, &item.Title)
            if err != nil {
               fmt.Println("unable to parse todo")
               panic(err.Error())
            }
            retVal = append(retVal, item)
        }
        return retVal;
}

