package model

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func NewArticle() *Article {
	a := new(Article)
	a.Title = "New title"
	a.Desc = "New Desc"
	a.Content = "New Content"
	return a
}

func ArticlesArray() *[]Article {

	var array = make([]Article, 2, 3)

	array[0] = Article{
		Id:      1,
		Title:   "Top",
		Desc:    "Top",
		Content: "Top",
	}
	array[1] = Article{
		Id:      2,
		Title:   "Bot",
		Desc:    "Bot",
		Content: "Bot",
	}

	return &array
}
