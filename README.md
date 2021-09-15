# Cloudflare Pages PoC

[<img src="https://gohugo.io/images/gopher-hero.svg" alt="hugo gopher hero logo" width="50"/>](https://gohugo.io/)
[![code style: prettier](https://img.shields.io/badge/code_style-prettier-ff69b4.svg?style=flat-square)](https://github.com/prettier/prettier)

A proof of concept for Cloudflare Pages.

My current workflow involves an SSG (Hugo) deployed (via a Makefile) to an S3 (at Scaleway), together with Cloudflare workers for an associated api. All, of course, fronted by Cloudflare and its handy caching.

So, interested to see how this can all be integrated into Cloudflare's Pages offering 😎

As an exercise in "Eating your own dog food", a guide detailing this PoC journey will be published using this repo's resulting site too! See the [Cloudflare Pages PoC](https://poc.shadowcryptic.com/cloudflare-pages-poc/) post.

## General

This repo contains the source code for a (Hugo generated) JAMstack. Documentation is available for:

-   [Development](./docs/development.md)

For general Hugo related documentation see the [Hugo docs](https://gohugo.io/documentation/).

## Note

I use YAML rather than TOML.
