package ifc

import (
	"context"
	"os"
	"strings"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/workspace"
)

func TestDumpTIfcEvalUnknownRanges(t *testing.T) {
	sc := CreateServices()
	fixture := fbtest.New(t, sc)
	raw, err := os.ReadFile("t.ifc")
	if err != nil {
		t.Fatal(err)
	}
	doc := fixture.Parse(string(raw))

	validator := sc.MustGet[workspace.DocumentValidator]()
	diags := validator.Validate(context.Background(), doc.Document)
	for _, d := range diags {
		if d.Code != evalUnknownCode {
			continue
		}
		start := doc.Document.TextDoc.PositionAt(int(d.Range.Start))
		end := doc.Document.TextDoc.PositionAt(int(d.Range.End))
		span := int(end.Line) - int(start.Line)
		if span > 10 || strings.Contains(d.Message, "WorldCoordinateSystem") || strings.Contains(d.Message, "parameter value") {
			t.Logf("eval-unknown L%d:%d-L%d:%d span=%d msg=%q",
				start.Line, start.Character,
				end.Line, end.Character,
				span, d.Message)
		}
	}
}
