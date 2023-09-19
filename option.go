package cmps

import (
	"fmt"
	"strconv"
	"strings"
)

type Options struct {
	Cmps   float64 `cmps:"1"`
	Fields []string
	Groups []string
	Order  int
}

func (o *Options) parse(tag string, noCmps bool) {
	tags := strings.Split(tag, ";")
	if !noCmps {
		cmps, err := strconv.ParseFloat(tags[0], 64)
		if err != nil {
			panic(fmt.Errorf("the tag: %v is not a number(%v)", tags[0], err))
		}
		o.Cmps = cmps
	}
	o.Order = 1
	for _, t := range tags {
		ts := strings.Split(t, ":")
		switch strings.ToLower(ts[0]) {
		case "fields":
			o.Fields = strings.Split(ts[1], ",")
		case "groups":
			o.Groups = strings.Split(ts[1], ",")
		case "order":
			if strings.ToLower(ts[1]) == "desc" {
				o.Order = -1
			}
		}
	}
}
