package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSnowFlake(t *testing.T) {
	incrVal := 123456789
	uid, _ := strconv.ParseInt(
		fmt.Sprintf("%d%06d", time.Now().Unix(), incrVal%1000000),
		10, 64)
	fmt.Printf("uid[%v]", uid)
}
