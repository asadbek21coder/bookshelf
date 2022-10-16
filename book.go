package bookshelf

type BookDB struct {
	Id      int    `json:"id" db:"id"`
	Isbn    string `json:"isbn" db:"isbn"`
	Status  int    `json:"status" db:"status"`
	UserKey string `json:"user_key" db:"user_key"`
}

type Isbn struct {
	Isbn string `json:"isbn" binding:"required"`
}

type BooksDB struct {
	Books []BookDB `json:"books"`
}

// type AllBooksResponse struct {
// 	data []FullBook
// }

type FullB struct {
	Id        int    `json:"id"`
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
	Pages     int    `json:"pages"`
}

type FullBook struct {
	Book   FullB `json:"book"`
	Status int   `json:"status"`
}
