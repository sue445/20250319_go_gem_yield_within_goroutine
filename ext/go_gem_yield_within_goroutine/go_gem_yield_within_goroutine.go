package main

/*
#include "go_gem_yield_within_goroutine.h"

VALUE rb_go_gem_yield_within_goroutine_sum(VALUE self, VALUE a, VALUE b);
*/
import "C"

import (
	"github.com/ruby-go-gem/go-gem-wrapper/ruby"
)

//export rb_go_gem_yield_within_goroutine_sum
func rb_go_gem_yield_within_goroutine_sum(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	longA := ruby.NUM2LONG(ruby.VALUE(a))
	longB := ruby.NUM2LONG(ruby.VALUE(b))

	sum := longA + longB

	return C.VALUE(ruby.LONG2NUM(sum))
}

//export Init_go_gem_yield_within_goroutine
func Init_go_gem_yield_within_goroutine() {
	rb_mGoGemYieldWithinGoroutine := ruby.RbDefineModule("GoGemYieldWithinGoroutine")
	ruby.RbDefineSingletonMethod(rb_mGoGemYieldWithinGoroutine, "sum", C.rb_go_gem_yield_within_goroutine_sum, 2)
}

func main() {
}
