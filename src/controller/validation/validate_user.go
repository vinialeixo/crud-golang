package validation

import (
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/goccy/go-json"
	"github.com/vinialeixo/crud-golang/src/configuration/rest_err"
)

//here we are gonna change the way how to return error message

var (
	Validate = validator.New() //variavel global para utilizar no controler e fazer o translate dessa informação
	transl   ut.Translator
)

//criar uma metodo para receber a menssagem  e fazer a validação

//init() converter para ingles

func init() {
	val, ok := binding.Validator.Engine().(*validator.Validate) //casting in go
	if ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	//jsonErr fazer um unmarshal da tipagem do erro
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationerror validator.ValidationErrors

	//comparar a tipagem se o validation_err é igual a jsonErr? tipos diferentes
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("invalid field type")
	} else if errors.As(validation_err, &jsonValidationerror) {
		//se tiver algum erro é por causa do binding criado para validar o erro
		errorsCauses := []rest_err.Causes{}

		//casting
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)

		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to conjvert fields")
	}
}
