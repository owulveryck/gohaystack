package gohaystack

import (
	"errors"
	"fmt"
	"math"
	"strconv"
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
				if v == nil {
					return errors.New("nil value")
				}
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
		b.WriteString("\n")
	}
	return nil
}

// MarshalZinc for a value
func (v *Value) MarshalZinc() ([]byte, error) {
	switch v.kind {
	case HaystackLastType:
		return nil, nil
	case HaystackTypeUndefined:
		return nil, errors.New("cannot marshal this type")
	case HaystackTypeGrid:
		b, err := v.g.MarshalZinc()
		if err != nil {
			return nil, err
		}
		return []byte("<<\n" + string(b) + "\n>>\n"), nil
	case HaystackTypeNull:
		return []byte(`N`), nil
	case HaystackTypeBool:
		if v.b {
			return []byte(`T`), nil
		}
		return []byte(`F`), nil
	case HaystackTypeMarker:
		return []byte(`M`), nil
	case HaystackTypeRemove:
		return []byte(`R`), nil
	case HaystackTypeNA:
		return []byte(`NA`), nil
	case HaystackTypeNumber:
		output := strconv.FormatFloat(float64(v.number.value), 'g', -1, 32)
		if v.number.value == math.MaxFloat32 {
			output = "INF"
		}
		if v.number.value == -math.MaxFloat32 {
			output = "-INF"
		}
		if v.number.unit != nil {
			output = output + " " + *v.number.unit
		}
		return []byte(output), nil
	case HaystackTypeRef:
		return []byte(`@` + *v.ref), nil
	case HaystackTypeStr:
		return []byte(`"` + *v.str + `"`), nil
	case HaystackTypeDate:
		return []byte(v.t.Format("2006-01-02")), nil
	case HaystackTypeTime:
		return []byte(v.t.Format("15:04:05.999999999")), nil
	case HaystackTypeDateTime:
		return []byte(v.t.Format("2006-02-01T15:04:05.999-0700 MST")), nil
	case HaystackTypeURI:
		return []byte("`" + v.u.String() + "`"), nil
	case HaystackTypeCoord:
		return []byte(fmt.Sprintf("C(%v,%v)", v.coord.lat, v.coord.long)), nil
	case HaystackTypeList:
		output := []byte("[")
		for i := 0; i < len(v.list); i++ {
			b, err := v.list[i].MarshalZinc()
			if err != nil {
				return nil, err
			}
			output = append(output, b...)
			if i < len(v.list)-1 {
				output = append(output, []byte(",")...)
			}
		}
		output = append(output, []byte("]")...)
		return output, nil
	case HaystackTypeDict:
		output := []byte("{")
		i := 0
		l := len(v.dict)
		for k, v := range v.dict {
			b, err := v.MarshalZinc()
			if err != nil {
				return nil, err
			}
			output = append(output, []byte(k+":")...)
			output = append(output, b...)
			if i < l-1 {
				output = append(output, []byte(",")...)
			}
			i++
		}
		output = append(output, []byte("}")...)
		return output, nil

	case HaystackTypeXStr:
		return nil, errors.New("not implemented")
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
	b.WriteString("\n")
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
		b.WriteString("\n")
	}
	return labels
}
