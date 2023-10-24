package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/raafly/inventory-management/entity"
)

func TestCreateTime(t *testing.T) {
	time := time.Now()
	data := entity.Item {
		Created_at: time,
	}

	fmt.Println(data)
}