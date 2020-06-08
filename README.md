
# Battery Monitor

Battery Monitor is an X-platform utility tool developed on Golang. It will notify the user about charging, discharging, and critically low battery state of the battery (surely if the battery is present).

 - [Installation](#installation)
 - [Issue Tracking](#issue-tracking)
 - [Roadmap](#roadmap)
 - [Changelog](#changelog)
 - [Contributors](#contributors)

## Installation

### Linux (64-bit)

To install the daily build, please ensure the following dependencies installed:

- git (>=2.17)
- golang (>=1.11)

Now run the following command:

```shell script
$ curl https://raw.githubusercontent.com/maateen/battery-monitor/dev/scripts/install.sh | bash -
```

## Roadmap

- [x] Porting into Golang
- [x] Adding X-platform support
- [x] Adding configuration support
- [ ] Adding system stray icon
- [ ] Adding support for notification duration
- [ ] Adding support for sound notification
- [ ] Adding support for permanent notification until another event
- [ ] Packaging for Linux
    - [ ] Packaging for Debian/Ubuntu
    - [ ] Packaging for CentOS/Fedora
- [ ] Packaging for MacOSX
- [ ] Packaging for Windows
- [ ] Using DBUS for notification on Linux
- [ ] Fixing notification on MacOSX
- [ ] Adding new version notification
- [ ] Adding CI
- [ ] Adding automated release

## Changelog

### v0.6

- [x] Restructured and reformatted the whole project.
- [x] Added system tray icon ([Issue #46](https://github.com/maateen/battery-monitor/issues/46))
- [x] Fixed [issue #51](https://github.com/maateen/battery-monitor/issues/51)
- [x] Added some new icons ([Issue #53](https://github.com/maateen/battery-monitor/issues/53))
- [x] Added Ubuntu 18.04 LTS support ([Issue #55](https://github.com/maateen/battery-monitor/issues/55))
- [x] Fixed [issue #61](https://github.com/maateen/battery-monitor/issues/61)

### v0.5.4

- [x] Fixed [issue #48](https://github.com/maateen/battery-monitor/issues/48)

### v0.5.3

- [x] Fixed [issue #45](https://github.com/maateen/battery-monitor/issues/45)
- [x] Support for Ubuntu 17.10 has been added.

### v0.5.2

- [x] Fixed [issue #41](https://github.com/maateen/battery-monitor/issues/41)
- [x] Fixed [issue #42](https://github.com/maateen/battery-monitor/issues/42)
- [x] Introduced a Test feature for developers.

### v0.5.1

- [x] Fixed [issue #35](https://github.com/maateen/battery-monitor/issues/35)
- [x] Fixed [issue #39](https://github.com/maateen/battery-monitor/issues/39)

### v0.5

- [x] Developed a GUI to manage the custom warning easily.
- [x] Minimized CPU consumption.
- [x] Added Makefile for easy installation and up-gradation.
- [x] Re-structured the project.
-  [x] Support for Ubuntu 14.04, 16.04, 16.10 and 17.04 has been added.

### v0.4

- [x] Reformatted the code in a new style.
- [x] Optimized the code in a way so that Battery Monitor consumes a little resource of your PC.

### v0.3

- [x] Fixed [issue #7](https://github.com/maateen/battery-monitor/issues/7), decreased CPU consuming from 40% to below 0.7%
- [x] Fixed [issue #4](https://github.com/maateen/battery-monitor/issues/4), Added warning at 30% battery life (temporary solution, will be replaced by a GUI in near future).
- [x] Fixed [issue #6](https://github.com/maateen/battery-monitor/issues/6), Added an entry in dash.

### v0.2.1

- [x] Added trusty support.

### v0.2

- [x] Added **Critically Low Battery** warning when the battery is below 10%.
- [x] Added `ctrl+C` pressing support to stop the `battery-monitor` command on terminal.

### v0.1

- [x] Initial release.

## Contributors

### [Safwan Rahman](https://github.com/safwanrahman)

He has reformatted the code in a new style. The style represents the code easier to understand. Also, he has optimized the code in a way that **Battery Monitor** consumes a little resource of your PC. Please take a minute to thank him.

### [Abdelhak Bougouffa](https://abougouffa.github.io/)

He has fixed some bugs and optimized **Battery Monitor** in a way so that it consumes lower CPU and RAM than before. Please take a minute to thank him.

### [Yochanan Marqos](https://github.com/yochananmarqos)

He is our official package maintainer in AUR. He has put Arch users' life at ease. Please take a minute to thank him.

### [Benjamin Staffin](https://github.com/benley)

He has improved the build process and added modern Python setuptools packaging. Please take a minute to thank him.
