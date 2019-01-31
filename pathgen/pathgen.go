package pathgen

import (
	"os/exec"

	"github.com/gobuffalo/genny"
)

func New(path string) *genny.Generator {
	g := genny.New()

	g.Command(exec.Command("mkdir", path))
	g.Command(exec.Command("mkdir", path+"/bin"))
	g.Command(exec.Command("mkdir", path+"/pkg"))
	g.Command(exec.Command("mkdir", path+"/src"))

	return g
}
