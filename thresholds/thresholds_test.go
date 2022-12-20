/*
	This file is part of icinga_plugins.

Icinga Plugins Support is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Icinga Plugins Support is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Icinga Plugins Support.  If not, see <http://www.gnu.org/licenses/>.
*/
package thresholds

import (
	"errors"
	"math"
	"strings"
	"testing"

	icinga "github.com/marshei/icinga_plugins"
	"github.com/marshei/icinga_plugins/perfdata"
)

func TestRangeListSuccess(t *testing.T) {
	list, err := ParseThresholdList("metric1,10:;metric2,@30:")
	if err != nil {
		t.Errorf("Unexpected error occured: %s", err.Error())
	}

	if len(list) != 2 {
		t.Errorf("Expecting list of length %d, got %d", 2, len(list))
	}

	var e icinga.ThresholdRange
	e.Definition = "10:"
	e.Metric = "metric1"
	e.Inside = true
	e.Start = 10
	e.End = math.Inf(1)
	expectRangeSuccess(t, list[0], e, nil)

	e.Definition = "@30:"
	e.Metric = "metric2"
	e.Inside = false
	e.Start = 30
	e.End = math.Inf(1)
	expectRangeSuccess(t, list[1], e, nil)

	var pd perfdata.PerformanceData
	pd.Label = "metric2"
	g := getThreshold(list, &pd)
	expectRangeSuccess(t, *g, e, nil)
}

func TestRangeListErrorEmptyMetric(t *testing.T) {
	_, err := ParseThresholdList("metric1,10:20;,@30:40")
	if err == nil {
		t.Errorf("Expecting an error here")
	}

	expected := "empty metric"
	if err.Error() != expected {
		t.Errorf("Expecting error %s, got %s", expected, err.Error())
	}

	_, err = ParseThresholdList("10:20;@30:40")
	if err == nil {
		t.Errorf("Expecting an error here")
	}

	expected = "missing metric"
	if err.Error() != expected {
		t.Errorf("Expecting error %s, got %s", expected, err.Error())
	}
}

func TestRangeListErrorMissingMetric(t *testing.T) {
	_, err := ParseThresholdList("metric1,10:20;@30:40")
	if err == nil {
		t.Errorf("Expecting an error here")
	}

	expected := "missing metric"
	if err.Error() != expected {
		t.Errorf("Expecting error %s, got %s", expected, err.Error())
	}
}

func TestRangesSuccess(t *testing.T) {
	rangeSuccess(t, "metric,11.34", "11.34", "metric", true, 0, 11.34)
	rangeSuccess(t, "0:11.34", "0:11.34", "", true, 0, 11.34)
	rangeSuccess(t, "-11.34:11.34", "-11.34:11.34", "", true, -11.34, 11.34)
	rangeSuccess(t, "11.34:", "11.34:", "", true, 11.34, math.Inf(1))
	rangeSuccess(t, ":11.34", ":11.34", "", true, math.Inf(-1), 11.34)
	rangeSuccess(t, "~:11.34", "~:11.34", "", true, math.Inf(-1), 11.34)

	rangeSuccess(t, "@11.34", "@11.34", "", false, 0, 11.34)
	rangeSuccess(t, "@0:11.34", "@0:11.34", "", false, 0, 11.34)
	rangeSuccess(t, "metric,@-11.34:11.34", "@-11.34:11.34", "metric", false, -11.34, 11.34)
	rangeSuccess(t, "@11.34:", "@11.34:", "", false, 11.34, math.Inf(1))
	rangeSuccess(t, "@:11.34", "@:11.34", "", false, math.Inf(-1), 11.34)
	rangeSuccess(t, "metric,@~:11.34", "@~:11.34", "metric", false, math.Inf(-1), 11.34)
}

func rangeSuccess(t *testing.T, rangeDef string, def string, metric string, inside bool, start float64, end float64) {
	var e icinga.ThresholdRange
	e.Definition = def
	e.Metric = metric
	e.Inside = inside
	e.Start = start
	e.End = end
	r, err := parseThreshold(rangeDef)
	expectRangeSuccess(t, r, e, err)
}

func expectRangeSuccess(t *testing.T, current icinga.ThresholdRange, expected icinga.ThresholdRange, err error) {
	if err != nil {
		t.Error(err)
	}

	if current.Definition != expected.Definition {
		t.Errorf("Difference for Definition: curr = %s, expected = %s", current.Definition, expected.Definition)
	}

	if current.Metric != expected.Metric {
		t.Errorf("Difference for Metric: curr = %s, expected = %s", current.Metric, expected.Metric)
	}

	if current.Inside != expected.Inside {
		t.Errorf("Difference for Inside: curr = %t, expected = %t", current.Inside, expected.Inside)
	}

	if current.Start != expected.Start {
		t.Errorf("Difference for Start: curr = %f, expected = %f", current.Start, expected.Start)
	}

	if current.End != expected.End {
		t.Errorf("Difference for End: curr = %f, expected = %f", current.End, expected.End)
	}
}

func TestRangesError(t *testing.T) {
	rangeError(t, "", "empty range")
	rangeError(t, "@", "empty range")
	rangeError(t, ":", "empty range")
	rangeError(t, "1:2:3", "invalid range: too many values")
	rangeError(t, "1:B", "parsing \"B\": invalid syntax")
	rangeError(t, "B", "parsing \"B\": invalid syntax")
	rangeError(t, "@B", "parsing \"B\": invalid syntax")
	rangeError(t, "metric,abc,@B", "invalid metric")
	rangeError(t, ",@1", "empty metric")
}

func rangeError(t *testing.T, rangeDef string, message string) {
	_, err := parseThreshold(rangeDef)
	expectRangeError(t, message, err)
}

func expectRangeError(t *testing.T, message string, err error) {
	if err == nil {
		t.Error(errors.New("Expecting an error but was successful"))
	}

	if !strings.Contains(err.Error(), message) {
		t.Errorf("Expecting error: %s, got = %s", message, err.Error())
	}
}

func TestRangeEvaluation(t *testing.T) {
	evaluateRange(t, "10:20", 10, false)
	evaluateRange(t, "10:20", 20, false)
	evaluateRange(t, "10:20", -10, true)
	evaluateRange(t, "@~:20", 20, true)
	evaluateRange(t, "@~:20", -20, true)
	evaluateRange(t, "@~:20", 40, false)
}

func evaluateRange(t *testing.T, rangeDef string, value float64, expectedAlert bool) {
	r, err := parseThreshold(rangeDef)
	if err != nil {
		t.Errorf("Failed to parse range definition: %s", rangeDef)
	}

	isAlert := isValueOutOfRange(r, value)

	if isAlert != expectedAlert {
		if expectedAlert {
			t.Errorf("Alert expected but not received for range %s and value %f", rangeDef, value)
		} else {
			t.Errorf("Alert not expected but received for range %s and value %f", rangeDef, value)
		}
	}
}

func TestStringToFloat(t *testing.T) {

	test_StringToFloatSuccess(t, "1.234", 1.234)
	test_StringToFloatSuccess(t, "-1000.234", -1000.234)
	test_StringToFloatSuccess(t, "~", math.Inf(-1))
}

func test_StringToFloatSuccess(t *testing.T, input string, expected float64) {
	result, err := stringToFloat(input)

	if err != nil {
		t.Errorf("Unexpected error occured: %s", err.Error())
	}

	if result != expected {
		t.Errorf("StringToFloat was incorrect, got: %f, want: %f.", result, expected)
	}
}
