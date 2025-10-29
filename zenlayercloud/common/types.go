package common

func Integer(v int) *int {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Uint(v uint) *uint {
	return &v
}

func Uint64(v uint64) *uint64 {
	return &v
}

func Float64(v float64) *float64 {
	return &v
}

func Bool(v bool) *bool {
	return &v
}

func String(v string) *string {
	return &v
}

func ToInteger(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func ToInt64(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func ToUint(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

func ToUint64(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func ToFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func ToBool(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func ToString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}
