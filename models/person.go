package models

import db "ginDemo/database"

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

/*
插入数据
*/
func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersonById(id string) (person Person, err error) {
	stmt, err := db.SqlDB.Prepare("select id, first_name, last_name from person where id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
