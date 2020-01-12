package main

import (
	"testing"
    "sort"
)

func TestTop10(t *testing.T) {
	TestString := "sledge iq iq iq thatcher hibana ash  ash montagne ash hibana montagne montagne montagne hibana hibana montagne hibana buck buck sledge sledge buck maverick maverick hibana maverick ash thermite maverick sledge thermite ash montagne  hibana iq thermite sledge montagne montagne iq sledge hibana sledge twitch thatcher blitz sledge iq sledge nomad fuze ash ash glaz buck hibana nomad maverick nokk amaru sledge lion"
	check := []string{"sledge", "hibana", "montagne" ,"ash", "iq", "maverick", "buck", "thermite", "thatcher", "nomad"}
	result := Top10(TestString)
	sort.Strings(check)
	sort.Strings(result)
	// я не понял как проще сравнить два слайса
	for i, v := range result {
		if v != check[i]{
			t.Fatalf("test failed %s not equal %s", v, check[i])
		}
	}
}
