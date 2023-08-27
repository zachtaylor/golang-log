package log

// LegacySourceFormatter uses detail with path prefix removal
func LegacySourceFormatter(gopaths ...string) SourceFormatter {
	return RestringSourceFormatter(
		DefaultSourceFormatter(),
		RestringerCutPrefixes(gopaths),
	)
}
