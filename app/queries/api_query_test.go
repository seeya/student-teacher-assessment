package queries

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/seeya/student-teacher-assessment/platform/database"
	"github.com/stretchr/testify/require"
)

func TestGetTeachers(t *testing.T) {
	_ = godotenv.Load("../../.env")
	db, err := database.OpenTestDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	ApiQuery := &ApiQuery{DB: db}

	result, err := ApiQuery.GetTeachers()

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	expected := []string{"teacherjoe@gmail.com", "teacherken@gmail.com"}

	require.Equal(t, len(expected), len(result))
	require.ElementsMatch(t, expected, result)
}

func TestFindCommonStudents(t *testing.T) {
	_ = godotenv.Load("../../.env")
	db, err := database.OpenTestDBConnection()

	require.NoError(t, err)

	ApiQuery := &ApiQuery{DB: db}

	// Single Teacher
	teacherEmails := []string{"teacherken@gmail.com"}

	common, err := ApiQuery.FindCommonStudents(teacherEmails)

	require.NoError(t, err)
	require.NotEmpty(t, common)
	require.ElementsMatch(t, []string{"studentagnes@gmail.com", "studentbob@gmail.com", "studentmiche@gmail.com"}, common)

	// Multiple Teachers
	teacherEmails = []string{"teacherken@gmail.com", "teacherjoe@gmail.com"}

	common, err = ApiQuery.FindCommonStudents(teacherEmails)
	require.NoError(t, err)
	require.NotEmpty(t, common)
	require.ElementsMatch(t, []string{"studentagnes@gmail.com", "studentbob@gmail.com"}, common)

}

func TestSuspendStudent(t *testing.T) {
	_ = godotenv.Load("../../.env")
	db, err := database.OpenTestDBConnection()

	require.NoError(t, err)

	ApiQuery := &ApiQuery{DB: db}

	student, err := ApiQuery.FindStudentIDByEmail("studenthon@gmail.com")
	require.NoError(t, err)
	require.Equal(t, student.IsSuspended, false)

	err = ApiQuery.SuspendStudent("studenthon@gmail.com")

	require.NoError(t, err)

	student, err = ApiQuery.FindStudentIDByEmail("studenthon@gmail.com")

	require.NoError(t, err)
	require.Equal(t, student.IsSuspended, true)
}

// studenthon is suspended, so he should not be returned
func TestStudentCanReceiveNotifications(t *testing.T) {
	_ = godotenv.Load("../../.env")
	db, err := database.OpenTestDBConnection()

	require.NoError(t, err)

	ApiQuery := &ApiQuery{DB: db}

	// 1. teacherjoe when seeded, has agnes, bob & hon
	// 2. include studentmiche in message
	// 3. studenthon is suspended in the previous test so he should not be returned
	students, err := ApiQuery.StudentCanReceiveNotifications("teacherjoe@gmail.com", "hello @studentmiche@gmail.com")

	require.NoError(t, err)
	log.Printf("students: %v", students)
	require.ElementsMatch(t, students, []string{"studentagnes@gmail.com", "studentbob@gmail.com", "studentmiche@gmail.com"})
}

func TestTeacherAddStudent(t *testing.T) {
	_ = godotenv.Load("../../.env")
	db, err := database.OpenTestDBConnection()

	require.NoError(t, err)

	ApiQuery := &ApiQuery{DB: db}

	err = ApiQuery.TeacherAddStudent("teacherjoe@gmail.com", []string{"studentmary@gmail.com"})

	require.NoError(t, err)

	students, err := ApiQuery.FindCommonStudents([]string{"teacherjoe@gmail.com"})

	require.NoError(t, err)
	require.ElementsMatch(t, students,
		[]string{"studentagnes@gmail.com", "studentbob@gmail.com",
			"studentmary@gmail.com", "studenthon@gmail.com"})

}
