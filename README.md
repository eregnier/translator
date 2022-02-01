# Translator

Simple translation tool using google translation api.

To use it you have to provide a valid service account as json file with path in the environment. This service account must be allowed to use GCP translations api.

usage:

```bash
GOOGLE_APPLICATION_CREDENTIALS='/tmp/cred.json' ./translator input.txt output.txt
```

input file should be like
```
house
wife
computer
```

so the output will be
```
loger
épouse
ordinateur
```

ℹ️ at the moment the program only translates from english to french. If you want extra language options feel free to open issue or create PR.


It is portable and works for windows 🪟 and linux 🐧. See releases for prebuild binaries.
