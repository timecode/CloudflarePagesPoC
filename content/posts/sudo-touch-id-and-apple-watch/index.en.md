---
weight: 2
title: "sudo: Authenticate with Touch ID, including Apple Watch"
date: 2020-02-08T21:52:36+01:00
# draft: true
tags: ["Touch ID","Apple Watch"]
categories: ["Command line"]
toc:
  enable: false
resources:
  - name: "featured-image"
    src: "featured-image.png"
---

Having to enter a password when asked by a `sudo` command gets old, very quickly, when you have a long password. More so, it's relatively insecure. Apart from the potential of keyboard loggers, I'm always paranoid that, however fast I can type, some hi-jinxing team mates could always be videoing my hands, fingers and keyboard! Using an external physical device, like a [YubiKey](https://www.yubico.com/products/yubikey-5-overview/) for example, is much safer. However, if you have a MacBook with Touch ID or an Apple Watch, they can be used to authenticate your sudo commands instead 😎

<!--more-->

Making an addition to the `/etc/pam.d/sudo` file is all that's required.

## Step 1

We need to temporarily allow the file to be modifiable, so:

```sh
# if required
sudo chmod 644 /etc/pam.d/sudo
```

## Step 2

Using a suitable editor (e.g. `sudo vi /etc/pam.d/sudo`), add the line `auth       sufficient     pam_tid.so` to the top of the file so that the contents look something like...

```txt
# sudo: auth account password session
auth       sufficient     pam_tid.so
auth       sufficient     pam_smartcard.so
auth       required       pam_opendirectory.so
account    required       pam_permit.so
password   required       pam_deny.so
session    required       pam_permit.so
```

## Step 3

**Important**: Remember to remove those earlier permissions from the file, e.g.:

```sh
# if required
sudo chmod 444 /etc/pam.d/sudo
```

## That's it

Now, when you attempt a `sudo` command, you'll be prompted with a Touch ID authentication in lieu of entering your administrator password. Either respond by placing an appropriate finger on the Touch ID reader or 'OK' the notification on your Apple Watch. Yay 🎉
