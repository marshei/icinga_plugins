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
package icinga

import (
	"encoding/json"
	"fmt"
)

type ThresholdRange struct {
	Definition string
	Metric     string
	Inside     bool
	Start      float64
	End        float64
}

// To print a struct can be represented as JSON
// Please note that float turned into string to handle the Inf values
func (tr *ThresholdRange) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%f", tr.Start)
	e := fmt.Sprintf("%f", tr.End)
	return json.Marshal(&struct {
		Definition string `json:"definition"`
		Metric     string `json:"metric"`
		Inside     bool   `json:"inside"`
		Start      string `json:"start"`
		End        string `json:"end"`
	}{
		Definition: tr.Definition,
		Metric:     tr.Metric,
		Inside:     tr.Inside,
		Start:      s,
		End:        e,
	})
}

// Plugin exit codes
type ExitCode int

const (
	ExitOk       ExitCode = 0
	ExitWarning  ExitCode = 1
	ExitCritical ExitCode = 2
	ExitUnknown  ExitCode = 3
)

func (code ExitCode) String() string {
	switch code {
	case ExitOk:
		return "OK"
	case ExitWarning:
		return "WARNING"
	case ExitCritical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// Get the resulting code from two codes
func (code ExitCode) GetResultCode(other ExitCode) ExitCode {
	switch code {
	case ExitOk:
		return other
	case ExitWarning:
		if other == ExitCritical {
			return other
		}
		return ExitWarning
	case ExitCritical:
		return ExitCritical
	default:
		if other == ExitCritical || other == ExitWarning {
			return other
		}
		return ExitUnknown
	}
}
