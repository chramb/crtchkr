% crtchkr.toml 5

## Name
crtchkr.toml - Syntax of configuration file for **[crtchkr(1)](crtchkr.1.md)**

## Format
The [TOML format][toml] is used as the encoding of the configuration file.  Every option and subtable listed
here  is nested under a global "storage" table.  No bare options are used. The format of TOML can be simpli‐
fied to:

    [table]
    option = value

    [table.subtable1]
    option = value

    [table.subtable2]

## Usage

**crtchkr.toml** uses 2 tables \[mail\] and \[discord\]
with specific options to specify address and parameters of messages to send.

For setting variables inside those parameters **[crtchkr(1)](crtchkr.1.md)** uses Go [Templates](https://pkg.go.dev/text/template).

The exposed type is FailedCerts, which is an array of struct FailedCert which contains:

- [\*x509.Certificate - https://pkg.go.dev/crypto/x509#Certificate](https://pkg.go.dev/crypto/x509#Certificate),
- Err: error of why it is invalid. You can use all those variables and
parameters of structs in them inside your messages.

### FailedCerts definition in source code.

    type FailedCerts []FailedCert

    type FailedCert struct {
        Cert *x509.Certificate
        Err  error
    }


## Examples

### Mail Table

    [mail.hello]
    server = "smtp.example.com:465"
    from = "certchecker@example.com"
    to = ["admin@example.com", "notifications@example.com"]
    message = """
    Subject: Test Email
    
    
    Your Certificate is no longer valid:
    {{ range . }}
    ---
    CN:\t{{ .Cert.Subject.CommonName}}
    Error:\t{{ .Err }}
    ---
    {{ end }}
    """
    [mail.pwr.auth]
    identity = ""
    username = "certchecker@example.com"
    password = "hunter2"
    host = "smtp.example.com"

### Discord Table
    
    [discord.exampletable]
    url = "https://discord.com/api/webhooks/<server>/<endpoint>/"
    request = """
    {
    "content": "Your Certificates are no longer valid!!\\n found those issues: {{ range . }} \\nCN: {{ .Cert.Subject.CommonName }} -- Error: {{ .Err }}{{ end }}",
    "username": "Certificate Tester"
    }
    """


## SEE ALSO
**[crtchkr-print(1)](crtchkr-print.1.md)**, **[crtchkr-check(1)](crtchkr-check.1.md)**, 
**[crtchkr-message(1)](crtchkr-message.1.md)**, **[crtchkr.toml(5)](crtchkr.toml.5.md)**

## HISTORY
2 Feb 2023, Manpage created by **[Chris \<https://github.com/chramb\>](https://github.com/chramb)**
