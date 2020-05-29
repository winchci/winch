package winch

type CommitType string

const (
	breakingChangeTitle = "BREAKING CHANGES"
	revertedTitle       = "Reverted"
	featuresTitle       = "Features"
	fixesTitle          = "Fixes"
	refactoringTitle    = "Refactoring"
	testsTitle          = "Tests"
	stylesTitle         = "Style"
	documentationTitle  = "Documentation"
	performanceTitle    = "Performance"
	changesTitle        = "Changes"
)

const (
	buildTypeKey    = "build"
	choreTypeKey    = "chore"
	ciTypeKey       = "ci"
	docsTypeKey     = "docs"
	fixTypeKey      = "fix"
	perfTypeKey     = "perf"
	refactorTypeKey = "refactor"
	revertTypeKey   = "revert"
	styleTypeKey    = "style"
	testTypeKey     = "test"
	featTypeKey     = "feat"
	changeTypeKey   = "change"
)

var typeDescriptions = map[CommitType]string{
	buildTypeKey:    "Build system changes",
	choreTypeKey:    "Chore, tidying, adding setup, non-style changes",
	ciTypeKey:       "Continuous Integration changes",
	docsTypeKey:     "Documentation changes",
	fixTypeKey:      "Fixing a bug",
	perfTypeKey:     "Performance changes",
	refactorTypeKey: "Refactoring",
	revertTypeKey:   "Reverting a prior change (reference the commit)",
	styleTypeKey:    "Style changes (linting, formatting)",
	testTypeKey:     "Tests added, changed or removed",
	featTypeKey:     "Feature implementation",
}

var knownTypes = map[CommitType]bool{
	buildTypeKey:    false,
	choreTypeKey:    false,
	ciTypeKey:       false,
	docsTypeKey:     false,
	fixTypeKey:      false,
	perfTypeKey:     false,
	refactorTypeKey: false,
	revertTypeKey:   false,
	styleTypeKey:    false,
	testTypeKey:     false,
	featTypeKey:     true,
}

var typeTitles = map[CommitType]string{
	buildTypeKey:    changesTitle,
	choreTypeKey:    changesTitle,
	ciTypeKey:       changesTitle,
	docsTypeKey:     documentationTitle,
	fixTypeKey:      fixesTitle,
	perfTypeKey:     performanceTitle,
	refactorTypeKey: refactoringTitle,
	revertTypeKey:   revertedTitle,
	styleTypeKey:    stylesTitle,
	testTypeKey:     testsTitle,
	featTypeKey:     featuresTitle,
	changeTypeKey:   changesTitle,
}

var titleOrder = []string{
	breakingChangeTitle,
	changesTitle,
	revertedTitle,
	featuresTitle,
	fixesTitle,
	refactoringTitle,
	testsTitle,
	stylesTitle,
	documentationTitle,
	performanceTitle,
}

func (t CommitType) String() string {
	return string(t)
}

func (t CommitType) Title() string {
	return typeTitles[t]
}

func (t CommitType) Description() string {
	return typeDescriptions[t]
}

func (t CommitType) IsPatch() bool {
	return !knownTypes[t]
}

func (t CommitType) IsMinor() bool {
	return knownTypes[t]
}

func NewType(s string) CommitType {
	if _, ok := knownTypes[CommitType(s)]; ok {
		return CommitType(s)
	}

	return CommitType("change")
}

func GetCommitTypes() []CommitType {
	t := make([]CommitType, 0, len(knownTypes))
	for k := range knownTypes {
		t = append(t, k)
	}
	return t
}
