package cmps

import (
	"fmt"
	"strconv"
	"strings"
)

type Options struct {
	Cmps   float64
	Fields []string
	Groups []string
}

func (o Options) OrderBy() []string {
	return []string{"Cmps"}
}

func (o *Options) parse(tag string) {
	tags := strings.Split(tag, ";")
	cmps, err := strconv.ParseFloat(tags[0], 64)
	if err != nil {
		panic(fmt.Errorf("The tag: %v is not a number(%v)", tags[0], err))
	}
	o.Cmps = cmps
	for _, t := range tags {
		ts := strings.Split(t, ":")
		switch strings.ToLower(ts[0]) {
		case "fields":
			o.Fields = strings.Split(ts[1], ",")
		case "groups":
			o.Groups = strings.Split(ts[1], ",")
		}
	}
}
