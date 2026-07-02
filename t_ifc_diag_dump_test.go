package ifc

import (
	"context"
	"os"
	"strings"
	"testing"

	fbtest "typefox.dev/fastbelt/test"
	"typefox.dev/fastbelt/util/service"
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

	validator := service.MustGet[workspace.DocumentValidator](sc)
	diags := validator.Validate(context.Background(), doc.Document, "test")
	for _, d := range diags {
		if d.Code != evalUnknownCode {
			continue
		}
		span := int(d.Range.End.Line) - int(d.Range.Start.Line)
		if span > 10 || strings.Contains(d.Message, "WorldCoordinateSystem") || strings.Contains(d.Message, "parameter value") {
			t.Logf("eval-unknown L%d:%d-L%d:%d span=%d msg=%q",
				d.Range.Start.Line, d.Range.Start.Column,
				d.Range.End.Line, d.Range.End.Column,
				span, d.Message)
		}
	}
}
