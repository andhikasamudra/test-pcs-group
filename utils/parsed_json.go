package utils

func ParsedInterfaceToMapInterface(data interface{}) map[string]string {
	if data == nil {
		return nil
	}
	var res = map[string]string{}
	for k, v := range data.(map[string]interface{}) {
		res[k] = v.(string)
	}

	return res
}
