package relevantreader

import "github.com/Ekram-B2/suggestionsmanager/results"

// RelevantReader supports reading relevant data from a persistant store. Relevant data is
// partial segment of the global data set with which a rank can be attributed
type RelevantReader interface {

	// ReadRelevant used to read in relevant data from a persistant store
	ReadRelevant(string) (results.Results, error)
}