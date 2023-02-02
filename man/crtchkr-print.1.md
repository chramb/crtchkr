% crtchkr-print 1

## Name
crtchkr-print - print certificate info

## SYNOPSIS
**crtchkr (print|p)** \[options\] \[arguments...\]

## OPTIONS

### **--dnsnames**
Print DNSNames specified in certificate (default: false)

### **--subject**
Print Subject specified in certificate (default: false)

### **--issuer**
Print Issuer specified in certificate (default: false)

### **--version**
Print Certificate Version specified in certificate (default: false)

### **--not-before**
print date not before which certificate is valid (default: false)

### **--not-after**
print date when certificate expires (default: false)

### **--all**, **-a**
print all (default: false)

## EXAMPLES
Print *all* information on certificate called *cert.pem*

    $ crtchkr print -a cert.pem

Print *DNSNames* information on certificate used by *https://example.com*.\

    $ crtchkr print --dnsnames https://example.com
**Note**: make sure to specify *https://* protocol to let **crtchkr** know that it has to fetch certificates.

## SEE ALSO
**[crtchkr(1)](crtchkr.1.md)**, **[crtchkr-check(1)](crtchkr-check.1.md)**, 
**[crtchkr-message(1)](crtchkr-message.1.md)**, **[crtchkr.toml(5)](crtchkr.toml.5.md)**

## HISTORY
2 Feb 2023, Manpage created by **[Chris \<https://github.com/chramb\>](https://github.com/chramb)**
