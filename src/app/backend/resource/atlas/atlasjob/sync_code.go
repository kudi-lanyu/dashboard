package atlasjob

import (
	"fmt"
	"log"
	"strings"
)

type (
	SyncMode           = string
	SyncSource         = string
	SyncImage          = string
	SyncGitProjectName = string
)

type SyncCodeArgs struct {
	SyncMode           SyncMode           `yaml:"syncMode";json:"syncMode"`                               // --syncMode: rsync, hdfs, git
	SyncSource         SyncSource         `yaml:"syncSource";json:"syncSource"`                           // --syncSource
	SyncImage          SyncImage          `yaml:"syncImage,omitempty";json:"syncImage"`                   // --syncImage
	SyncGitProjectName SyncGitProjectName `yaml:"syncGitProjectName,omitempty";json:"syncGitProjectName"` // --syncImage
}

func (sc *SyncCodeArgs) HandleSyncCode() error {

	switch sc.SyncMode {
	case "":
		log.Println("No action for sync Code")
	case "git":
		log.Println("Check and prepare sync code with git")
		if sc.SyncSource == "" {
			return fmt.Errorf("--syncSource should be set when syncMode is set")
		}

		// split test.git to test

		parts := strings.Split(strings.Trim(sc.SyncSource, "/"), "/")
		sc.SyncGitProjectName = strings.Split(parts[len(parts)-1], ".git")[0]
		log.Println("Try to split %s to get project name %s", sc.SyncSource, sc.SyncGitProjectName)
	case "rsync":
		log.Println("Check and prepare sync code with rsync")
		if sc.SyncSource == "" {
			return fmt.Errorf("--syncSource should be set when syncMode is set")
		}

	default:
		log.Fatalf("Unknown sync mode: %s", sc.SyncMode)
		return fmt.Errorf("Unknown sync mode: %s, it should be git or rsync", sc.SyncMode)
	}

	return nil
}
