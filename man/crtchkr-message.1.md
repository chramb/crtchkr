% crtchkr-message 1

## Name
crtchkr-message - execute the message action

## SYNOPSIS
**crtchkr** (-c|--config)=*path* **(message|m)** \[options\] \[action\]

## DESCRIPTION
Verify validity of provided certificate by comparing it
to ones stored in trust system's store and going down
the chain of trust

## OPTIONS

### **--discord**=*action*, **-d**=*action*
Post a webhook to server specified in **[crtchkr.toml(5)](crtchkr.toml.5.md)**

### **--mail**=*action*, **-x**=*action*
Send mail through smtp server specified in **[crtchkr.toml(5)](crtchkr.toml.5.md)**

## EXAMPLES
Send message to discord server specified in **crtchkr.toml** as \[discord.*variable*\]

    $ crtchkr --config crtchkr.toml message --discord variable

Send mail with parameters specified in **config.toml** in table \[mail.*hello*\]

    $ crtchkr -c /etc/crtchkr/config.toml m -m hello

## SEE ALSO
**[crtchkr(1)](crtchkr.1.md)**, **[crtchkr-print(1)](crtchkr-print.1.md)**, 
**[crtchkr-check(1)](crtchkr-check.1.md)**, **[crtchkr.toml(5)](crtchkr.toml.5.md)**

## HISTORY
2 Feb 2023, Manpage created by **[Chris \<https://github.com/chramb\>](https://github.com/chramb)**
