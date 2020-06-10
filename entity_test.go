package gohaystack

import (
	"testing"
)

func TestEntity_GetID(t *testing.T) {
	id := NewHaystackID("")
	type fields struct {
		id   *HaystackID
		Dis  string
		tags map[*Label]*Value
	}
	tests := []struct {
		name   string
		fields fields
		want   *HaystackID
	}{
		{
			"simple test",
			fields{
				id: id,
			},
			id,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:   tt.fields.id,
				Dis:  tt.fields.Dis,
				tags: tt.fields.tags,
			}
			if got := e.GetID(); got != tt.want {
				t.Errorf("Entity.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_SetTag(t *testing.T) {
	id := NewHaystackID("")
	type fields struct {
		id   *HaystackID
		Dis  string
		tags map[*Label]*Value
	}
	type args struct {
		l *Label
		v *Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"simple non existent",
			fields{
				id:   id,
				tags: make(map[*Label]*Value),
			},
			args{
				l: nil,
				v: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:   tt.fields.id,
				Dis:  tt.fields.Dis,
				tags: tt.fields.tags,
			}
			e.SetTag(tt.args.l, tt.args.v)
		})
	}
}
