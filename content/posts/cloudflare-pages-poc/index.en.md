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

## Basic Production Site

1.  Checkout the CF docs

    Follow Cloudflare's [Deploy a Hugo site docs](https://developers.cloudflare.com/pages/framework-guides/deploy-a-hugo-site).

    I'm going to be using the [CodeIT](https://github.com/sunt-programator/CodeIT.git) theme, so I'm using that where the docs use `themes/terminal`. Similarly, my config is set up as seen in [my repo](https://github.com/timecode/CloudflarePagesPoC).

2.  Set a custom DNS entry

    Set a Custom Domain to allow easier naivigation to your site (set at Cloudflare's hosted DNS, for example) such as `https://poc.shadowcryptic.com` (set subdomain `poc` entry `CNAME` to `cloudflarepagespoc.pages.dev`)

3.  Build settings

    -   Build command: `hugo`
    -   Build output directory: `/public`
    -   Root directory: `/`

    Build dependencies versions

    {{< admonition >}}
    Default versions for build dependencies may well be quite old.
    {{< /admonition >}}

    There are separate environment variables for `Production` and `Preview` builds, so always try out new versions in a Preview build before upgrading the Production ENV.

    Either way, environment variables should be set to match your development environment.

    Go to your Pages Project page > Settings > Environment Variables

    Add the following, for example:

    | Variable name      | Value        |
    | ------------------ | ------------ |
    | `GO_VERSION`       | `1.17`       |
    | `HUGO_VERSION`     | `0.88.1`     |
    | `HUGO_ENVIRONMENT` | `production` |

    **Note**: `HUGO_ENVIRONMENT` is used by Hugo as a selector for its configuration. This repo's config directory, for example, has two such options 'production' and 'development'

4.  Deploy

    Deployment is simply a matter of pushing your local commits to the repo. The default branch for production builds is `main` so commit or merge to that and Cloudflare should initiate a build and deploy to your production site!

    The build can we followed by clicking on `View Build` and following the Build log, for example:

    ```txt
    13:12:51.819 Initializing build environment. This may take up to a few minutes to complete
    13:15:27.988 Success: Finished initializing build environment
    13:15:27.988 Cloning repository...
    13:15:32.126 Success: Finished cloning repository files
    13:15:33.008 Installing dependencies
    13:15:33.011 Python version set to 2.7
    13:15:34.226 v12.18.0 is already installed.
    13:15:34.736 Now using node v12.18.0 (npm v6.14.4)
    13:15:34.777 Started restoring cached build plugins
    13:15:34.781 Finished restoring cached build plugins
    13:15:34.910 Attempting ruby version 2.7.1, read from environment
    13:15:35.944 Using ruby version 2.7.1
    13:15:36.221 Using PHP version 5.6
    13:15:36.252 5.2 is already installed.
    13:15:36.257 Using Swift version 5.2
    13:15:36.258 Started restoring cached node modules
    13:15:36.260 Finished restoring cached node modules
    13:15:36.262 Started restoring cached yarn cache
    13:15:36.265 Finished restoring cached yarn cache
    13:15:36.268 Installing yarn at version 1.22.4
    13:15:36.272 [37mInstalling Yarn![0m
    13:15:36.272 [36m> Downloading tarball...[0m
    13:15:36.281
    13:15:36.281 [1/2]: https://yarnpkg.com/downloads/1.22.4/yarn-v1.22.4.tar.gz --> /tmp/yarn.tar.gz.zT8fSfiMql
    13:15:36.282   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
    13:15:36.283                                  Dload  Upload   Total   Spent    Left  Speed
    13:15:36.369
      0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
    100    79  100    79    0     0    900      0 --:--:-- --:--:-- --:--:--   908
    13:15:36.506
    100    93  100    93    0     0    413      0 --:--:-- --:--:-- --:--:--   413
    13:15:36.616
      0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
    100   625  100   625    0     0   1868      0 --:--:-- --:--:-- --:--:-- 11574
    13:15:36.722
    100 1215k  100 1215k    0     0  2758k      0 --:--:-- --:--:-- --:--:-- 2758k
    13:15:36.722
    13:15:36.722 [2/2]: https://yarnpkg.com/downloads/1.22.4/yarn-v1.22.4.tar.gz.asc --> /tmp/yarn.tar.gz.zT8fSfiMql.asc
    13:15:36.735
    100    83  100    83    0     0   6655      0 --:--:-- --:--:-- --:--:--  6655
    13:15:36.743
    100    97  100    97    0     0   4617      0 --:--:-- --:--:-- --:--:--  4617
    13:15:36.746
    100   629  100   629    0     0  26002      0 --:--:-- --:--:-- --:--:-- 26002
    13:15:36.749
    100  1028  100  1028    0     0  38460      0 --:--:-- --:--:-- --:--:-- 38460
    13:15:36.780 [36m> Verifying integrity...[0m
    13:15:36.805 gpg: Signature made Mon 09 Mar 2020 03:52:13 PM UTC using RSA key ID 69475BAA
    13:15:36.812 gpg: Good signature from "Yarn Packaging <yarn@dan.cx>"
    13:15:36.815 gpg: Note: This key has expired!
    13:15:36.815 Primary key fingerprint: 72EC F46A 56B4 AD39 C907  BBB7 1646 B01B 86E5 0310
    13:15:36.815      Subkey fingerprint: 6D98 490C 6F1A CDDD 448E  4595 4F77 6793 6947 5BAA
    13:15:36.816 [32m> GPG signature looks good[0m
    13:15:36.816 [36m> Extracting to ~/.yarn...[0m
    13:15:36.868 [36m> Adding to $PATH...[0m
    13:15:36.872 [36m> We've added the following to your /opt/buildhome/.bashrc
    13:15:36.873 > If this isn't the profile of your current shell then please add the following to your correct profile:
    13:15:36.873
    13:15:36.873 export PATH="$HOME/.yarn/bin:$HOME/.config/yarn/global/node_modules/.bin:$PATH"
    13:15:36.873 [0m
    13:15:37.182 [32m> Successfully installed Yarn 1.22.4! Please open another terminal where the `yarn` command will now be available.[0m
    13:15:37.503 Installing NPM modules using Yarn version 1.22.4
    13:15:37.900 yarn install v1.22.4
    13:15:37.944 [1/4] Resolving packages...
    13:15:37.969 [2/4] Fetching packages...
    13:15:38.493 [3/4] Linking dependencies...
    13:15:38.547 [4/4] Building fresh packages...
    13:15:38.555 Done in 0.66s.
    13:15:38.572 NPM modules installed using Yarn
    13:15:38.787 Installing Hugo 0.88.1
    13:15:39.639 hugo v0.88.1-5BC54738+extended linux/amd64 BuildDate=2021-09-04T09:39:19Z VendorInfo=gohugoio
    13:15:39.640 Started restoring cached go cache
    13:15:39.643 Finished restoring cached go cache
    13:15:39.644 Installing Go version 1.17
    13:15:45.024
    13:15:45.024 unset GOOS;
    13:15:45.024 unset GOARCH;
    13:15:45.024 export GOROOT='/opt/buildhome/.gimme_cache/versions/go1.17.linux.amd64';
    13:15:45.024 export PATH="/opt/buildhome/.gimme_cache/versions/go1.17.linux.amd64/bin:${PATH}";
    13:15:45.024 go version >&2;
    13:15:45.024
    13:15:45.026 export GIMME_ENV="/opt/buildhome/.gimme_cache/env/go1.17.linux.amd64.env"
    13:15:45.033 go version go1.17 linux/amd64
    13:15:45.034 Installing missing commands
    13:15:45.034 Verify run directory
    13:15:45.035 Executing user command: hugo
    13:15:45.101 Start building sites …
    13:15:45.101 hugo v0.88.1-5BC54738+extended linux/amd64 BuildDate=2021-09-04T09:39:19Z VendorInfo=gohugoio
    13:15:45.658
    13:15:45.659                    | EN
    13:15:45.659 -------------------+-----
    13:15:45.659   Pages            | 73
    13:15:45.659   Paginator pages  |  1
    13:15:45.659   Non-page files   | 24
    13:15:45.659   Static files     | 96
    13:15:45.659   Processed images |  0
    13:15:45.659   Aliases          | 21
    13:15:45.659   Sitemaps         |  1
    13:15:45.659   Cleaned          |  0
    13:15:45.659
    13:15:45.659 Total in 601 ms
    13:15:45.665 Finished
    13:15:45.665 Validating asset output directory
    13:15:46.875 Deploying your site to Cloudflare's global network...
    13:15:53.603 Success: Your site was deployed!
    ```

5.  Marvel at the deployed site

    - The deployed site gets a TLS certificate provided by Cloudflare 😎.
    - Deployments can be rolled back (to any version) at any time.
    - Each deployment has its own unique subdomain, such as `https://331e1fb6.cloudflarepagespoc.pages.dev/` for example.
    - Access to Preview deployments can be control with Cloudflare Access.
    - Web Analytics can be enabled at Cloudflare.
    - Check your site's score on Google's [PageSpeed Insights](https://developers.google.com/speed/pagespeed/insights/)... for example, [here's this site](https://developers.google.com/speed/pagespeed/insights/?url=https%3A%2F%2Fpoc.shadowcryptic.com%2F&tab=desktop)

## Preview Site

Preview branches are classed as "all non-Production branches". Try one out...

1.  Create a branch... `git checkout -b my-preview-version-1`
2.  Make you changes
3.  Commit your changes (repeat the previous step for more changes, or continue)
4.  Push to the repo `git push origin`
5.  Watch Cloudflare Pages pick up the branch and deploy a preview! The deployment gets its own subdomain, such as `https://df58e22a.cloudflarepagespoc.pages.dev/` for example, however it also gets a handy alias containing the branch name, such as `https://my-preview-version-1.cloudflarepagespoc.pages.dev`

When you're ready to merge the branch to `main`, open a pull request, watch as the Cloudflare Pages build completes to allow your Merge, which will then trigger a Production build and deployment!

## Adding a (custom, dynamically generated) Cloudflare Worker

To show the build process can include more than just the static site generation, we'll add a Cloudflare Worker to act as a simple API. The worker will be dynamically generated as well, just "because we can".

The api will simply return a json response containing the worker's code creation time (to satisfy the dynamic requirement updated at build time) as well as the current time whenever the api is called. The result will therefore be something like:

```json
{
  "time_build": "2021-09-15T17:02:03Z",
  "time_now": "The local time in GB is 6:24:40 PM"
}
```

1.  Add environment variables `CF_` shown below:

    | Variable name          | Value                   |
    | ---------------------- | ----------------------- |
    | `CF_WORKERS_API_TOKEN` | `<ADD API TOKEN HERE>`  |
    | `CF_ACCOUNT_ID`        | `<ADD ACCOUNT ID HERE>` |
    | `CF_ZONE_ID`           | `<ADD ZONE ID HERE>`    |
    | `HUGO_VERSION`         | `0.88.1`                |
    | `HUGO_ENVIRONMENT`     | `production`            |

2.  Add the cloudflare worker deployment code as seen in the [repo](https://github.com/timecode/CloudflarePagesPoC/tree/main/gocode/).

3.  Add code to hook in to the Hugo deployment as seen in the repo's [Makefile](https://github.com/timecode/CloudflarePagesPoC/tree/main/Makefile). Update the Cloudflare Pages `Build command` from the regular `hugo` command to now use the Makefile with the command `make cloudflare-deploy`

4.  Add a CNAME to the DNS to allow the api e.g. `CNAME api-poc-dev shadowcryptic.com`

Deployment should now include/update this Cloudflare worker whenever the site is updated (the `time_build` field should be seen to update).

Clicking on the [api endpoint](https://api-poc.shadowcryptic.com/time) should provide something like the above example. We could of course add a simple piece of JavaScript to a page that automatically calls the API on load and updates a page dynamically, but that's another story.

Adding dynamic elements to the Hugo generated SSG is now fairly simple to develop and deploy using Cloudflare Pages.
🥂 😎 🥂
