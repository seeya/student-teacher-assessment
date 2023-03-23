package queries

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/seeya/student-teacher-assessment/app/models"
)

type ApiQuery struct {
	*sqlx.DB
}

func (q *ApiQuery) GetTeachers() ([]string, error) {
	var teachers []string
	err := q.Select(&teachers, "SELECT email FROM teachers")
	return teachers, err
}

func (q *ApiQuery) FindCommonStudents(teacherEmails []string) error {
	// Find all IDs of teachers
	var teacherIDs []string
	for _, email := range teacherEmails {
		t, err := q.FindTeacherIDByEmail(email)

		if err != nil {
			return err
		}

		teacherIDs = append(teacherIDs, strconv.Itoa(int(t.ID)))
	}

	// TODO: Test out this query
	query := `SELECT student_id, COUNT(*) as occurrences
				FROM teachings
				GROUP BY student_id
				HAVING occurrences = ?;`

	log.Printf("Teacher IDs: %v", teacherIDs)
	return nil
}

// func (q *ApiQuery) FindCommonStudents(teacherEmails []string) error {
// 	// Find all IDs of teachers
// 	query := `SELECT id FROM teachers
//     			WHERE email IN (` + strings.TrimSuffix(strings.Repeat("?,", len(teacherEmails)), ",") + `)`

// 	stmt, err := q.DB.Preparex(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	args := make([]interface{}, len(teacherEmails))

// 	for i, email := range teacherEmails {
// 		args[i] = email
// 	}

// 	rows, err := stmt.Queryx(args...)

// 	if err != nil {
// 		return err
// 	}

// 	defer rows.Close()

// 	var teacherIDs []string

// 	for rows.Next() {
// 		var teacherID int
// 		err := rows.Scan(&teacherID)

// 		if err != nil {
// 			return err
// 		}

// 		teacherIDs = append(teacherIDs, strconv.Itoa(teacherID))
// 	}

// 	log.Printf("Teacher IDs: %v", teacherIDs)

// 	return err
// }

func (q *ApiQuery) TeacherAddStudent(teacherEmail string, studentsEmail []string) error {
	s := `INSERT INTO teachings (student_id, teacher_id) VALUES (?, ?)`

	teacher, err := q.FindTeacherIDByEmail(teacherEmail)

	if err != nil {
		return err
	}

	for _, email := range studentsEmail {
		student, err := q.FindStudentIDByEmail(email)

		if err != nil {
			return err
		}

		_, err = q.DB.Exec(s, student.ID, teacher.ID)

		if err != nil {
			return err
		}

		log.Printf("Student: %v", student)
	}

	return nil
}

func (q *ApiQuery) FindTeacherIDByEmail(email string) (*models.Teacher, error) {
	s := `SELECT * FROM teachers WHERE email = ? LIMIT 1`

	var teacher = models.Teacher{}
	err := q.DB.Get(&teacher, s, email)

	if err != nil {
		return nil, err
	}

	log.Printf("Teacher: %v", teacher)

	return &teacher, err
}

func (q *ApiQuery) FindStudentIDByEmail(email string) (*models.Student, error) {
	s := `SELECT * FROM students WHERE email = ? LIMIT 1`

	var student = models.Student{}
	err := q.DB.Get(&student, s, email)

	if err != nil {
		return nil, err
	}

	log.Printf("Student: %v", student)

	return &student, err

}

func (q *ApiQuery) SeedTeachers() error {
	var teachers = []string{"ken", "joe"}

	s := `INSERT INTO teachers (email) VALUES (?)`

	for _, name := range teachers {
		_, err := q.DB.Exec(s, fmt.Sprintf("teacher%s@gmail.com", name))

		if err != nil {
			log.Printf("Failed to seed Teachers table: %v\n", err)
			return err
		}

		log.Printf("Seeded Teachers table")
	}

	return nil
}

func (q *ApiQuery) SeedStudents() error {
	var students = []string{"agnes", "bob", "miche", "mary", "jon", "hon"}

	s := `INSERT INTO students(email) VALUES (?)`

	for _, name := range students {
		_, err := q.DB.Exec(s, fmt.Sprintf("student%s@gmail.com", name))

		if err != nil {
			log.Printf("Failed to seed Students table: %v\n", err)
			return err
		}

		log.Printf("Seeded Students table")
	}

	return nil
}
