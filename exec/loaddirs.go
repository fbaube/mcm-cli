package exec

import(
	"github.com/fbaube/mcfile"
	FU "github.com/fbaube/fileutils"
	SU "github.com/fbaube/stringutils"
	L "github.com/fbaube/mlog"
)

func LoadDirpathsContentFSs(ff []FU.FSItem) ([]*mcfile.ContentityFS) {
     if ff == nil || len(ff) == 0 {
     	return nil
	}
     var pFSs []*mcfile.ContentityFS
     var pFS   *mcfile.ContentityFS

     // For every input FSItem
     for iDir, pDir := range ff {
     	 var shortName = SU.EnsureTrailingPathSep(
	     SU.Tildotted(pDir.FPs.AbsFP.S()))
	 L.L.Info("InDir[%d]: %s", iDir, shortName)
	 pFS = mcfile.NewContentityFS(pDir.FPs.AbsFP.S(), nil)
	 L.L.Okay("Found %d item(s) total (%d dirs, %d files)",
	 	pFS.ItemCount(), pFS.DirCount(), pFS.FileCount())
	 if pFS.FileCount() == 0 {
	    	L.L.Info("No content inputs to process: " + shortName)
		continue
	 }
	 pFSs = append(pFSs, pFS) 
     }
     return pFSs
}