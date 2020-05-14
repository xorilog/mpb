package mpb

import (
	"bytes"
	"testing"
	"unicode/utf8"
)

func TestDraw(t *testing.T) {
	// key is termWidth
	testSuite := map[int][]struct {
		name           string
		style          string
		total, current int64
		barWidth       int
		trimSpace      bool
		reverse        bool
		rup            int64
		want           string
	}{
		0: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "…  ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "",
			},
		},
		1: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "…  ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "",
			},
		},
		2: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "  ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[]",
			},
		},
		3: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    "  ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[-]",
			},
		},
		4: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [] ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[>-]",
			},
		},
		5: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [-] ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[>--]",
			},
		},
		6: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [>-] ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[>---]",
			},
		},
		7: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [>--] ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[=>---]",
			},
		},
		8: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [>---] ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[=>----]",
			},
		},
		80: {
			{
				name:    "t,c{60,20}",
				total:   60,
				current: 20,
				want:    " [========================>---------------------------------------------------] ",
			},
			{
				name:      "t,c{60,20}trim",
				total:     60,
				current:   20,
				trimSpace: true,
				want:      "[=========================>----------------------------------------------------]",
			},
			{
				name:     "t,c,bw{60,20,60}",
				total:    60,
				current:  20,
				barWidth: 60,
				want:     " [==================>---------------------------------------] ",
			},
			{
				name:      "t,c,bw{60,20,60}trim",
				total:     60,
				current:   20,
				barWidth:  60,
				trimSpace: true,
				want:      "[==================>---------------------------------------]",
			},
		},
		100: {
			{
				name:    "t,c{100,0}",
				total:   100,
				current: 0,
				want:    " [------------------------------------------------------------------------------------------------] ",
			},
			{
				name:      "t,c{100,0}trim",
				total:     100,
				current:   0,
				trimSpace: true,
				want:      "[--------------------------------------------------------------------------------------------------]",
			},
			{
				name:    "t,c{100,1}",
				total:   100,
				current: 1,
				want:    " [>-----------------------------------------------------------------------------------------------] ",
			},
			{
				name:      "t,c{100,1}trim",
				total:     100,
				current:   1,
				trimSpace: true,
				want:      "[>-------------------------------------------------------------------------------------------------]",
			},
			{
				name:    "t,c{100,33}",
				total:   100,
				current: 33,
				want:    " [===============================>----------------------------------------------------------------] ",
			},
			{
				name:      "t,c{100,33}trim",
				total:     100,
				current:   33,
				trimSpace: true,
				want:      "[===============================>------------------------------------------------------------------]",
			},
			{
				name:      "t,c{100,33}trim,rev",
				total:     100,
				current:   33,
				trimSpace: true,
				reverse:   true,
				want:      "[------------------------------------------------------------------<===============================]",
			},
			{
				name:    "t,c,rup{100,33,33}",
				total:   100,
				current: 33,
				rup:     33,
				want:    " [+++++++++++++++++++++++++++++++>----------------------------------------------------------------] ",
			},
			{
				name:      "t,c,rup{100,33,33}trim",
				total:     100,
				current:   33,
				rup:       33,
				trimSpace: true,
				want:      "[+++++++++++++++++++++++++++++++>------------------------------------------------------------------]",
			},
			{
				name:      "t,c,rup{100,33,33}trim,rev",
				total:     100,
				current:   33,
				rup:       33,
				trimSpace: true,
				reverse:   true,
				want:      "[------------------------------------------------------------------<+++++++++++++++++++++++++++++++]",
			},
			{
				name:    "t,c,rup{100,40,33}",
				total:   100,
				current: 40,
				rup:     33,
				want:    " [++++++++++++++++++++++++++++++++=====>----------------------------------------------------------] ",
			},
			{
				name:      "t,c,rup{100,40,33}trim",
				total:     100,
				current:   40,
				rup:       33,
				trimSpace: true,
				want:      "[++++++++++++++++++++++++++++++++======>-----------------------------------------------------------]",
			},
			{
				name:    "t,c{100,99}",
				total:   100,
				current: 99,
				want:    " [==============================================================================================>-] ",
			},
			{
				name:      "t,c{100,99}trim",
				total:     100,
				current:   99,
				trimSpace: true,
				want:      "[================================================================================================>-]",
			},
			{
				name:    "t,c{100,100}",
				total:   100,
				current: 100,
				want:    " [================================================================================================] ",
			},
			{
				name:      "t,c{100,100}trim",
				total:     100,
				current:   100,
				trimSpace: true,
				want:      "[==================================================================================================]",
			},
			{
				name:    "[=の-] t,c{100,100}",
				style:   "[=の-]",
				total:   100,
				current: 1,
				want:    " [の---------------------------------------------------------------------------------------------…] ",
			},
		},
	}

	var tmpBuf bytes.Buffer
	for tw, cases := range testSuite {
		for _, tc := range cases {
			s := newTestState(tc.style, tc.reverse)
			s.reqWidth = tc.barWidth
			s.total = tc.total
			s.current = tc.current
			s.trimSpace = tc.trimSpace
			if tc.rup > 0 {
				if f, ok := s.filler.(interface{ SetRefill(int64) }); ok {
					f.SetRefill(tc.rup)
				}
			}
			tmpBuf.Reset()
			tmpBuf.ReadFrom(s.draw(newStatistics(tw, s)))
			by := tmpBuf.Bytes()

			got := string(by[:len(by)-1])
			if !utf8.ValidString(got) {
				t.Fail()
			}
			if got != tc.want {
				t.Errorf("termWidth:%d %q want: %q %d, got: %q %d\n", tw, tc.name, tc.want, utf8.RuneCountInString(tc.want), got, utf8.RuneCountInString(got))
			}
		}
	}
}

func newTestState(style string, reverse bool) *bState {
	if style == "" {
		style = DefaultBarStyle
	}
	s := &bState{
		filler: NewBarFiller(style, reverse),
		bufP:   new(bytes.Buffer),
		bufB:   new(bytes.Buffer),
		bufA:   new(bytes.Buffer),
	}
	return s
}
