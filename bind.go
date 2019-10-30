/**
 * @Author: DollarKiller
 * @Description: 参数绑定
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:48 2019-09-29
 */
package erguotou

import (
	"errors"
	"github.com/dollarkillerx/erguotou/fasthttp"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (c *Context) BindValue(obj interface{}) error {
	return bind(c.Ctx, obj)
}

func (c *Context) BindFrom(obj interface{}) error {
	err := bindFormPost(c.Ctx, obj)
	if err != nil {
		return err
	}
	// 进行validate 验证
	validate := validator.New()
	return validate.Struct(obj)
}

func (c *Context) BindJson(obj interface{}) error {
	err := bindJson(c.Ctx, obj)
	if err != nil {
		return err
	}
	// 进行validate 验证
	validate := validator.New()
	return validate.Struct(obj)
}

func (c *Context) BindGet(obj interface{}) error {
	err := bindFormGet(c.Ctx, obj)
	if err != nil {
		return err
	}
	// 进行validate 验证
	validate := validator.New()
	return validate.Struct(obj)
}

func bind(req *fasthttp.RequestCtx, obj interface{}) error {

	contentType := string(req.Request.Header.ContentType())
	//如果是简单的json
	if strings.Contains(strings.ToLower(contentType), "application/json") {
		return bindJson(req, obj)
	}
	if strings.Contains(strings.ToLower(contentType), "application/x-www-form-urlencoded") {
		return bindFormPost(req, obj)
	}
	if contentType == "" {
		return bindFormGet(req, obj)
	}

	return errors.New("bind error: unsupported method")
}

// 绑定Json数据
func bindJson(req *fasthttp.RequestCtx, obj interface{}) error {
	data := req.PostBody()
	return Jsonp.Unmarshal(data, obj)
}

// 绑定Form数据
func bindFormPost(req *fasthttp.RequestCtx, obj interface{}) error {
	getValues := req.PostArgs()
	return mapForm(obj, getValues)
}

func bindFormGet(req *fasthttp.RequestCtx, obj interface{}) error {
	getValues := req.QueryArgs()
	return mapForm(obj, getValues)
}

//自动绑定方法
func mapForm(ptr interface{}, form *fasthttp.Args) error {
	typ := reflect.TypeOf(ptr).Elem()  // 获取type elem
	val := reflect.ValueOf(ptr).Elem() // 获取val elem

	// NumField 获取 个数
	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i) // 获取第 i 个
		structField := val.Field(i)
		// 如果不能修改就跳过
		if !structField.CanSet() {
			continue
		}

		// 获取值类型
		structFieldKind := structField.Kind()
		inputFieldName := typeField.Tag.Get("form") // 获取type tag

		// 如果 没有 tag 就用原来的名称示意
		if inputFieldName == "" {
			inputFieldName = typeField.Tag.Get("json") // 获取type tag
			if inputFieldName == "" {
				inputFieldName = typeField.Name
			}
		}

		// 允许用户不发生绑定
		if inputFieldName == "-" {
			continue
		}

		// 如果是一个struct 就进行递归
		if structFieldKind == reflect.Struct {
			err := mapForm(structField.Addr().Interface(), form)
			if err != nil {
				return err
			}
			continue
		}

		// 获取 form中元素
		inputValue := string(form.Peek(inputFieldName))
		//clog.Println(inputFieldName)
		//clog.Println(inputValue)
		if inputValue == "" {
			continue
		}

		//numElems := len(inputValue)

		// 如果是slice 专业解析
		//if structFieldKind == reflect.Slice && numElems > 0 {
		//	// 获取 slice 的类型
		//	sliceOf := structField.Type().Elem().Kind()
		//	// 创建slice
		//	slice := reflect.MakeSlice(structField.Type(), numElems, numElems)
		//	for i := 0; i < numElems; i++ {
		//		// 设置
		//		if err := setWithProperType(sliceOf, inputValue[i], slice.Index(i)); err != nil {
		//			return err
		//		}
		//	}
		//	// 设置
		//	val.Field(i).Set(slice)
		//} else {
		//
		//	// 如果是时间类型 用特定解析
		//	if _, isTime := structField.Interface().(time.Time); isTime {
		//		if err := setTimeField(inputValue[0], typeField, structField); err != nil {
		//			return err
		//		}
		//		continue
		//	}
		//	if err := setWithProperType(typeField.Type.Kind(), inputValue[0], structField); err != nil {
		//		return err
		//	}
		//}

		// 如果是时间类型 用特定解析
		if _, isTime := structField.Interface().(time.Time); isTime {
			if err := setTimeField(inputValue, typeField, structField); err != nil {
				return err
			}
			continue
		}
		if err := setWithProperType(typeField.Type.Kind(), inputValue, structField); err != nil {
			return err
		}
	}

	return nil
}

// 设置 元素  类型,val名称,设置目标地址
func setWithProperType(valueKind reflect.Kind, val string, structField reflect.Value) error {
	switch valueKind {
	case reflect.Int:
		return setIntField(val, 0, structField)
	case reflect.Int8:
		return setIntField(val, 8, structField)
	case reflect.Int16:
		return setIntField(val, 16, structField)
	case reflect.Int32:
		return setIntField(val, 32, structField)
	case reflect.Int64:
		return setIntField(val, 64, structField)
	case reflect.Uint:
		return setUintField(val, 0, structField)
	case reflect.Uint8:
		return setUintField(val, 8, structField)
	case reflect.Uint16:
		return setUintField(val, 16, structField)
	case reflect.Uint32:
		return setUintField(val, 32, structField)
	case reflect.Uint64:
		return setUintField(val, 64, structField)
	case reflect.Bool:
		return setBoolField(val, structField)
	case reflect.Float32:
		return setFloatField(val, 32, structField)
	case reflect.Float64:
		return setFloatField(val, 64, structField)
	case reflect.String:
		structField.SetString(val)
	default:
		return errors.New("Unknown type")
	}
	return nil
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	if val == "" {
		val = "false"
	}
	boolVal, err := strconv.ParseBool(val)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func setTimeField(val string, structField reflect.StructField, value reflect.Value) error {
	timeFormat := structField.Tag.Get("time_format")
	//2018-01-02 01:02:03

	if timeFormat == "" {
		timeFormat = "2006-01-02 15:04:05"
		val = strings.Replace(val, "/", "-", 0)
		num := len(strings.Split(val, " "))
		if num == 1 {
			val = val + " 00:00:00"
		} else {
			//2018-01-02 00
			num = len(strings.Split(val, ":"))

			if num == 1 {
				val = val + ":00:00"
			} else if num == 2 {
				val = val + ":00"
			}
		}

	}

	if val == "" {
		value.Set(reflect.ValueOf(time.Time{}))
		return nil
	}

	l := time.Local
	if isUTC, _ := strconv.ParseBool(structField.Tag.Get("time_utc")); isUTC {
		l = time.UTC
	}

	if locTag := structField.Tag.Get("time_location"); locTag != "" {
		loc, err := time.LoadLocation(locTag)
		if err != nil {
			return err
		}
		l = loc
	}

	t, err := time.ParseInLocation(timeFormat, val, l)
	if err != nil {
		return err
	}

	value.Set(reflect.ValueOf(t))
	return err
}
