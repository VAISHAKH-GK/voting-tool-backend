package handlers

import (
	"fmt"

	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/sessions"
	"github.com/VAISHAKH-GK/voting-tool-backend/pkg/user"
	"github.com/VAISHAKH-GK/voting-tool-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	var user user.User
	var err = c.BodyParser(&user)
	if err != nil {
		fmt.Println(err)
	}

	var otp = utils.GenerateOTP()

	sessions.Add(c, "signup-username", user.Username)
	sessions.Add(c, "signup-email", user.Email)
	sessions.Add(c, "signup-password", user.Password)
	sessions.Add(c, "signup-otp", otp)

	utils.SendMail(user.Email, fmt.Sprintf("<h3>This is the OTP for your voting tool website account creation </h3><br><h1> %d </h1>", otp))

	return c.JSON("SEND OTP")
}

func SubmitOTP(c *fiber.Ctx) error {
	var otp = sessions.Get(c, "signup-otp")
	fmt.Println(otp)

	var data map[string]int
	c.BodyParser(&data)

	if otp.(int) != data["otp"] {
		return c.Status(401).JSON("Wrong OTP")
	}

	var username = sessions.Get(c, "signup-username").(string)
	var email = sessions.Get(c, "signup-email").(string)
	var password = sessions.Get(c, "signup-password").(string)

	password = utils.BcryptGenerateHash(password)

	var user = user.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	user.InsertUser()

	return c.JSON("Signup Success")
}

func Login(c *fiber.Ctx) error {
	var u map[string]interface{}
	var err = c.BodyParser(&u)
	if err != nil {
		fmt.Println(err)
	}

	var user = user.GetUser(u["username"].(string))

	if !utils.BcryptCompare(u["password"].(string), user.Password) {
		return c.Status(401).JSON("Wrong Password")
	}

	sessions.Add(c, "username", u["username"].(string))
	sessions.Add(c, "isLoggedIn", true)

	return c.JSON("Login Success")
}
