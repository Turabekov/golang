package main

import (
	"errors"
	"fmt"
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
	GetArticleById(id string) Article
	GetAllArticle(limit, offset int, search, auhtor_id string) error
	UpdateArticle(article Article) error
	DeleteArticle(article Article) error
}

type AuthorCRUD interface {
	CreateAuthor() error
	GetAuthorById(id string) Author
	GetAllAuthor(limit, offset int, search, auhtor_id string) error
	UpdateAuthor(article Article) error
	DeleteAuthor(article Article) error
}

type storyDatabase struct {
	articles []Article
	authors  []Author
}

// Author CRUD ==============================================================================================================================
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

// Article CRUD ==============================================================================================================================
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

// Article CRUD ==============================================================================================================================
func main() {
	author1 := Author{
		id:   "qw23dasdg34fds",
		name: "Rashford",
		age:  23,
	}
	author2 := Author{
		id:   "xcvxclef324342",
		name: "Rooney",
		age:  38,
	}

	article1 := Article{
		id:        "24234sddfsf",
		title:     "Derby in EPL",
		desc:      "We want to describe",
		author_id: "qw23dasdg34fds",
	}

	database := storyDatabase{}
	// create authors
	database.CreateAuthor(author1)
	database.CreateAuthor(author2)

	// create articels
	database.CreateArticle(article1)

	fmt.Println(database)

}
