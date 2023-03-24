package queries

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/seeya/student-teacher-assessment/app/models"
	"golang.org/x/exp/maps"
)

type ApiQuery struct {
	*sqlx.DB
}

func (q *ApiQuery) GetTeachers() ([]string, error) {
	var teachers []string
	err := q.Select(&teachers, "SELECT email FROM teachers")
	return teachers, err
}

func (q *ApiQuery) FindCommonStudents(teacherEmails []string) ([]string, error) {
	// Find all IDs of teachers
	var teacherIDs []string
	for _, email := range teacherEmails {
		t, err := q.FindTeacherIDByEmail(email)

		if err != nil {
			return nil, err
		}

		teacherIDs = append(teacherIDs, strconv.Itoa(int(t.ID)))
	}

	// 1. Select all rows where the teacher_id is in the list of teacherIDs
	// 2. Group by student_id and count the total occurance
	// 3. Compare the total occurance to the number of teacherIDs we input using HAVING
	// 4. Left join to get the email of the student
	query := `SELECT student_id, COUNT(*) as occurrences, email
				FROM teachings
				LEFT JOIN students ON teachings.student_id = students.id
				WHERE teacher_id IN (` + strings.TrimSuffix(strings.Repeat("?,", len(teacherIDs)), ",") + `)
				GROUP BY student_id
				HAVING occurrences = ?;`

	stmt, err := q.DB.Preparex(query)

	if err != nil {
		return nil, err
	}

	args := make([]interface{}, len(teacherIDs)+1)
	for i, id := range teacherIDs {
		args[i] = id
	}
	args[len(args)-1] = len(teacherEmails)

	log.Printf("Total Teachers Length: %v", len(teacherEmails))
	rows, err := stmt.Queryx(args...)

	if err != nil {
		return nil, err
	}

	var students []string

	for rows.Next() {
		var studentID int
		var count int
		var email string
		err := rows.Scan(&studentID, &count, &email)

		if err != nil {
			return nil, err
		}

		students = append(students, email)

		log.Printf("Common Student: %v, %v, %v", studentID, count, email)
	}

	log.Printf("Teacher IDs: %v", teacherIDs)
	return students, nil
}

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

func (q *ApiQuery) SuspendStudent(email string) error {
	s := `UPDATE students SET is_suspended = true WHERE email = ?`

	stmt, err := q.DB.Preparex(s)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(email)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Result: %v", affected)
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

func (q *ApiQuery) StudentCanReceiveNotifications(teacherEmail string, notification string) ([]string, error) {
	teacher, err := q.FindTeacherIDByEmail(teacherEmail)

	if err != nil {
		return nil, err
	}

	s := `SELECT email FROM teachings 
			LEFT JOIN students ON students.id = student_id 
			WHERE teacher_id = ? 
			AND students.is_suspended = FALSE;`

	rows, err := q.DB.Queryx(s, teacher.ID)

	if err != nil {
		return nil, nil
	}

	var students = map[string]int{}

	var email string
	for rows.Next() {
		_ = rows.Scan(&email)

		students[email] = 1
	}

	// Parse the notification for @email
	pattern := "@[\\d\\w]{1,}@[\\d\\w]{1,}.[\\d\\w]{1,}"
	re := regexp.MustCompile(pattern)
	emails := re.FindAllString(notification, -1)
	log.Printf("Match: %v", emails)

	// Add the emails found into our dictionary to remove duplicates
	for _, email := range emails {
		// Remove the first @ at the front
		students[email[1:]] = 1
	}

	return maps.Keys(students), nil
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
