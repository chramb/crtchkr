% crtchkr-check 1

## Name
crtchkr-cert - check and validate X.509 certificate

## SYNOPSIS
**crtchkr (check|c|verify|v)** \[options\] \[arguments...\]

## DESCRIPTION
Verify validity of provided certificate by comparing it 
to ones stored in trust system's store and going down 
the chain of trust 

## OPTIONS

### **--oneline**, **-1** 
Write validity of certificate on single line

### **--on-fail**=*action*, **-x**=*action*
Specifies action to perform in case certificate is found invalid.
Uses action defined in **[crtchkr.toml(5)](crtchkr.toml.5.md)**


## EXAMPLES
Check validity of certificate of *https://pwr.edu.pl* and print results in one line,
in case it fails write a mail specified in **crtchkr.toml** as \[*mail.admins*\].

    $ crtchkr check -1 --on-fail mail.adminis https://pwr.edu.pl

Check validity of certificate *cert.pem* located in /var/www/ssl/ if invalid post a webhook to Discord server,
specified in **crtchkr.toml** as \[*discord.warncert*\].

    $ crtchrk c --on-fail discord.warncert /var/www/ssl/cert.pem 

## SEE ALSO
**[crtchkr(1)](crtchkr.1.md)**, **[crtchkr-print(1)](crtchkr-print.1.md)**, 
**[crtchkr-message(1)](crtchkr-message.1.md)**, **[crtchkr.toml(5)](crtchkr.toml.5.md)**

## HISTORY
2 Feb 2023, Manpage created by **[Chris \<https://github.com/chramb\>](https://github.com/chramb)**
