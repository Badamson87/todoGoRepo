package todo

type Getter interface {GetAll() []Item}

type Adder interface {Add(item Item)}

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

func (r *Repo) GetAll() []Item {
    return r.Items
}
