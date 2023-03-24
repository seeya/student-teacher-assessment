package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/seeya/student-teacher-assessment/app/queries"
	"github.com/seeya/student-teacher-assessment/platform/database"
)

func main() {
	_ = godotenv.Load(".env")

	db, err := database.OpenDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	ApiQuery := queries.ApiQuery{DB: db}
	ApiQuery.SeedTeachers()
	ApiQuery.SeedStudents()

	_ = ApiQuery.TeacherAddStudent("teacherken@gmail.com",
		[]string{"studentjon@gmail.com",
			"studentbob@gmail.com", "studentmiche@gmail.com"})

	_ = ApiQuery.TeacherAddStudent("teacherjoe@gmail.com",
		[]string{"studentjon@gmail.com",
			"studentagnes@gmail.com", "studenthon@gmail.com"})

	common, err := ApiQuery.FindCommonStudents([]string{"teacherjoe@gmail.com", "teacherken@gmail.com"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Common Students: %v", common)

	ApiQuery.SuspendStudent("studentmiche@gmail.com")
	students, err := ApiQuery.StudentCanReceiveNotifications("teacherhc@gmail.com", "Hello students! @studentagnes@gmail.com @studentmiche@gmail.com @hehe@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Students: %v", students)

	log.Println("Hello world")
}
