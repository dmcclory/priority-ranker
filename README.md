# ranker
ranker is a tool for finding the priority in a list of options. You can create and switch between lists with the 'list init' and 'list subcommands'

'ranker vote' - You find the relative priority by voting on randomly drawn pairs of options.

'ranker results' - Once you've voted a number of times, you can see a ranking of the options.

The rankings are calculated using the [Bradley Terry model](https://en.wikipedia.org/wiki/Bradley–Terry_model).

## install

You can either run the install script or grab a binary from the [latest release](https://github.com/dmcclory/priority-ranker/releases/latest)

`gh` is required for the install script.

```
gh api -H 'Accept: application/vnd.github.v3.raw' "repos/dmcclory/priority-ranker/contents/install.sh" | bash
```
