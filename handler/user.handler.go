package handler

import (
	"log"
	"mini-project/database"
	"mini-project/model/entity"
	"mini-project/model/request"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User{
		Nama: user.Nama,
		KataSandi: user.KataSandi,
		Telepon: user.Telepon,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		Tentang: user.Tentang,
		Pekerjaan: user.Pekerjaan,
		Email: user.Email,
		IDProvinsi: user.IDProvinsi,
		IDKota: user.IDKota,
		IsAdmin: user.IsAdmin,
	}

	errCreateUser := database.DB.Save(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data.",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data": newUser,
	})


}