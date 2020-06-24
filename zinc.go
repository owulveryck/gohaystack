package gohaystack

import (
	"errors"
	"strings"
)

// MarshalZinc encode the grrind in zinc format.
// https://www.project-haystack.org/doc/Zinc
func (g *Grid) MarshalZinc() ([]byte, error) {
	var b strings.Builder
	err := g.marshalMetaZinc(&b)
	if err != nil {
		return nil, err
	}
	// Writing columns
	labels := g.marshalColsZinc(&b)
	// Values
	err = g.marshalRowsZinc(&b, labels)
	if err != nil {
		return nil, err
	}
	return []byte(b.String()), nil
}

func (g *Grid) marshalRowsZinc(b *strings.Builder, labels []*Label) error {
	i := 0
	for _, entity := range g.entities {
		for _, l := range labels {
			if v, ok := entity.GetTags()[l]; ok {
				z, err := v.MarshalZinc()
				if err != nil {
					return err
				}
				b.Write(z)
			}

			if i < len(labels)-1 {
				b.WriteString(`,`)
			}
			i++
		}
		b.WriteString(`\n`)
	}
	return nil
}

// MarshalZinc for a value
func (v *Value) MarshalZinc() ([]byte, error) {
	if v == nil {
		return nil, errors.New("cannot marshal this type")
	}
	return nil, nil
}

func (g *Grid) marshalMetaZinc(b *strings.Builder) error {
	var hasVer bool
	var version string
	if v, ok := g.Meta["Ver"]; ok {
		hasVer = true
		version = v
	}
	if v, ok := g.Meta["ver"]; ok {
		hasVer = true
		version = v
	}
	if !hasVer {
		return errors.New("Bad formatting, missing version tag")
	}
	if version != "3.0" {
		return errors.New("Unsupported version " + version)
	}
	// Writing header
	b.WriteString(`ver:"3.0"`)
	for t, v := range g.Meta {
		if t == "ver" || t == "Ver" {
			continue
		}
		b.WriteString(` ` + t + `:`)
		b.WriteString(`"` + v + `"`)
	}
	b.WriteString(`\n`)
	return nil
}

func (g *Grid) marshalColsZinc(b *strings.Builder) []*Label {
	labelsDic := make(map[*Label]struct{}, 0)
	for _, entity := range g.entities {
		for label := range entity.tags {
			labelsDic[label] = struct{}{}
		}
	}
	labels := make([]*Label, 0, len(labelsDic))
	i := 0
	for l := range labelsDic {
		b.WriteString(l.Value)
		if l.Display != "" {
			b.WriteString(` dis:"` + l.Display + `"`)
		}
		labels = append(labels, l)
		if i < len(labelsDic)-1 {
			b.WriteString(`,`)
		}
		i++
	}
	if len(labels) > 0 {
		b.WriteString(`\n`)
	}
	return labels
}
