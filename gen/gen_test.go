package gen

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("package tmp \n")
	//err := genQuery(&buf, []string{"people", "price"})
	err := queryGen(&buf, "", map[string]string{})
	fmt.Println(err)
	err = writeFile("../tmp", "people.query", buf)
	fmt.Println(err)
}
