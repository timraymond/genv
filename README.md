genv
====

genv is a utility for generating an isolated $GOPATH for new and existing Go
projects. It uses `direnv` to manage adjusting paths and $GOPATH.

Creating a new project
----------------------

```
$ genv new <name of project>
```

Cloning an existing project
---------------------------

```
$ genv clone <git URI>
```

Changing Go versions
--------------------

```
$ genv switch 1.12
```
