package npm

import (
	"os"
	"path/filepath"

	"golang.org/x/xerrors"

	"github.com/aquasecurity/fanal/analyzer"
	"github.com/aquasecurity/fanal/analyzer/library"
	"github.com/aquasecurity/fanal/types"
	"github.com/aquasecurity/fanal/utils"
	"github.com/aquasecurity/go-dep-parser/pkg/npm"
)

//func init() {
//	analyzer.RegisterAnalyzer(&npmLibraryAnalyzer{})
//}

const version = 1

var requiredFiles = []string{"package-lock.json"}

type npmLibraryAnalyzer struct{}

func (a npmLibraryAnalyzer) Analyze(target analyzer.AnalysisTarget) (*analyzer.AnalysisResult, error) {
	res, err := library.Analyze(types.Npm, target.FilePath, target.Content, npm.Parse)
	if err != nil {
		return nil, xerrors.Errorf("unable to parse package-lock.json: %w", err)
	}
	return res, nil
}

func (a npmLibraryAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	fileName := filepath.Base(filePath)
	return utils.StringInSlice(fileName, requiredFiles)
}

func (a npmLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypeNpm
}

func (a npmLibraryAnalyzer) Version() int {
	return version
}
