/*
	This file is part of icinga_plugins.

Foobar is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Foobar is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
*/
package thresholds

import (
	"errors"
	"math"
	"strconv"
	"strings"

	icinga "github.com/marshei/icinga_plugins"
	"github.com/marshei/icinga_plugins/perfdata"
)

/*
 * Parsing the input into the threshold range structure
 */

// ParseThresholdList parses the provided string into a list of threshold ranges
func ParseThresholdList(thresholdDef string) ([]icinga.ThresholdRange, error) {
	var list []icinga.ThresholdRange
	countNoMetric := 0
	metricParts := strings.FieldsFunc(thresholdDef, func(r rune) bool { return r == ';' })
	for _, p := range metricParts {
		r, err := parseThreshold(p)
		if err != nil {
			return list, err
		}
		if r.Metric == "" {
			countNoMetric++
		}
		list = append(list, r)
	}

	if countNoMetric >= 1 && len(list) > 1 {
		return list, errors.New("missing metric")
	}

	return list, nil
}

func parseThreshold(thresholdDef string) (icinga.ThresholdRange, error) {
	var err error
	var thresholdRange icinga.ThresholdRange
	thresholdRange.Inside = true
	thresholdRange.Start = math.Inf(-1)
	thresholdRange.End = math.Inf(1)

	if strings.HasPrefix(thresholdDef, ",") {
		return thresholdRange, errors.New("empty metric")
	}

	rangeDef := thresholdDef
	// get metric name if given
	if strings.Contains(rangeDef, ",") {
		s1 := strings.FieldsFunc(rangeDef, func(r rune) bool { return r == ',' })
		if len(s1) != 2 {
			return thresholdRange, errors.New("invalid metric")
		}
		thresholdRange.Metric = s1[0]
		rangeDef = s1[1]
	}

	thresholdRange.Definition = rangeDef
	thresholdRange.Inside = !strings.HasPrefix(rangeDef, "@")
	rangeDef = strings.TrimPrefix(rangeDef, "@")

	if rangeDef == "" {
		return thresholdRange, errors.New("empty range")
	}

	if !strings.Contains(rangeDef, ":") {
		// A single number will be interpreted as end with a start of 0
		thresholdRange.Start = 0
		thresholdRange.End, err = stringToFloat(rangeDef)

		if err != nil {
			return thresholdRange, err
		}
		return validateThreshold(thresholdRange)
	}

	s := strings.FieldsFunc(rangeDef, func(r rune) bool { return r == ':' })

	if len(s) == 0 {
		return thresholdRange, errors.New("empty range")
	}

	if len(s) > 2 {
		return thresholdRange, errors.New("invalid range: too many values")
	}

	if len(s) == 1 {
		if strings.HasSuffix(rangeDef, ":") {
			thresholdRange.Start, err = stringToFloat(s[0])
			if err != nil {
				return thresholdRange, err
			}
		} else {
			thresholdRange.End, err = stringToFloat(s[0])
			if err != nil {
				return thresholdRange, err
			}
		}
	} else {
		thresholdRange.Start, err = stringToFloat(s[0])
		if err != nil {
			return thresholdRange, err
		}
		thresholdRange.End, err = stringToFloat(s[1])
		if err != nil {
			return thresholdRange, err
		}
	}

	return validateThreshold(thresholdRange)
}

func stringToFloat(value string) (float64, error) {
	if value == "" {
		return 0, errors.New("empty value")
	}

	if value == "~" {
		return math.Inf(-1), nil
	}

	result, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func validateThreshold(thresholdRange icinga.ThresholdRange) (icinga.ThresholdRange, error) {
	if thresholdRange.Start > thresholdRange.End {
		return thresholdRange, errors.New("invalid range: start greater than end")
	}
	return thresholdRange, nil
}

/*
 * Evaluate a given value against the threshold ranges
 */
func Evaluate(warningList []icinga.ThresholdRange, criticalList []icinga.ThresholdRange,
	value float64, perfData *perfdata.PerformanceData) icinga.ExitCode {

	thresholdWarning := getThreshold(warningList, perfData)
	thresholdCritical := getThreshold(criticalList, perfData)

	if perfData != nil {
		if thresholdCritical != nil {
			perfData.SetCritical(thresholdCritical.Definition)
		}
		if thresholdWarning != nil {
			perfData.SetWarning(thresholdWarning.Definition)
		}
	}

	if thresholdCritical != nil {
		if isValueOutOfRange(*thresholdCritical, value) {
			return icinga.ExitCritical
		}
	}

	if thresholdWarning != nil {
		if isValueOutOfRange(*thresholdWarning, value) {
			return icinga.ExitWarning
		}
	}

	return icinga.ExitOk
}

// Get threshold
func getThreshold(list []icinga.ThresholdRange, perfData *perfdata.PerformanceData) *icinga.ThresholdRange {
	metric := ""
	if perfData != nil {
		metric = perfData.Label
	}

	for _, l := range list {
		if l.Metric == metric {
			return &l
		}
	}
	return nil
}

func isValueOutOfRange(thresholdRange icinga.ThresholdRange, value float64) bool {
	if thresholdRange.Inside {
		// normally value is inside range
		if value < thresholdRange.Start || value > thresholdRange.End {
			return true
		}
		return false
	}

	// normally value is outside of range
	if value >= thresholdRange.Start && value <= thresholdRange.End {
		return true
	}
	return false
}

