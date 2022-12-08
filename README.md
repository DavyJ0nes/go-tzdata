# go-tzdata

## Description

This is a quick experiment into seeing why statically built binaries were having
issues trying to get timezone data.

In hindsight it's super obvious that the tzdata can't be referenced because
all the libraries that it needs are included in the binary so doesn't call out
to OS.

Go includes the ability to inject this data at build time from the:
https://pkg.go.dev/time/tzdata package.

There are two ko files that are used as build configuration.

- `.ko-before.yaml` does not include the tzdata file.
- `.ko-after.yaml` includes the tzdata file.

## Example

```
$ cp .ko-before.yaml .ko.yaml
overwrite .ko.yaml? (y/n [n]) y
.ko-before.yaml -> .ko.yaml

$ ko build -L --preserve-import-paths
2022/12/08 19:34:40 Using base gcr.io/distroless/static:nonroot@sha256:ed05c7a5d67d6beebeba19c6b9082a5513d5f9c3e22a883b9dc73ec39ba41c04 for github.com/davyj0nes/go-tzdata
2022/12/08 19:34:41 Using build config test for github.com/davyj0nes/go-tzdata
2022/12/08 19:34:41 Building github.com/davyj0nes/go-tzdata for linux/amd64
2022/12/08 19:34:42 Loading ko.local/github.com/davyj0nes/go-tzdata:82ef31cde205ba8192902878c1db0f5b769eb75d4dba2870ff58c75051e0d2d1
2022/12/08 19:34:42 Loaded ko.local/github.com/davyj0nes/go-tzdata:82ef31cde205ba8192902878c1db0f5b769eb75d4dba2870ff58c75051e0d2d1
2022/12/08 19:34:42 Adding tag latest
2022/12/08 19:34:42 Added tag latest
ko.local/github.com/davyj0nes/go-tzdata:82ef31cde205ba8192902878c1db0f5b769eb75d4dba2870ff58c75051e0d2d1

$ docker run ko.local/github.com/davyj0nes/go-tzdata:82ef31cde205ba8192902878c1db0f5b769eb75d4dba2870ff58c75051e0d2d1
✅ got:  Pacific/Enderbury
💀 err:  unknown time zone Pacific/Kanton

$ cp .ko-after.yaml .ko.yaml
overwrite .ko.yaml? (y/n [n]) y
.ko-after.yaml -> .ko.yaml

$ ko build -L --preserve-import-paths
2022/12/08 19:35:06 Using base gcr.io/distroless/static:nonroot@sha256:ed05c7a5d67d6beebeba19c6b9082a5513d5f9c3e22a883b9dc73ec39ba41c04 for github.com/davyj0nes/go-tzdata
2022/12/08 19:35:07 Using build config test for github.com/davyj0nes/go-tzdata
2022/12/08 19:35:07 Building github.com/davyj0nes/go-tzdata for linux/amd64
2022/12/08 19:35:07 Loading ko.local/github.com/davyj0nes/go-tzdata:d75ad91f74464cd0ac2e0d98a095e150199260facdd50275e018c87711fae55f
2022/12/08 19:35:07 Loaded ko.local/github.com/davyj0nes/go-tzdata:d75ad91f74464cd0ac2e0d98a095e150199260facdd50275e018c87711fae55f
2022/12/08 19:35:07 Adding tag latest
2022/12/08 19:35:07 Added tag latest
ko.local/github.com/davyj0nes/go-tzdata:d75ad91f74464cd0ac2e0d98a095e150199260facdd50275e018c87711fae55f

$ docker run ko.local/github.com/davyj0nes/go-tzdata:d75ad91f74464cd0ac2e0d98a095e150199260facdd50275e018c87711fae55f
✅ got:  Pacific/Enderbury
✅ got:  Pacific/Kanton
done
```
