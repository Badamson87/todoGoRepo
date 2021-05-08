package todo
import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Getter interface {GetAll() []Item}

type Adder interface {Add(item Item)}

type Deleter interface {Delete(items []Item)}

type Updater interface {Update(item Item)}

type Item struct {
    Id int `json:"id"`
    Checked bool  `json:"checked"`
    Title string  `json:"title"`
}

func New() *Repo {
    return &Repo{
        Items: []Item{},
    }
}

type Repo struct {
    Items []Item
}

type Todo struct {
    Id int `json:"id"`
    Checked bool `json:"checked"`
    Title string `json:"title"`
}

func (r *Repo) Add(item Item) {
    r.Items = append(r.Items, item)
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

func (r Repo) Delete(items []Item) {
//     for index, element := range r {
//         for i, item := range items {
//             if element.Id == item.Id {
//                 RemoveIndex(r, index)
//             }
//         }
//     }
//     return r.Items
}

func (r Repo) Update (item Item) {

}
