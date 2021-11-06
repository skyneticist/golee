# golee
[![Go Report Card](https://goreportcard.com/badge/github.com/skyneticist/golee)](https://goreportcard.com/report/github.com/skyneticist/golee)

## cli tool for git productivity

This cli tool is intended to make frequent git work in the command line more convenient and less time-consuming.

The main goals of this application are:

  - To ~significantly reduce the time typing when working in git cli, without relying on hooks 
  - Note: (since I work on a team for a big company, it's preferred to have some versatility here)

  - To be as lightweight and efficient/performant as possible
  - To be incredibly portable
  - To be easily updated
  - Easy to work on
  - Easy to release
 
 
 
#Installing
-----------

One thing that IS nice about this tool is how easy it is to install and to keep updated. Also, the size of the executable and the way its being added to PATH is fast, safe, convenient.

To achieve all of this, Goreleaser and a custom postinstall scrit is used in combination with a package.json file. This essentially allows for a platform-independent way to install the git-golee tool with a simple npm command.

```npm i -g @skyneticist/git-golee@latest```

That's it. Now you have access to git-golee via `net-helper`

Currently working on a way to set the command to `gg` instead of `net-helper` without setting an alias

