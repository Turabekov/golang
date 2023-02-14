package main

import (
	"errors"
	"fmt"
	"strings"
)

type Article struct {
	id        string
	title     string
	desc      string
	author_id string
}

type Author struct {
	id   string
	name string
	age  int
}

type ArticleCRUD interface {
	CreateArticle(article Article) error
	GetArticleById(id string) (Article, Author, error)
	GetAllArticle(offset, limit int, search, auhtor_id string) ([]Article, error)
	UpdateArticle(article Article) error
	DeleteArticle(article Article) error
}

type AuthorCRUD interface {
	CreateAuthor() error
	GetAuthorById(id string) (Author, error)
	GetAllAuthor(limit, offset int, search string) ([]Author, error)
	UpdateAuthor(article Article) (Author error)
	DeleteAuthor(article Article) error
}

type storyDatabase struct {
	articles []Article
	authors  []Author
}

// Author CRUD ==============================================================================================================================
// Create Author
func (s *storyDatabase) CreateAuthor(author Author) error {
	isContinue := false
	// Check exsiting of author
	for _, val := range s.authors {
		if val.id == author.id {
			isContinue = true
			break
		} else {
			isContinue = false
		}
	}
	if isContinue {
		// return error if author exist
		return errors.New("Article with this info already exist")
	}
	// add new author to DB
	s.authors = append(s.authors, author)

	return nil
}

// Get Author by ID
func (s *storyDatabase) GetAuthorById(id string) (Author, []Article, error) {
	author := Author{}
	for _, v := range s.authors {
		if v.id == id {
			author = v
		}
	}

	authorArticles := []Article{}
	for _, v := range s.articles {
		if v.author_id == id {
			authorArticles = append(authorArticles, v)
		}
	}

	return author, authorArticles, nil
}

// Update Author
func (s *storyDatabase) UpdateAuthor(author Author) (Author, error) {
	for i, v := range s.authors {
		if v.id == author.id {
			s.authors[i].name = "Updated name"
			s.authors[i].age = 77
			return s.authors[i], nil
		}
	}

	return Author{}, errors.New("user didn't found for updating")
}

// Delete Author
func (s *storyDatabase) DeleteAuthor(author Author) error {
	for i, v := range s.authors {
		if v.id == author.id {
			s.authors = append(s.authors[:i], s.authors[i+1:]...)
			return nil
		}
	}

	return errors.New("user didn't found for deleting")
}

// Get All authors
func (s *storyDatabase) GetAllAuthor(offset, limit int, search string) ([]Author, error) {

	// // if searched not given or empty
	if len(search) <= 0 {
		if limit+offset > len(s.authors) {
			fmt.Println("Hello")
			if offset > len(s.authors) {
				return []Author{}, nil
			}
			fmt.Println("Hello")
			return s.authors[offset:], nil
		}

		return s.authors[offset : limit+offset], nil
	}
	// // Search
	res := []Author{}
	for _, v := range s.authors {
		fullName := strings.ToLower(v.name)
		if strings.Contains(fullName, strings.ToLower(strings.Replace(search, " ", "", -1))) {
			res = append(res, v)
		}
	}

	// if user not found
	if len(res) <= 0 {
		return []Author{}, nil
	}

	// if offset and limit out of range
	if limit+offset > len(res) {
		// offset more than length of res
		if offset > len(res) {
			return []Author{}, nil
		}

		return res[offset:], nil
	}

	return res[offset : limit+offset], nil
}

// Article CRUD ==============================================================================================================================

// Create Article
func (s *storyDatabase) CreateArticle(article Article) error {
	isContinue := false
	// Check exsiting of article
	for _, val := range s.articles {
		if val.id == article.id {
			isContinue = true
			break
		} else {
			isContinue = false
		}
	}
	if isContinue {
		// return error if article exist
		return errors.New("Article with this info already exist")
	}
	// add new article to DB
	s.articles = append(s.articles, article)
	return nil
}

// Get article by id
func (s *storyDatabase) GetArticleById(id string) (Article, Author, error) {
	var article Article
	for _, v := range s.articles {
		if v.id == id {
			article = v
		}
	}

	for _, v := range s.authors {
		if v.id == article.author_id {
			return article, v, nil
		}
	}

	return Article{}, Author{}, errors.New("Not found")
}

func (s *storyDatabase) GetAllArticle(offset, limit int, search, auhtor_id string) ([]Article, error) {
	if len(auhtor_id) > 0 {
		authorArticles := []Article{}
		for _, v := range s.articles {
			if v.author_id == auhtor_id {
				authorArticles = append(authorArticles, v)
			}
		}

		articles, err := searchFunc(offset, limit, search, authorArticles)

		if err != nil {
			return []Article{}, err
		}

		return articles, nil

	}

	articles, err := searchFunc(offset, limit, search, s.articles)

	if err != nil {
		return []Article{}, err
	}

	return articles, nil

}

