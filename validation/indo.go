package validation

import (
	"fmt"
	"log"
	"unicode"
	"bytes"

	"github.com/go-playground/locales/id"
	"github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	id_translations "gopkg.in/go-playground/validator.v9/translations/id"
)

func FormErrorID(a interface{}) []map[string]string {

	translator := id.New()
	uni := ut.New(translator, translator)
	var list []map[string]string 

	// this is usually known or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, found := uni.GetTranslator("id")
	if !found {
		log.Println("translator not found")
	}

	v := validator.New()
	v.SetTagName("validate")

	if err := id_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Println(err)
	}

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} dibutuhkan", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", addSpace(fe.Field()))
		return t
	})

	_ = v.RegisterTranslation("unique", trans, func(ut ut.Translator) error {
		return ut.Add("unique", "{0} harus unik", true) 
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique", addSpace(fe.Field()))
		return t
	})

	_ = v.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} tidak valid", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", addSpace(fe.Field()))
		return t
	})

	_ = v.RegisterTranslation("numeric", trans, func(ut ut.Translator) error {
		return ut.Add("numeric", "{0} bukan angka", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("numeric", addSpace(fe.Field()))
		return t
	})

	_ = v.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min","{0} tidak boleh kurang dari {1} karakter", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", addSpace(fe.Field()), fe.Param())
		return t
	})

	_ = v.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "{0} tidak boleh lebih dari {1} karakter", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", addSpace(fe.Field()), fe.Param())
		return t
	})
	
	_ = v.RegisterTranslation("gt", trans, func(ut ut.Translator) error {
		return ut.Add("gt", "{0} tidak lebih dari {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gt", addSpace(fe.Field()), fe.Param())
		return t
	})

	_ = v.RegisterTranslation("passwd", trans, func(ut ut.Translator) error {
		return ut.Add("passwd", "{0} kurang panjang", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwd", addSpace(fe.Field()))
		return t
	})

	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})
	
	
	err := v.Struct(a)

	if(err != nil){

		for _, e := range err.(validator.ValidationErrors) {
			data := fmt.Sprintf("%v", e.Translate(trans))
			// utils.SetFlashdata(c, fmt.Sprintf("%v-msg", e.Field()), data)
			list = append(list, map[string]string{
				fmt.Sprintf("%v",e.Field()): data,
			})
		}
		
		return list

	} else {
		return list
	}
		
}

func addSpace(s string) string {
    buf := &bytes.Buffer{}
    for i, rune := range s {
        if unicode.IsUpper(rune) && i > 0 {
            buf.WriteRune(' ')
        }
        buf.WriteRune(rune)
    }
    return buf.String()
}