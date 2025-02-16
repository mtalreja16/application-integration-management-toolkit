## integrationcli authconfigs create

Create an authconfig

### Synopsis

Create an authconfig

```
integrationcli authconfigs create [flags]
```

### Options

```
  -e, --encrypted-file string     Base64 encoded, Cloud KMS encrypted Auth Config JSON file path
  -k, --encryption-keyid string   Cloud KMS key for decrypting Auth Config; Format = locations/*keyRings/*/cryptoKeys/*
  -f, --file string               Auth Config JSON file path
  -h, --help                      help for create
```

### Options inherited from parent commands

```
  -a, --account string       Path Service Account private key in JSON
      --apigee-integration   Use Apigee Integration; default is false (Application Integration)
      --disable-check        Disable check for newer versions
      --no-output            Disable printing API responses from the control plane
  -p, --proj string          Integration GCP Project name
  -r, --reg string           Integration region name
  -t, --token string         Google OAuth Token
      --verbose              Enable verbose output from integrationcli
```

### SEE ALSO

* [integrationcli authconfigs](integrationcli_authconfigs.md)	 - Manage integration auth configurations

###### Auto generated by spf13/cobra on 17-Mar-2023
