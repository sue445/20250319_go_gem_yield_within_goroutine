package main

/*
#include "go_gem_yield_within_goroutine.h"

VALUE rb_go_gem_yield_within_goroutine_sum(VALUE self, VALUE a, VALUE b);
VALUE rb_go_gem_yield_within_goroutine_with_block(VALUE self, VALUE arg);
*/
import "C"

import (
	"github.com/ruby-go-gem/go-gem-wrapper/ruby"
	"sync"
)

//export rb_go_gem_yield_within_goroutine_sum
func rb_go_gem_yield_within_goroutine_sum(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	longA := ruby.NUM2LONG(ruby.VALUE(a))
	longB := ruby.NUM2LONG(ruby.VALUE(b))

	sum := longA + longB

	return C.VALUE(ruby.LONG2NUM(sum))
}

//export rb_go_gem_yield_within_goroutine_with_block
func rb_go_gem_yield_within_goroutine_with_block(_ C.VALUE, arg C.VALUE) C.VALUE {
	if ruby.RbBlockGivenP() == 0 {
		ruby.RbRaise(ruby.VALUE(C.rb_eArgError), "Block not given")
	}

	var wg sync.WaitGroup
	wg.Add(1)

	var blockResult ruby.VALUE

	go func() {
		defer wg.Done()
		blockResult = ruby.RbYield(ruby.VALUE(arg))
	}()

	wg.Wait()

	return C.VALUE(blockResult)
}

//export Init_go_gem_yield_within_goroutine
func Init_go_gem_yield_within_goroutine() {
	rb_mGoGemYieldWithinGoroutine := ruby.RbDefineModule("GoGemYieldWithinGoroutine")
	ruby.RbDefineSingletonMethod(rb_mGoGemYieldWithinGoroutine, "sum", C.rb_go_gem_yield_within_goroutine_sum, 2)
	ruby.RbDefineSingletonMethod(rb_mGoGemYieldWithinGoroutine, "with_block", C.rb_go_gem_yield_within_goroutine_with_block, 1)
}

func main() {
}
