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
package icinga

import (
	"fmt"

	"github.com/marshei/icinga_plugins/perfdata"
)

func Print(message string, code ExitCode) ExitCode {
	fmt.Printf("%s - %s\n", code.String(), message)
	return code
}

func PrintWithPerformanceData(message string, code ExitCode, perfDataList []perfdata.PerformanceData) ExitCode {
	if len(perfDataList) == 0 {
		return Print(message, code)
	}
	perfData := ""
	for _, pd := range perfDataList {
		if perfData == "" {
			perfData = pd.String()
		} else {
			perfData += " " + pd.String()
		}
	}
	fmt.Printf("%s - %s | %s\n", code.String(), message, perfData)
	return code
}

func PrintUnknown(message string) ExitCode {
	return Print(message, ExitUnknown)
}

func PrintWarning(message string) ExitCode {
	return Print(message, ExitWarning)
}

func PrintWarningWithPerformanceData(message string, perfData *perfdata.PerformanceData) ExitCode {
	if perfData == nil {
		return PrintWarning(message)
	}
	return PrintWithPerformanceData(message, ExitWarning, []perfdata.PerformanceData{*perfData})
}

func PrintCritical(message string) ExitCode {
	return Print(message, ExitCritical)
}

func PrintCriticalWithPerformanceData(message string, perfData *perfdata.PerformanceData) ExitCode {
	if perfData == nil {
		return PrintCritical(message)
	}
	return PrintWithPerformanceData(message, ExitCritical, []perfdata.PerformanceData{*perfData})
}

func PrintOk(message string) ExitCode {
	return Print(message, ExitOk)
}

func PrintOkWithPerformanceData(message string, perfData *perfdata.PerformanceData) ExitCode {
	if perfData == nil {
		return PrintOk(message)
	}
	return PrintWithPerformanceData(message, ExitOk, []perfdata.PerformanceData{*perfData})
}
