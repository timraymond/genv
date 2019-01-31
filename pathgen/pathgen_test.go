package pathgen_test

import (
	"testing"

	"github.com/gobuffalo/genny/gentest"
	"github.com/timraymond/genv/pathgen"
)

func Test_CreatesGopathDirs(t *testing.T) {
	res, err := gentest.RunNew(pathgen.New("testpath"), nil)
	if err != nil {
		t.Fatal("Unexpected err while running pathgen: err:", err)
	}

	expected := []string{
		"mkdir testpath",
		"mkdir testpath/bin",
		"mkdir testpath/pkg",
		"mkdir testpath/src",
	}

	err = gentest.CompareCommands(expected, res.Commands)
	if err != nil {
		t.Fatal("Unexpected err comparing executed commands: err:", err)
	}
}
