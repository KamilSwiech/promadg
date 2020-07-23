# prometheus-alerts-doc-gen

Generate prometheus rules and alerts into doc with markdown format

> CREATION IN PROGRESS

### TODO
- tests
- markdown config in tpl
- full api support

### How to use
The quickiest way to get output:
```
promadg get -data rules > output.json
```

### Commands
* set:
    * context -- set config file with prometheus
    * confing -- set config format file for parsing jsons
    * time -- set time for queries. Defaults to "now"
* get:
    * data -- get selected data (rules, alerts)
    * url -- get provided url
    * time -- get configured time
