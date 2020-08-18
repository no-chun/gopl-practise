package main

import "testing"

func Test_tree_String(t1 *testing.T) {
	type fields struct {
		value int
		left  *tree
		right *tree
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "test 1", fields: struct {
			value int
			left  *tree
			right *tree
		}{value: 1, left: nil, right: nil}, want: "1"},

		{name: "test 2", fields: struct {
			value int
			left  *tree
			right *tree
		}{value: 1, left: &tree{
			value: 0,
			left:  nil,
			right: nil,
		}, right: &tree{
			value: 2,
			left:  nil,
			right: nil,
		}}, want: "012"},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &tree{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
