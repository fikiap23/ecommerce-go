package utils

import "time"

func PtrString(s string) *string       { return &s }
func PtrInt(i int) *int               { return &i }
func PtrTime(t time.Time) *time.Time  { return &t }
func PtrBool(b bool) *bool             { return &b }
