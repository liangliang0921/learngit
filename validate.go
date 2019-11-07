package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

type Test struct {
	// Email string `validate:"email"`
	Size  int    `validate:"max=10,min=1"`
	Start int    `validate:"required"`
	End   int    `validate:"required,gtfield=Start"`
	Str   string `validate:"len=5"`
}

func test1() {

	// 优化错误的展示
	enu := en.New()
	uni := ut.New(enu, enu)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(&Test{9, 20, 30, "12345"})
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs.Translate(trans))
	}

}

type User struct {
	Name string `validate:"is-zhou"`
}

func (u *User) userValidator() error {
	validate := validator.New()
	validate.RegisterValidation("is-zhou", ValidateMyVal)
	err := validate.Struct(u)
	return err
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "zhou"
}

func test2() {
	user1 := User{
		Name: "zhou",
	}
	if err := user1.userValidator(); err != nil {
		fmt.Println("user1", err)
	}
	fmt.Println(user1.Name)

	user2 := User{
		Name: "zhou1",
	}
	if err := user2.userValidator(); err != nil {
		fmt.Println("user2", err)
	}
	fmt.Println(user2.Name)
}

type TemS struct {
	s uint64
}

func text3() bool {
	a := 1
	b := 2
	c := 3
	return a < b && a > c
}

func main() {
	// test1()
	// test2()
}
