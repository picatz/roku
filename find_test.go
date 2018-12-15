package roku

import "testing"
import "fmt"

func TestFind(t *testing.T) {
	devices, err := Find(5)

	if err != nil {
		t.Error(err)
	}

	t.Log(devices)
}

func ExampleFind() {
	devices, err := Find(2)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(devices))
	// Output: 1
}
