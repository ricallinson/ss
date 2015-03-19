package main

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestRsa(t *testing.T) {
	Describe("ProcessFiles()", func() {
		It("should count all occurrences of 'GET' in a single file", func() {
			q := Query("GET", nil, 0)
			f := GetFiles("./fixtures/basic/1.txt")
			t := ProcessFiles(f, q)
			AssertEqual(t, 3)
		})
		It("should count all occurrences of 'GET'", func() {
			q := Query("GET", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 15)
		})
		It("should count all occurrences of 'GET' AND '302'", func() {
			q := Query("GET 302", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 5)
		})
		It("should count all occurrences of NOT '302'", func() {
			q := Query("NOT 302", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 15)
		})
		It("should count all occurrences of 'GET' NOT '302'", func() {
			q := Query("GET NOT 302", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 10)
		})
		It("should count all occurrences of 'PUT' OR 'POST'", func() {
			q := Query("PUT OR POST", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 10)
		})
		It("should count all occurrences of '200' NOT 'PUT' OR 'POST'", func() {
			q := Query("200 NOT PUT OR POST", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 10)
		})
		It("should count all occurrences of '200' OR '302' NOT 'PUT'", func() {
			q := Query("200 OR 302 NOT PUT", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 20)
		})
		It("should count all occurrences of '200' OR '302' NOT 'PUT' OR 'POST'", func() {
			q := Query("200 OR 302 NOT PUT OR POST", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 15)
		})
		It("should count all occurrences of '200' 'GET' OR 'PUT'", func() {
			q := Query("200 GET OR PUT", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 15)
		})
		It("should count all occurrences of NOT '200' 'GET' OR 'PUT'", func() {
			q := Query("NOT 200 GET OR PUT", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 5)
		})
		It("should count all occurrences of 'GET' OR 'PUT' OR 'POST'", func() {
			q := Query("GET OR PUT OR POST", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 25)
		})
		It("should count all occurrences of NOT 'GET' OR 'PUT' OR 'POST'", func() {
			q := Query("NOT GET OR PUT OR POST", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 0)
		})
		It("should count all occurrences of 'GET' OR 'PUT' and 'status' OR 'update'", func() {
			q := Query("GET OR PUT status OR update", nil, 0)
			f := GetFiles("./fixtures/basic")
			t := ProcessFiles(f, q)
			AssertEqual(t, 10)
		})
	})

	Describe("ProcessFiles() Bug Fixes", func() {
		It("should count all occurrences of NOT '302'", func() {
			q := Query("NOT 302", nil, 0)
			f := GetFiles("./fixtures")
			t := ProcessFiles(f, q)
			AssertEqual(t, 15)
		})
	})

	Report(t)
}
