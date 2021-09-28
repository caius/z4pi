# Z4Pi

[I own a 2003 BMW Z4][z4 tweet]. It has some irritations, and I think I can overcome or augment them with technology.

## Background

The E85 BMW Z4 has a kbus network in it for non-critical systems (think door sensors, roof commands, windows up/down, etc.) https://modbmw.com sell a USB adapter for the kbus, exposing it via serial port to computer. Raspberry Pi Zero W fits the bill nicely, sits in the boot with the adapter plugged into the CD Changer kbus cables. <https://github.com/ezeakeal/pyBus> or <https://github.com/qcasey/gokbus> are existing libraries to decipher/speak kbus traffic.

There’s a third party device you can plug in to allow one-press roof up/down functionality, which sounds like it just broadcasts onto the kbus network the same packets as if you were holding the button(s) down constantly. Also I hate pressing the sport button after starting the car *every time*. (Sport mode just makes the throttle linear.)

[z4 tweet]: https://twitter.com/Caius/status/1426857721912074240

## Functionality

Currently implemented is … nothing.

- [ ] Serial reading logged to file
- [ ] Webserver listening / exposing logfile
