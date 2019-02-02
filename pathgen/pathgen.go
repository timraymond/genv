package pathgen

import (
	"os/exec"
	"path"

	"github.com/gobuffalo/genny"
)

func New(project, username string) *genny.Generator {
	g := genny.New()

	g.Command(exec.Command("mkdir", project))
	g.Command(exec.Command("mkdir", project+"/bin"))
	g.Command(exec.Command("mkdir", project+"/pkg"))
	g.Command(exec.Command("mkdir", project+"/src"))

	projPath := path.Join(project, "src", "github.com", username, project)
	g.Command(exec.Command("mkdir", "-p", projPath))

	return g
}
