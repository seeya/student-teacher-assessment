package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/seeya/student-teacher-assessment/app/queries"
)

type RegisterStudentRequest struct {
	Email   string   `json:"email"`
	Student []string `json:"student"`
}

type CommonStudentRequest struct {
	Teachers []string `query:"teacher,required"`
}

type SuspendStudentRequest struct {
	Student string `json:"student"`
}

type RetrieveForNotificationsRequest struct {
	Teacher      string `json:"teacher"`
	Notification string `json:"notification"`
}

func errorMiddleware(c *fiber.Ctx, err error) error {
	return c.Status(500).JSON(fiber.Map{
		"message": err.Error(),
	})

}

func InitApiRoutes(router fiber.Router, api *queries.ApiQuery) {
	router.Get("/seed", func(c *fiber.Ctx) error {
		api.SeedTeachers()
		api.SeedStudents()
		api.SeedTeacherStudentRelationship()

		return c.SendString("Seeded")
	})

	router.Get("/commonstudents", func(c *fiber.Ctx) error {
		req := new(CommonStudentRequest)

		if err := c.QueryParser(req); err != nil {
			return errorMiddleware(c, err)
		}

		log.Println(req)

		students, err := api.FindCommonStudents(req.Teachers)

		if err != nil {
			return errorMiddleware(c, err)
		}

		return c.JSON(students)
	})

	router.Post("/retrievefornotifications", func(c *fiber.Ctx) error {
		req := new(RetrieveForNotificationsRequest)

		if err := c.BodyParser(req); err != nil {
			return errorMiddleware(c, err)
		}

		students, err := api.StudentCanReceiveNotifications(req.Teacher, req.Notification)

		if err != nil {
			return errorMiddleware(c, err)
		}

		return c.JSON(students)
	})

	router.Post("/register", func(c *fiber.Ctx) error {
		req := new(RegisterStudentRequest)

		// json.Unmarshal([]byte(c.Body()), &req)

		// TODO: Why cannot parse properly.
		if err := c.BodyParser(req); err != nil {
			return errorMiddleware(c, err)

		}

		log.Printf("Request: %v", req)

		err := api.TeacherAddStudent(req.Email, req.Student)

		if err != nil {
			return errorMiddleware(c, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

	router.Post("/suspend", func(c *fiber.Ctx) error {
		req := new(SuspendStudentRequest)

		if err := c.BodyParser(req); err != nil {
			return errorMiddleware(c, err)
		}

		err := api.SuspendStudent(req.Student)

		if err != nil {
			return errorMiddleware(c, err)
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

}
