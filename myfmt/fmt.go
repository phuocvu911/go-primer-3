package myfmt

import "goprimer/mystrconv"

func Sprint(v any) string {
	if v == nil {
		return "<nil>"
	}

	switch data := v.(type) {
	case error:
		return data.Error()
	case string:
		return data
	case bool:
		if data == true {
			return "true"
		} else {
			return "false"
		}
	case int:
		str, _ := mystrconv.FormatInt(int64(data), 10)
		return str
	case int64:
		str, _ := mystrconv.FormatInt(data, 10)
		return str
	case uint64:
		str, _ := mystrconv.FormatUint(data, 10)
		return str
	case []byte:
		return string(data)
	default:
		return "%!(unsupported type)"
	}
}
