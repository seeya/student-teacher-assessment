package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/seeya/student-teacher-assessment/app/queries"
)

type RegisterStudentRequest struct {
	Email   string   `json:"email"`
	Student []string `json:"student[]"`
}

type CommonStudentRequest struct {
	Teachers []string `query:"teacher"`
}

type SuspendStudentRequest struct {
	Student string `json:"student"`
}

type RetrieveForNotificationsRequest struct {
	Teacher      string `json:"teacher"`
	Notification string `json:"notification"`
}

func InitApiRoutes(router fiber.Router, api *queries.ApiQuery) {

	router.Get("/commonstudents", func(c *fiber.Ctx) error {
		req := new(CommonStudentRequest)

		if err := c.QueryParser(req); err != nil {
			return err
		}

		log.Printf("Request: %v", req.Teachers)
		students, err := api.FindCommonStudents(req.Teachers)

		if err != nil {
			return err
		}

		return c.JSON(students)
	})

	router.Post("/retrievefornotifications", func(c *fiber.Ctx) error {
		req := new(RetrieveForNotificationsRequest)

		if err := c.BodyParser(req); err != nil {
			return err
		}

		students, err := api.StudentCanReceiveNotifications(req.Teacher, req.Notification)

		if err != nil {
			return err
		}

		return c.JSON(students)
	})

	router.Post("/register", func(c *fiber.Ctx) error {
		req := new(RegisterStudentRequest)

		if err := c.BodyParser(req); err != nil {
			return err
		}

		log.Printf("Request: %v", req)

		return c.SendStatus(fiber.StatusNoContent)
	})

	router.Post("/suspend", func(c *fiber.Ctx) error {
		req := new(SuspendStudentRequest)

		if err := c.BodyParser(req); err != nil {
			return err
		}

		err := api.SuspendStudent(req.Student)

		if err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusNoContent)
	})

}
