---
weight: 1
title: "Cloudflare Pages PoC"
date: 2021-09-15T09:45:40+01:00
description: This article describes my Cloudflare Pages PoC.
toc:
  enable: false
resources:
  - name: featured-image
    src: featured-image.png
  - name: featured-image-preview
    src: featured-image-preview.png
tags:
  - Cloudflare
  - Pages
categories:
  - Cloudflare
---
This article describes my Cloudflare Pages PoC.
<!--more-->

{{< admonition >}}
This PoC is WIP, but here's what I have so far...
{{< /admonition >}}

## Step 1 - Checkout the CF docs

Follow Cloudflare's [Deploy a Hugo site docs](https://developers.cloudflare.com/pages/framework-guides/deploy-a-hugo-site).

I'm going to be using the [CodeIT](https://github.com/sunt-programator/CodeIT.git) theme, so I'm using that where the docs use `themes/terminal`. Similarly, config is set up as seen in this repo.

## Step 2 - Set a custom DNS entry

Set a Custom Domain to allow easier naivigation to your site (set at Cloudflare's hosted DNS, for example) such as `https://poc.shadowcryptic.com` (set subdomain `poc` entry `CNAME` to `cloudflarepagespoc.pages.dev`)

## Step 3 - Build settings

- Build command: `hugo`
- Build output directory: `/public`
- Root directory: `/`

### Build dependencies versions

{{< admonition >}}
Default versions for build dependencies may well be quite old.
{{< /admonition >}}

There are separate environment variables for `Production` and `Preview` builds, so always try out new versions in a Preview build before upgrading the Production ENV.

Either way, environment variables should be set to match your development environment.

Go to your Pages Project page > Settings > Environment Variables

Add the following, for example:

| Variable name  | Value    |
| -------------- | -------- |
| `GO_VERSION`   | `1.17`   |
| `HUGO_VERSION` | `0.88.1` |

## Step 4 - Deploy

Deployment is simply a matter of pushing your local commits to the repo. The default branch for production builds is `main` so commit or merge to that and Cloudflare should initiate a build and deploy to your production site!
