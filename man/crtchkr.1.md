% crtchkr 1

## Name
crtchkr - check and validate X.509 certificate

## SYNOPSIS
**crtchkr** \[global options\] command \[command options\] \[arguments...\]

## DESCRIPTION

A program to check and validate X.509 certificates. Checks chain of trust from root certificate
down to the lowest one validating each one on the way.

In case one of them fails an action (**[crtchkr-message(1)](crtchkr-message.1.md)**) is sent,
which has to be provided and configured in **crtchkr.toml**

## COMMANDS

| **Command**        | **Description**                                      |
|--------------------|------------------------------------------------------|
| crtchkr-print(1)   | Print X.509 certificate information.                 |
| crtchkr-check(1)   | Checks validity of X.509 certificate.                |
| crtchkr-message(1) | Executes message action defined in config-toml(5).   |
| crtchkr-help       | Shows full list of commands or help for one command. |

## GLOBAL OPTIONS


### **--config**=*path*, **-c**=*path* 
Load configuration from provided *path*

### **--verson**, **-v**
Show version of **crtchkr**

## SEE ALSO
**[crtchkr-print(1)](crtchkr-print.1.md)**, **[crtchkr-check(1)](crtchkr-check.1.md)**, 
**[crtchkr-message(1)](crtchkr-message.1.md)**, **[crtchkr.toml(5)](crtchkr.toml.5.md)**

## HISTORY
2 Feb 2023, Manpage created by **[Chris \<https://github.com/chramb\>](https://github.com/chramb)**
