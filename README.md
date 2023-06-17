# About

This uses VNC to send keystrokes to a VM using the packer `boot_command` syntax.

This was used to try to troubleshoot the issue at https://github.com/cirruslabs/packer-plugin-tart/issues/71.

# Usage (macOS)

Start the VM:

```bash
curl -fsSL \
    -o debian-12.0.0-arm64-netinst.iso \
    https://cdimage.debian.org/debian-cd/12.0.0/arm64/iso-cd/debian-12.0.0-arm64-netinst.iso
tart create --linux test-vnc --disk-size 16
tart run test-vnc --graphics --vnc-experimental --disk debian-12.0.0-arm64-netinst.iso:ro
```

At the VM console, in the debian bootloader prompt press ENTER to start the installation application.

Switch to tty2 by pressing ALT+F2 (or WIN+F2 in a PC keyboard).

Activate the console by pressing ENTER.

Type:

```bash
# tty 2
cat
```

Switch to tty3 by pressing ALT+F3 (or WIN+F2 in a PC keyboard).

Activate the console by pressing ENTER.

Type:

```bash
# tty 3
cat
```

At the host machine, open a new terminal window.

Build this application:

```bash
brew install go
make
```

From the `tart run` terminal window, locate a line alike:

```
Opening vnc://:enhance-chase-volume-push@127.0.0.1:59415...
```

And use those details to execute this application as, e.g.:

```bash
vnc_address='127.0.0.1:59415'
vnc_password='enhance-chase-volume-push'
boot_command='<leftCtrlOn><f2><leftCtrlOff>'
boot_command='<leftSuperOn><f2><leftSuperOff>'
boot_command='<leftAltOn><f2><leftAltOff>'
./use-packer-vnc-bootcommand \
    -address $vnc_address \
    -password $vnc_password \
    -boot-command $boot_command
```
