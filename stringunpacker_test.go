package stringunpacker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {

	cases := []struct {
		in  string
		out string
		err bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"ф4bc2d5e", "ффффbccddddde", false},
		{"abcd", "abcd", false},
		{`45`, ``, true},
		{`4a5b`, `aaaaab`, false},
		// {`qwe\4\5`, `qwe45`, false},
		// {`qwe\45`, `qwe44444`, false},
		// {`qwe\\5`, `qwe\\\\\`, false}, // ?
	}

	for _, testCase := range cases {
		t.Logf("Expect ERR:%v IN: %v, OUT: %v", testCase.err, testCase.in, testCase.out)
		out, err := UnpackString(testCase.in)
		if (err != nil) == testCase.err {
			// t.Logf("OK")
		} else {
			t.Errorf(" Get %v %v FAIL\n", err, testCase.out)
			continue
		}

		if out == testCase.out {
			t.Logf("OK\n")
		} else {

			for i:=0; i <len(out); i+=1 {
				fmt.Printf("%x ", out[i])
			}
			fmt.Println()

			for i:=0; i <len(testCase.out); i+=1 {
				fmt.Printf("%x ", testCase.out[i])
			}
			fmt.Println()

			t.Errorf("Get %v %v %v FAIL\n", out, reflect.TypeOf(out), reflect.TypeOf(testCase.out))
			t.Errorf("Get %v %v FAIL\n", len(out), len(testCase.out))
		}

	}

}
