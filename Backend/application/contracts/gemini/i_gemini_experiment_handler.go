package contracts


// IExperimentHandler defines the contract for handling experiments in Gemini.
type IExperimentHandler interface {
	GetExperimentsForChapter(chapterName string) ([]string, error)
	GetExperimentContent(experimentName string) ([]byte, error)
}