package fields

import (
	"fmt"
	"github.com/junxie6/go-form-it/common"
	"reflect"
	"strconv"
)

// // Text field input type.
// type TextFieldType struct {
// 	Field
// }

// // Password field input type.
// type PasswordFieldType struct {
// 	Field
// }

// // Textarea field input type
// type TextAreaFieldType struct {
// 	Field
// }

// // Hidden field input type.
// type HiddenFieldType struct {
// 	Field
// }

// TextField creates a default text input field based on the provided name.
func TextField(name string) *Field {
	return FieldWithType(name, formcommon.TEXT)
}

// PasswordField creates a default password text input field based on the provided name.
func PasswordField(name string) *Field {
	return FieldWithType(name, formcommon.PASSWORD)
}

// =========== TEXT AREA

// TextAreaField creates a default textarea input field based on the provided name and dimensions.
func TextAreaField(name string, rows, cols int) *Field {
	ret := FieldWithType(name, formcommon.TEXTAREA)
	ret.SetParam("rows", strconv.Itoa(rows))
	ret.SetParam("cols", strconv.Itoa(cols))
	return ret
}

// ========================

// HiddenField creates a default hidden input field based on the provided name.
func HiddenField(name string) *Field {
	return FieldWithType(name, formcommon.HIDDEN)
}

// TextFieldFromInstance creates and initializes a text field based on its name, the reference object instance and field number.
// It uses i object's [fieldNo]-th field content to set the field content.
func TextFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := TextField(name)
	ret.SetValue(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}

// PasswordFieldFromInstance creates and initializes a password field based on its name, the reference object instance and field number.
// It uses i object's [fieldNo]-th field content to set the field content.
func PasswordFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := PasswordField(name)
	ret.SetValue(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}

// TextFieldFromInstance creates and initializes a text field based on its name, the reference object instance and field number.
// This method looks for "form_rows" and "form_cols" tags to add additional parameters to the field.
// It also uses i object's [fieldNo]-th field content to set the field content.
func TextAreaFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	t := reflect.TypeOf(i).Field(fieldNo).Tag
	var rows, cols int64 = 30, 50
	var err error
	if v := t.Get("form_rows"); v != "" {
		rows, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
	}
	if v := t.Get("form_col"); v != "" {
		cols, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil
		}
	}
	ret := TextAreaField(name, int(rows), int(cols))
	ret.SetText(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}

// HiddenFieldFromInstance creates and initializes a hidden field based on its name, the reference object instance and field number.
// It uses i object's [fieldNo]-th field content to set the field content.
func HiddenFieldFromInstance(i interface{}, fieldNo int, name string) *Field {
	ret := HiddenField(name)
	ret.SetValue(fmt.Sprintf("%s", reflect.ValueOf(i).Field(fieldNo).String()))
	return ret
}
