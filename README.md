# prometheus-alerts-doc-gen

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
