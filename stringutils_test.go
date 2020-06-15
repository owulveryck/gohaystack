package gohaystack

import (
	"reflect"
	"testing"
)

func Test_trimDoubleQuoteRight(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"valid uniq quote",
			args{
				[]byte(`bla"`),
			},
			[]byte(`bla`),
		},
		{
			"valid uniq two quote",
			args{
				[]byte(`"bla"`),
			},
			[]byte(`"bla`),
		},
		{
			"valid middle entry",
			args{
				[]byte(`bla"bla`),
			},
			[]byte(`bla"bla`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimDoubleQuoteRight(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimDoubleQuoteRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimDoubleQuoteLeft(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"valid uniq quote",
			args{
				[]byte(`"bla`),
			},
			[]byte(`bla`),
		},
		{
			"valid uniq two quote",
			args{
				[]byte(`"bla"`),
			},
			[]byte(`bla"`),
		},
		{
			"valid middle entry",
			args{
				[]byte(`bla"bla`),
			},
			[]byte(`bla"bla`),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimDoubleQuoteLeft(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimDoubleQuoteLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimDoubleQuote(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"valid uniq quote",
			args{
				[]byte(`"bla`),
			},
			[]byte(`bla`),
		},
		{
			"valid uniq two quote",
			args{
				[]byte(`"bla"`),
			},
			[]byte(`bla`),
		},
		{
			"valid middle entry",
			args{
				[]byte(`bla"bla`),
			},
			[]byte(`bla"bla`),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimDoubleQuote(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimDoubleQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidString(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"valid",
			args{
				[]byte(`""`),
			},
			true,
		},
		{
			"invalid",
			args{
				[]byte(`"`),
			},
			false,
		},
		{
			"invalid",
			args{
				[]byte(``),
			},
			false,
		},
		{
			"invalid",
			args{
				[]byte(`{""}`),
			},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidString(tt.args.data); got != tt.want {
				t.Errorf("isValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
