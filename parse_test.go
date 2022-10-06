package main

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	utcLoc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(fmt.Sprintf("failed to load location 'UTC': %s", err))
	}

	t.Run("2022-10-05t19:59:25.644225z", func(t *testing.T) {
		result, err := Parse("2022-10-05t19:59:25.644225z")
		if err != nil {
			t.Error(err)
		}
		result = result.In(utcLoc)
		expected := time.Date(2022, 10, 5, 19, 59, 25, 644225000, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("2022-09-21T09:52:19Z", func(t *testing.T) {
		result, err := Parse("2022-09-21T09:52:19Z")
		if err != nil {
			t.Error(err)
		}
		result = result.In(utcLoc)
		expected := time.Date(2022, 9, 21, 9, 52, 19, 0, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("2022-10-05T16:27:08.419-04:00", func(t *testing.T) {
		result, err := Parse("2022-10-05T16:27:08.419-04:00")
		if err != nil {
			t.Error(err)
		}
		result = result.In(utcLoc)
		expected := time.Date(2022, 10, 5, 20, 27, 8, 419000000, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("2022-10-05T09:10:19 EDT", func(t *testing.T) {
		result, err := Parse("2022-10-05T09:10:19 EDT")
		if err != nil {
			t.Error(err)
		}
		result = result.In(utcLoc)
		expected := time.Date(2022, 10, 5, 13, 10, 19, 0, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("unix time (seconds)", func(t *testing.T) {
		result, err := Parse("1665001628")
		if err != nil {
			t.Error(err)
		}
		result = result.In(utcLoc)
		expected := time.Date(2022, 10, 5, 20, 27, 8, 0, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("unix time (milliseconds)", func(t *testing.T) {
		result, err := Parse("1665001628419")
		if err != nil {
			t.Error(err)
		}
		result = result.In(utcLoc)
		expected := time.Date(2022, 10, 5, 20, 27, 8, 419000000, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("unix time (nanoseconds)", func(t *testing.T) {
		result, err := Parse("1665001628419000123")
		if err != nil {
			t.Error(err)
		}
		result = result.In(utcLoc)
		expected := time.Date(2022, 10, 5, 20, 27, 8, 419000123, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})
}
