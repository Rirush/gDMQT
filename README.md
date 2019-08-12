# gDMQT
> Go implementation of a Datagram Message Queuing Transport.

[![Travis CI](https://img.shields.io/travis/com/Hexawolf/gDMQT.svg?style=flat-square&logo=Linux)](https://travis-ci.com/Hexawolf/gDMQT)
[![CodeCov](https://img.shields.io/codecov/c/github/Hexawolf/gDMQT.svg?style=flat-square)](https://codecov.io/gh/Hexawolf/gDMQT)
[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg?style=flat-square)](https://github.com/emersion/stability-badges#experimental)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FHexawolf%2FgDMQT.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FHexawolf%2FgDMQT?ref=badge_shield)

gDMQT is an experimental broker implementation, based on MQTT-SN protocol.
Any client that implements MQTT-SN properly may use this broker for sending and
receiving messages, though it must be expected that gDMQT violates common
broker behaviour.

The standard is currently work-in-progress and gDMQT just implements MQTT-SN
right now.

## Getting started

gDMQT is mainly tested against Linux and must work on any common distribution
that can run a Go compiler. Building should be simple:

```bash
go build
```

After building, rename `broker.example.cfg` into `broker.cfg` and change
values according to your needs. Hopefully, now you are ready to run it.

## License

The code is under MIT license. See [LICENSE](LICENSE) for more information.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FHexawolf%2FgDMQT.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FHexawolf%2FgDMQT?ref=badge_large)
