package main

import (
	"reflect"
	"testing"
)

type Data struct {
	X int
}

// ============================================
//
// Command:
//
// go test -run None -bench .
//
// Results:
//
// goos: darwin
// goarch: amd64
// pkg: yifei.com/reflection
// cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
// BenchmarkPlainSet-8         	1000000000	         0.4644 ns/op
// BenchmarkReflectioSet-8     	98572141	        10.65 ns/op
// BenchmarkPlainCall-8        	554181681	         2.090 ns/op
// BenchmarkReflectionCall-8   	 7473148	       162.1 ns/op
// PASS
// ok  	yifei.com/reflection	4.832s

// ============================================
// setup

var data = new(Data)
var field = reflect.ValueOf(data).Elem().FieldByName("X")
var newVal = 100
var newValValue = reflect.ValueOf(newVal)

func set() {
	data.X = newVal
}

func reflectioSet() {
	field.Set(newValValue)
}

func BenchmarkPlainSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		set()
	}
}

func BenchmarkReflectioSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflectioSet()
	}
}

// ============================================
// setup

var method = reflect.ValueOf(data).MethodByName("Inc")

func (data *Data) Inc() {
	data.X++
}

func BenchmarkPlainCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data.Inc()
	}
}

func BenchmarkReflectionCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		method.Call(nil)
	}
}