func searchFunc(offset, limit int, search string, arr []Article) ([]Article, error) {
	// // if searched not given or empty
	if len(search) <= 0 {
		if limit+offset > len(arr) {
			if offset > len(arr) {
				return []Article{}, nil
			}
			return arr[offset:], nil
		}

		return arr[offset : limit+offset], nil
	}
	// // Search
	res := []Article{}
	for _, v := range arr {
		fullName := strings.ToLower(v.title)
		if strings.Contains(fullName, strings.ToLower(strings.Replace(search, " ", "", -1))) {
			res = append(res, v)
		}
	}

	// if user not found
	if len(res) <= 0 {
		return []Article{}, nil
	}

	// if offset and limit out of range
	if limit+offset > len(res) {
		// offset more than length of res
		if offset > len(res) {
			return []Article{}, nil
		}

		return res[offset:], nil
	}

	return res[offset : limit+offset], nil
}

// Update Article
func (s *storyDatabase) UpdateArticle(article Article) (Article, error) {
	for i, v := range s.articles {
		if v.id == article.id {
			s.articles[i].title = "Updated title"
			s.articles[i].desc = "updated descr"
			return s.articles[i], nil
		}
	}

	return Article{}, errors.New("user didn't found for updating")
}

// Delete Author
func (s *storyDatabase) DeleteArticle(article Article) error {
	for i, v := range s.articles {
		if v.id == article.id {
			s.articles = append(s.articles[:i], s.articles[i+1:]...)
			return nil
		}
	}

	return errors.New("user didn't found for deleting")
}

// CRUD ==============================================================================================================================
func main() {
	database := storyDatabase{
		authors: []Author{
			{
				id:   "qw23dasdg34fds",
				name: "Rashford",
				age:  23,
			},
			{
				id:   "xcvxclef324342",
				name: "Rooney",
				age:  38,
			},
			{
				id:   "asdcxvdfg",
				name: "Fabinio",
				age:  43,
			},
			{
				id:   "asdcxvdfg",
				name: "Ronaldo",
				age:  38,
			},
		},
		articles: []Article{
			{
				id:        "24234sddfsf",
				title:     "Derby in EPL Rooney",
				desc:      "We want to describe",
				author_id: "xcvxclef324342",
			},
			{
				id:        "24ww4sddfsf",
				title:     "Nasa cars Rooney",
				desc:      "Lorem ipsum sdfds",
				author_id: "xcvxclef324342",
			},
			{
				id:        "24ww4sdssdfsf",
				title:     "Aaa bgdbg Rooney",
				desc:      "Lorem ipsum sdfds",
				author_id: "xcvxclef324342",
			},
			{
				id:        "24wws4sdssdfsf",
				title:     "Rashfords article bgdbg",
				desc:      "Lorem ipsum sdfds",
				author_id: "qw23dasdg34fds",
			},
		},
	}
	// =========CRUD AUTHOR=============================================================================
	// CreateAuthor
	// err := database.CreateAuthor(Author{
	// 	id:   "djfhsvbndjskfh",
	// 	name: "Khumoyun",
	// 	age:  20,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 		return
	// } else {
	// 	fmt.Println("After created", database.authors)
	// }

	// getAuthor by Id
	// author, authorArticles, err := database.GetAuthorById("xcvxclef324342")

	// if err != nil {
	// 	fmt.Println(err)
	// return
	// return
	// } else {
	// 	fmt.Println("Get by id", author, "articles", authorArticles)
	// }

	// update author
	// updatedUser, err := database.UpdateAuthor(database.authors[0])
	// if err != nil {
	// 	fmt.Println(err)
	//  return
	// } else {
	// 	fmt.Println("Updated User:", updatedUser)
	// 	fmt.Println(database.authors)
	// }

	// Delete Author
	// err := database.DeleteAuthor(database.authors[0])
	// if err != nil {
	// 	fmt.Println(err)
	// return
	// } else {
	// 	fmt.Println(database.authors)
	// }

	// Get all authors
	// users, err := database.GetAllAuthor(1, 1, "Ro")
	// if err != nil {
	// 	fmt.Println(err)
	// return
	// } else {
	// 	fmt.Println(users)
	// }

	// =========CRUD ARTICLE=============================================================================
	// CreateArticle
	database.CreateArticle(Article{
		id:        "sdfsdf34324",
		title:     "BBC news",
		desc:      "Hello Sam",
		author_id: "asdcxvdfg",
	})

	fmt.Println("Creating:", database.articles)
	// Get by id
	article, author, err := database.GetArticleById("sdfsdf34324")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Get by id author:", author)
		fmt.Println("Get by id article:", article)
	}
	// Update article
	updatedArticle, err := database.UpdateArticle(database.articles[1])
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("After updated", updatedArticle)
		fmt.Println(database.articles)
	}

	// Delete Article
	err = database.DeleteArticle(database.articles[0])
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("After deleted", database.articles)
	}

	// Get all article
	articles, err := database.GetAllArticle(0, 100, "a", "qw23dasdg34fds")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Get all", articles)
	}

}
