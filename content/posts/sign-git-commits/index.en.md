---
weight: 2
title: "git: Signing commits"
date: 2021-04-03T15:14:12+01:00
# draft: true
tags: ["git"]
categories: ["Code"]
toc:
  enable: true
resources:
  - name: "featured-image"
    src: "featured-image.png"
  - name: "featured-image-preview"
    src: "featured-image-preview.png"
---

Git allows verification that work is from trusted or known sources by way of signed commits.

<!--more-->

## SSH setup

Firstly, if you don't have this setup already, to access git from the command line using an ssh key, simply ensure the `~/.ssh/config` file includes the key you want to use in an `IdentityFile` value. For example...

```txt
############################################################################
Host *
  AddKeysToAgent yes
  IgnoreUnknown UseKeychain
  UseKeychain yes
  IdentityFile ~/.ssh/id_ed25519
  .
  .
  .
```

You'll need to add that key to your github account too. See [Adding a new SSH key to your GitHub account](https://docs.github.com/en/github/authenticating-to-github/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)

## GPG setup

-   check available keys

    ```sh
    gpg --list-secret-keys --keyid-format LONG
    # ~/.gnupg/pubring.gpg
    # -------------------------------------
    # sec   rsa4096/876EA2A69C6C55F0 2018-07-10 [SC]
    #       0C10A2A69C19DC0D3806A51435C1336C876E55F0
    # uid                 [ultimate] Your Name <your.email@example.com>
    # ssb   rsa4096/F5C5F988807BAACF 2018-07-10 [E] [expires: 2021-01-10]
    ```

-   add knowledge of the key to git

    ```sh
    git config --global user.signingkey 0C10A2A69C19DC0D3806A51435C1336C876E55F0
    # set to always sign commits !!!
    git config --global commit.gpgsign true
    ```

-   to test

    ```sh
    echo "test" | gpg --clearsign
    ```

-   [Optionally] to remove a passphrase from the key

    Issue the command, then type `passwd` in the prompt.
    It will ask you to provide your current passphrase and then the new one.
    Just hit `Enter` for no passphrase.
    Then type `quit` to quit the program.

    ```sh
    gpg --edit-key <keyid>

    ┌────────────────────────────────────────────────────────────────┐
    │ Please enter the passphrase to unlock the OpenPGP secret key:  │
    │ "Your Name <your.email@example.com>"                        │
    │ 4096-bit RSA key, ID 876EA2A69C6C55F0,                         │
    │ created 2018-07-10.                                            │
    │                                                                │
    │                                                                │
    │ Passphrase: __________________________________________________ │
    │                                                                │
    │         <OK>                                    <Cancel>       │
    └────────────────────────────────────────────────────────────────┘
    ```

## Reference

-   [Git Tools: Signing Your Work](https://git-scm.com/book/en/v2/Git-Tools-Signing-Your-Work)
-   [GitLab: Signing commits with GPG](https://docs.gitlab.com/ee/user/project/repository/gpg_signed_commits/)
-   [Bitbucket: Using GPG keys](https://confluence.atlassian.com/bitbucketserver/using-gpg-keys-913477014.html)
