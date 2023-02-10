package main

import "fmt"

type Record struct {
	id    string
	name  string
	email string
}

type Database interface {
	Insert(record Record)
	Update(record Record)
	Delete(record Record)
	GetAll(limit, offset int) []Record
}

type InMemoryDB struct {
	records []Record
}

// Insert
func (inDB *InMemoryDB) Insert(record Record) {
	isContinue := false

	for _, val := range inDB.records {
		if val.id == record.id {
			isContinue = true
			break
		} else {
			isContinue = false
		}
	}

	if isContinue {
		fmt.Println("Not found for inserting")
		return
	}

	inDB.records = append(inDB.records, record)
}

// Update
func (inDB *InMemoryDB) Update(record Record) {
	isContinue := true

	for i, val := range inDB.records {
		if val.id == record.id {
			inDB.records[i].name = "Updated name"
			inDB.records[i].email = "updated email"
			break
		} else {
			isContinue = false
		}
	}

	if !isContinue {
		fmt.Println("Not found for updating")
		return
	}

}

// Delete
func (inDB *InMemoryDB) Delete(record Record) {
	isContinue := true

	for i, val := range inDB.records {
		if val.id == record.id {
			inDB.records = append(inDB.records[:i], inDB.records[i+1:]...)
			break
		} else {
			isContinue = false
		}
	}

	if !isContinue {
		fmt.Println("Not found for deleting")
		return
	}
}

// GetAll
func (inDB *InMemoryDB) GetAll(limit, offset int) []Record {
	newSlice := []Record{}
	count := 0

	for i := offset; i < len(inDB.records); i++ {
		if count < limit {
			newSlice = append(newSlice, inDB.records[i])
		}
		count++
	}

	return newSlice
}

// GetOnyById
func (inDB *InMemoryDB) GetById(record Record) Record {
	for _, val := range inDB.records {
		if val.id == record.id {
			return val
		}
	}
	return Record{}
}

func main() {
	obj1 := Record{
		id:    "23123weqw3234e",
		name:  "Name example",
		email: "example@gmail.com",
	}

	obj2 := Record{
		id:    "dsda23123weqw3234e",
		name:  "Name2 example",
		email: "example2@gmail.com",
	}
	obj3 := Record{
		id:    "asddsda23123weqw3234e",
		name:  "Name3 example",
		email: "example3@gmail.com",
	}
	obj4 := Record{
		id:    "a4sddsda23123weqw3234e",
		name:  "Name4 example",
		email: "example4@gmail.com",
	}

	inDB := InMemoryDB{}
	inDB.Insert(obj1)
	inDB.Insert(obj2)
	inDB.Insert(obj3)
	inDB.Insert(obj4)

	fmt.Println("Data:", inDB.records)
	inDB.Update(obj1)
	fmt.Println("After updated:", inDB.records)

	inDB.Delete(obj1)
	fmt.Println("After deleting:", inDB.records)

	fmt.Println("GetAll:", inDB.GetAll(0, 2))

}
