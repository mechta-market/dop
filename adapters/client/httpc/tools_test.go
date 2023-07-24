package httpc

import (
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"github.com/mechta-market/dop/dopTools"
)

func TestObject2UrlValues(t *testing.T) {
	type embStruct struct {
		EF1 int `json:"j_ef1" form:"ef1"`
	}

	tests := []struct {
		obj  any
		want url.Values
	}{
		{
			obj: struct {
				embStruct
				F1 int       `json:"j_f1" form:"f1"`
				F2 int64     `json:"j_f2" form:"f2"`
				F3 uint8     `json:"j_f3" form:"f3"`
				F4 float64   `json:"j_f4" form:"f4"`
				F5 float32   `json:"j_f5" form:"f5"`
				F6 *int      `json:"j_f6" form:"f6"`
				F7 []int     `json:"j_f7" form:"f7"`
				F8 []string  `json:"j_f8" form:"f8"`
				F9 *[]string `json:"j_f9" form:"f9"`
			}{
				embStruct: embStruct{EF1: 77},
				F1:        1,
				F2:        -2,
				F3:        3,
				F4:        3.14,
				F5:        1.32,
				F6:        dopTools.NewPtr(7),
				F7:        []int{1, 2, 3},
				F8:        []string{"h", "e", "l", "l", "o"},
			},
			want: map[string][]string{
				"ef1": {"77"},
				"f1":  {"1"},
				"f2":  {"-2"},
				"f3":  {"3"},
				"f4":  {"3.14"},
				"f5":  {"1.32"},
				"f6":  {"7"},
				"f7":  {"1", "2", "3"},
				"f8":  {"h", "e", "l", "l", "o"},
			},
		},
	}
	for ttI, tt := range tests {
		t.Run(strconv.Itoa(ttI+1), func(t *testing.T) {
			if got := Object2UrlValues(tt.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Object2UrlValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
