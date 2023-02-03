package main

import "fmt"

type User struct {
	Id                  uint32
	FirstName, LastName string
}

type Option struct {
	Id        uint8
	Content   string
	IsCorrect bool
}

type Question struct {
	Id      uint32
	Content string
	Options [4]Option
}

type QuestionMap map[uint32]Question
type UserMap map[uint32]User

func ConvertToMapUser(users []User) (userMap UserMap) {
	userMap = UserMap{}

	for _, user := range users {
		userMap[user.Id] = user
	}

	return
}

func ConvertToMapQuestion(questions []Question) (questionMap QuestionMap) {
	questionMap = QuestionMap{}

	for _, question := range questions {
		questionMap[question.Id] = question
	}

	return
}

func (q Question) getOption(optionId uint8) Option {

	for _, option := range q.Options {
		if option.Id == optionId {
			return option
		}
	}

	return Option{}
}

func main() {

	var users = []User{
		{
			Id:        1,
			FirstName: "Shoxrux",
			LastName:  "Safarov",
		},
		{
			Id:        2,
			FirstName: "Humoyun",
			LastName:  "Turabekov",
		},
	}

	var questions = []Question{
		{
			Id:      1,
			Content: "Apple qanday meva?",
			Options: [4]Option{
				{
					Id:        1,
					Content:   "Olma",
					IsCorrect: true,
				},
				{
					Id:        2,
					Content:   "Banan",
					IsCorrect: false,
				},
				{
					Id:        3,
					Content:   "Nok",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Anor",
					IsCorrect: false,
				},
			},
		},
		{
			Id:      2,
			Content: "Golang qanday til?",
			Options: [4]Option{
				{
					Id:        1,
					Content:   "Mol tili",
					IsCorrect: false,
				},
				{
					Id:        2,
					Content:   "Dasturlash tili",
					IsCorrect: true,
				},
				{
					Id:        3,
					Content:   "Odam tili",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Til emas",
					IsCorrect: false,
				},
			},
		},
		{
			Id:      3,
			Content: "APL komandasini tanlang?",
			Options: [4]Option{
				{
					Id:        1,
					Content:   "Manchester United",
					IsCorrect: true,
				},
				{
					Id:        2,
					Content:   "Real Madrid",
					IsCorrect: false,
				},
				{
					Id:        3,
					Content:   "PSG",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Al Nasr",
					IsCorrect: false,
				},
			},
		},
		{
			Id:      4,
			Content: "C.Ronaldo hozirda qaysi komandada?",
			Options: [4]Option{
				{
					Id:        1,
					Content:   "Manchester United",
					IsCorrect: false,
				},
				{
					Id:        2,
					Content:   "Real Madrid",
					IsCorrect: false,
				},
				{
					Id:        3,
					Content:   "PSG",
					IsCorrect: false,
				},
				{
					Id:        4,
					Content:   "Al Nasr",
					IsCorrect: true,
				},
			},
		},
		{
			Id:      5,
			Content: "Bil Geyts kompaniyasining nomini tanlang?",
			Options: [4]Option{
				{
					Id:        1,
					Content:   "Alibaba",
					IsCorrect: false,
				},
				{
					Id:        2,
					Content:   "Apple",
					IsCorrect: false,
				},
				{
					Id:        3,
					Content:   "Microsoft",
					IsCorrect: true,
				},
				{
					Id:        4,
					Content:   "Neftlix",
					IsCorrect: true,
				},
			},
		},
	}

	userMap := ConvertToMapUser(users)
	questionMap := ConvertToMapQuestion(questions)
	fmt.Println(userMap[1].FirstName + " " + userMap[1].LastName)

	answer := map[uint32]uint8{}
	ball := 0
	// answer to questions
	answer[1] = 1
	answer[2] = 3
	answer[3] = 1
	answer[4] = 4
	answer[5] = 3

	for questionId, optionId := range answer {

		fmt.Printf("%d. %s\n", questionMap[questionId].Id, questionMap[questionId].Content)

		option := questionMap[questionId].getOption(optionId)

		fmt.Printf("\t[%c]. %s\n", 64+option.Id, option.Content)
		if option.IsCorrect {
			ball += 10
			fmt.Printf("\t[CORRECT]\n")
		} else {
			ball -= 5
			fmt.Printf("\t[WRONG]\n")
		}

	}
	fmt.Println("Total ball:", ball)
}
