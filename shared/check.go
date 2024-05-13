package shared

func CheckMeta(meta map[string]interface{}) []string {
	warnings := []string{}
	warnings = append(warnings, hasRequiredTags(meta)...)
	warnings = append(warnings, hasRedundantTags(meta)...)
	return warnings
}
