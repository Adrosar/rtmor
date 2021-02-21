# RtMoR

_Real-time Modification of Requests_
- [github.com/Adrosar/rtmor](https://github.com/Adrosar/rtmor)
- [bitbucket.org/Adrosar/rtmor](https://bitbucket.org/Adrosar/rtmor)

**RtMoR** is alternative to **Requestly**

- [requestly.io](https://requestly.io) _(new website)_
- [www.requestly.in](https://www.requestly.in) _(old website)_

The project uses [github.com/elazarl/goproxy](https://github.com/elazarl/goproxy) as a proxy server.



## Examples (Linux, bash)

Remember to build a project!
```bash
./scripts/dist.sh
```

Show help:
```bash
./build/linux-amd64/rtmor -help
```

Run proxy server that is listening on all network interfaces:
```bash
./build/linux-amd64/rtmor -start -log -listen 0.0.0.0:8888
```

Run with a configuration that includes an example rule:
```bash
./build/linux-amd64/rtmor -start -log -listen 0.0.0.0:8888 -cfg ./configs/sample.yaml
```



## Examples (Windows, cmd)

Run with a configuration that includes an example rule:
```cmd
build\windows-amd64\rtmor.exe -start -log -listen 0.0.0.0:8888 -cfg configs\sample.yaml
```



## HTTPS and Certificate

For HTTPS redirection to work, install the certificate on the device.
**CA Root:** `./vendor/github.com/elazarl/goproxy/ca.pem`



## Runnable binaries

[Download](https://drive.google.com/drive/folders/1K4XvLZYB10pQ1iTYsRh0FlLP_PzwhNp4?usp=sharing) a copy of the repository and **binaries** ready to run.



## License
I put the software temporarily under the Go-compatible **BSD** license. If this prevents someone from using the software, do let me know and I'll consider changing it.

The software uses repositories:

- [github.com/elazarl/goproxy](https://github.com/elazarl/goproxy) _(Go-compatible BSD license)_
- [github.com/go-yaml/yaml](https://github.com/go-yaml/yaml) _(Apache License 2.0)_
- [github.com/fatih/color](https://github.com/fatih/color) _(MIT License)_



## Author

Adrian Gargula | [github.com/Adrosar](https://github.com/Adrosar) | [bitbucket.org/Adrosar](https://bitbucket.org/Adrosar)