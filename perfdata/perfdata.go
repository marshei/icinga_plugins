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
package perfdata

import (
	"fmt"
)

// Performance data
type PerformanceData struct {
	Label    string
	Value    float64
	UOM      string
	Warning  string
	Critical string
	Minimum  string
	Maximum  string
}

// CreatePerformanceData creates and returns a new PerformanceData object
func CreatePerformanceData(Label string, Value float64, UOM string) *PerformanceData {
	pd := new(PerformanceData)
	pd.Label = Label
	pd.Value = Value
	pd.UOM = UOM

	return pd
}

// SetWarning sets the warning value of a PerformanceData object
func (pd *PerformanceData) SetWarning(warning string) {
	pd.Warning = warning
}

// SetCritical sets the critical value of a PerformanceData object
func (pd *PerformanceData) SetCritical(critical string) {
	pd.Critical = critical
}

// SetMinimum sets the minimum value of a PerformanceData object
func (pd *PerformanceData) SetMinimum(minimum string) {
	pd.Minimum = minimum
}

// SetMaximum sets the maximum value of a PerformanceData object
func (pd *PerformanceData) SetMaximum(maximum string) {
	pd.Maximum = maximum
}

// String returns a PerformanceData object as formatted string without a separator
func (pd *PerformanceData) String() string {
	return fmt.Sprintf("'%s'=%s%s;%s;%s;%s;%s", pd.Label, getValueString(pd.Value),
		pd.UOM, pd.Warning, pd.Critical, pd.Minimum, pd.Maximum)
}

func getValueString(value float64) string {
	if value == float64(int64(value)) {
		return fmt.Sprintf("%d", int64(value))
	}
	return fmt.Sprintf("%f", value)
}
