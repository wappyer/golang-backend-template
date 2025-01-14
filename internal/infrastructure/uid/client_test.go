package uid

import (
	"fmt"
	"testing"
)

func TestLongUid(t *testing.T) {
	timeDiff := 9999999999 >> 8
	fmt.Println(timeDiff)
	fmt.Println(timeDiff * 100 / 365 / 86400)

	id := LongUid(1).GetInt64()
	fmt.Println(id)
}

func TestShortUid(t *testing.T) {
	id := ShortUid(1).GetInt64()
	fmt.Println(id)
}

func TestMiniUid(t *testing.T) {
	id := MiniUid(1).GetInt64()
	fmt.Println(id)
}
