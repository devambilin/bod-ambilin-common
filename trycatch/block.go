package trycatch

import "errors"

// Block : for catch unhandling error
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Exception : type of exception
type Exception interface{}

// Throw : throw exception
func Throw(up Exception) {
	panic(up)
}

// Do : run your function with error handling
func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				switch r.(type) {
				case string:
					tcf.Catch(errors.New(r.(string)))
					break
				default:
					tcf.Catch(r)
				}
			}
		}()
	}
	tcf.Try()
}

type OnError func(Exception)

// Catch error
func Catch(onerror OnError) {
	if r := recover(); r != nil {
		switch r.(type) {
		case string:
			onerror(errors.New(r.(string)))
			break
		default:
			onerror(r)
		}
	}
}
