# Prometheus Alerts Documentation Generator

Generate prometheus rules and alerts into doc with markdown format or your
own custom format.

## TODO
- tests
- full api support
## Installation
Download binary:
```
curl -sSL 'https://github.com/kamilswiec/promadg/releases/download/0.0.2/promadg' -o /usr/local/bin/promadg
chmod +x /usr/local/bin/promadg
```
## How to use
The quickiest way to get output:
```
promadg -p your-prometheus-name.com > output.md
```
Results in:

# EXAMPLE GROUP
## NotUp
### Query:
up < 1
### Duration: 5m
### Labels:
>  severity: critical
### Annotations:
>  description: Prometheus is down


### Custom template
You can provide your own template for Sprig with `-t` flag. For example:
```
promadg -p your-prometheus-name.com -t /tmp/my-template > output.md
```
