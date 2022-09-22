package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	Subjects []string
}

func NewStudent(rollno int, name string, address string, Subjects[] string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.Subjects= Subjects



	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, Subjects[] string) *Student {
	st := NewStudent(rollno, name, address, Subjects)
	ls.list = append(ls.list, st)
	return st
}

func (ls *StudentList) Print() {
	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("Student Rollno: %d\n", ls.list[i].rollnumber)
		fmt.Printf("Student Name: %s\n", ls.list[i].name)
		fmt.Printf("Student Address: %s\n\n", ls.list[i].address)
		fmt.Printf("Student Subjects: %s\n\n", ls.list[i].Subjects)

	}
}

func main(){
	
	student := new(StudentList)
	stuuudent:=[] string {"History","BLockchain"}
	stuuudent1:=[] string {"TBW","BLockchain"}
	student.CreateStudent(2000, "Azaz", "Khushab",stuuudent)
	student.CreateStudent(2192, "Zaka", "Jannat-Ul-Firdoos",stuuudent1)
	student.Print()
	for i:=0;i<len(student.list);i++{
	Hashh:=strconv.Itoa(student.list[i].rollnumber)+ student.list[i].name+student.list[i].address 
	
	for x:=0;x<len((student.list[i].Subjects));x++{
		Hashh=Hashh+student.list[i].Subjects[x]
	}
	sum := sha256.Sum256([]byte(Hashh))

	fmt.Printf("%x\n", sum) //hexadecimal
}
}




