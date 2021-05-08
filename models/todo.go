package todo
import (
// "fmt"
"github.com/techinscribed/global-db/sqldb"
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

type Repo struct {
    Items []Item
}

func New() *Repo {
    return &Repo{
        Items: []Item{},
    }
}

func (r *Repo) Add(item Item) {
    r.Items = append(r.Items, item)
}

func (db DataBase) GetAll() []Item {
    results,err := db.Query("SELECT * FROM todo")
    if err != nil {
            panic(err.Error())
        }
        return results;
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
