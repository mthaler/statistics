package main

import (
	"bytes"
	"testing"

	testdataloader "github.com/peteole/testdata-loader"
	"github.com/stretchr/testify/assert"
)

func TestReadCsvFile(t *testing.T) {
	b := testdataloader.GetTestFile("testdata/1990s/1992-93/eng.1.csv")
	r := bytes.NewReader(b)
	d, _ := readCsvFile(r)
	assert.Equal(t, []string{"1", "Sat Aug 15 1992", "Arsenal FC", "2-4", "Norwich City FC"}, d[1])
}
