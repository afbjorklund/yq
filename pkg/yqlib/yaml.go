package yqlib

type YamlPreferences struct {
	LeadingContentPreProcessing bool
	PrintDocSeparators          bool
	UnwrapScalar                bool
	EvaluateTogether            bool
	CompactSequenceIndent       bool
}

func NewDefaultYamlPreferences() YamlPreferences {
	return YamlPreferences{
		LeadingContentPreProcessing: true,
		PrintDocSeparators:          true,
		UnwrapScalar:                true,
		EvaluateTogether:            false,
		CompactSequenceIndent:       false,
	}
}

var ConfiguredYamlPreferences = NewDefaultYamlPreferences()
