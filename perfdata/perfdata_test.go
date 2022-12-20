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

import "testing"

func TestCreatePerformanceData(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")
	want := "'testing'=123.430000;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "s")
	want = "'testing'=123.430000s;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "us")
	want = "'testing'=123.430000us;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "ms")
	want = "'testing'=123.430000ms;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 99.43, "%")
	want = "'testing'=99.430000%;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "B")
	want = "'testing'=123.430000B;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "KB")
	want = "'testing'=123.430000KB;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "MB")
	want = "'testing'=123.430000MB;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "TB")
	want = "'testing'=123.430000TB;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}

	pd = CreatePerformanceData("testing", 123.43, "c")
	want = "'testing'=123.430000c;;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}
}

func TestSetWarning(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetWarning("62")

	want := "'testing'=123.430000;62;;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}
}

func TestSetCritical(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetCritical("62")

	want := "'testing'=123.430000;;62;;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}
}

func TestSetMinimum(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetMinimum("62")

	want := "'testing'=123.430000;;;62;"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}
}

func TestSetMaximum(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.43, "")

	pd.SetMaximum("62")

	want := "'testing'=123.430000;;;;62"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}
}

func TestString(t *testing.T) {
	pd := CreatePerformanceData("testing", 123.00, "s")

	pd.SetWarning("48")
	pd.SetCritical("55")
	pd.SetMinimum("13")
	pd.SetMaximum("875")

	want := "'testing'=123s;48;55;13;875"

	if pd.String() != want {
		t.Errorf("CreatePerformanceData was incorrect, got: %s, want: %s.", pd.String(), want)
	}
}
