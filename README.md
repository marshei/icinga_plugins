# Icinga Plugins Support

![CI](https://github.com/marshei/icinga_plugins/workflows/CI/badge.svg?branch=main)

This repository contains source code written in Go to parse and
evaluate Icinga thresholds or threshold ranges as defined [here](https://icinga.com/docs/icinga-2/latest/doc/05-service-monitoring/#threshold-ranges)

The definition has been extended to allow a plugin to accept more than one threshold definitions for different performance metrics.
Then the typical format:

```-w 10:20```

can be extended for use with more metrics e.g.

```-c metric1,10:20;metric2,@30:40```

If no metric is given only one threshold will be accepted.

**Note**
Depending on the paramter handling escaping might be required.

## Example

The moethod `ParseThresholdList` parses the provided string from the CLI into a list of threshold ranges.
```
func ParseThresholdList(thresholdDef string) ([]icinga.ThresholdRange, error)
```

These list will then be used to evaluate a given value with or without metric
```
func Evaluate(warningList []icinga.ThresholdRange, 
              criticalList []icinga.ThresholdRange,
	          value float64, perfData *perfdata.PerformanceData) icinga.ExitCode
```

The returned plugin exit code can finally be printed along with a message
```
func Print(message string, code ExitCode) ExitCode

func PrintWithPerformanceData(message string, code ExitCode,
                              perfDataList []perfdata.PerformanceData) ExitCode
```

The exit code of the plugin should be `int(exitCode)`, e.g.
```
func exit(code icinga.ExitCode) {
	os.Exit(int(code))
}
```
