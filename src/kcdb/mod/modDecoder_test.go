package mod

import (
	"strings"
	"testing"
)

func TestDecodeMod(t *testing.T) {
	out, err := DecodeModule(strings.NewReader(`
    (module CR2032-white-smd (layer F.Cu) (tedit 566FEBDB)
      (fp_text reference REF** (at 0 -3.3) (layer F.SilkS)
        (effects (font (size 1 1) (thickness 0.15)))
      )
      (fp_text value CR2032-white-smd (at 0 3.35) (layer F.Fab)
        (effects (font (size 1 1) (thickness 0.15)))
      )
      (fp_line (start -1.8 7.45) (end 1.8 7.45) (layer F.SilkS) (width 0.15))
      (fp_line (start -1.8 -7.9) (end 1.8 -7.9) (layer F.SilkS) (width 0.15))
      (fp_text user BAT (at 7 0) (layer F.SilkS)
        (effects (font (size 1 1) (thickness 0.25)))
      )
      (fp_line (start -5.5 0) (end -2.5 2.5) (layer F.SilkS) (width 0.15))
      (fp_line (start 2.5 -2.5) (end 5.5 0) (layer F.SilkS) (width 0.15))
      (fp_text user Adhesive (at 0 0) (layer F.SilkS)
        (effects (font (size 1 1) (thickness 0.15)))
      )
      (fp_line (start -5.5 -2.5) (end -5.5 2.5) (layer F.SilkS) (width 0.15))
      (fp_line (start 5.5 -2.5) (end 5.5 2.5) (layer F.SilkS) (width 0.15))
      (fp_line (start -5.5 -2.5) (end 5.5 -2.5) (layer F.SilkS) (width 0.15))
      (fp_line (start -5.5 2.5) (end 5.5 2.5) (layer F.SilkS) (width 0.15))
      (pad 1 smd rect (at 14.65 0) (size 2.6 3.6) (layers F.Cu F.Paste F.Mask))
      (pad 2 smd rect (at -14.65 0) (size 2.6 3.6) (layers F.Cu F.Paste F.Mask))
    )

    `))

	if err != nil || out == nil {
		t.Errorf("Expected value and no error, got err = %v, out = %+v", err, out)
	}
}
