package common

func StringVal(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func IntVal(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func FindMinValInt(values []int) float64 {
	var min int = values[0]
	for _, val := range values {
		if min > val {
			min = val
		}
	}
	return float64(min)
}

func FindMaxValueInt(values []int) float64 {
	var max int = values[0]
	for _, val := range values {
		if max < val {
			max = val
		}
	}
	return float64(max)
}

func FindAvgValueInt(values []int) float64 {
	sum := 0
	for _, val := range values {
		sum += val
	}

	return float64(sum / len(values))
}

func FindMedianValueInt(values []int) float64 {
	l := len(values)
	if l == 0 {
		return 0
	}
	if l%2 == 0 {
		return float64(values[l/2-1]+values[l/2]) / 2.0
	} else {
		return float64(values[l/2])
	}
	return 0
}
