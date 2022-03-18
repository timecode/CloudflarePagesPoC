# Development

See the [Prerequisites](#Prerequisites) section below for requirements when setting up and maintaining this repo. A separate [development resources doc](./dev-resources.md) is also available.

Invoke local server with:

```sh
make local
# alternative aliases ...
# make dev
# make serve
```

## Submodule

The site currently uses a base theme, which is included as a gitsubmodule.

-   Add Submodule

    ```sh
    git submodule add https://github.com/sunt-programator/CodeIT.git themes/CodeIT
    ```

-   Update Submodule

    ```sh
    # Occasionally update the submodule to a new version:
    git -C themes/CodeIT checkout [<new version>]

    # related commands ...
    git add themes/CodeIT
    git commit -m "update themes/CodeIT submodule to latest version"

    # See the list of submodules in a superproject
    git submodule status
    ```

-   [When/If required] Hard reset

    ```sh
    make theme-reset
    ```

## Prerequisites

-   **Editor**: VSCode

    Prettier is used for code formatting. The editor being used should therefore have Prettier installed (see [Prettier: Editor Integration](https://prettier.io/docs/en/editors.html)). A local (pinned) version is installed in this repo to maintain consistency (as recommended in the Prettier [docs](https://github.com/prettier/prettier-vscode#prettier-resolution)). Update, as and when necessary:

    -   Install local plug-in

        ```sh
        # yarn remove prettier
        yarn add --dev --exact prettier@2.6.0

        # additional plugin for go-template
        # yarn remove prettier-plugin-go-template
        yarn add --dev --exact prettier-plugin-go-template@0.0.11
        ```

-   **DNS**: Cloudflare

    -   Minimum TLS Version 1.2 (Safari does not support 1.3 by default, still!). Verify browser compatibility from this [test page](https://www.cloudflare.com/en-gb/ssl/encrypted-sni/)

-   **Firewall**: Cloudflare

    Firewall Rules 'may' be set at Cloudflare, for example, to block/allow/challenge particular requests. Verify all rules on a separate domain before setting on a prod domain.

    For SSG sites a firewall rule such as the following would be good:

    -   Action: `BLOCK` (returns a 403)

        ```txt
        (http.request.uri.path contains "wp-") or
        (http.request.uri.path contains "admin") or
        (http.request.uri.path contains ".php") or
        (http.request.uri.path contains ".asp") or
        (http.request.uri.path contains ".env")
        ```

    To blacklist an old site:

    -   Action: `BLOCK` (returns a 403)

        ```txt
        (http.request.uri.path contains "Web_2020")
        ```

-   **Page Rules**: Cloudflare

    Page Rules may also be set at Cloudflare, for example, to redirect particular requests (with a 301) rather than block them (with a 403). There is a limit under the Free Tier at Cloudflare though, so this may only be of use to redirect requests to a particular path or sub path.

-   **Logging**: Cloudflare App

    [Logflare](https://logflare.app/) is a Cloudflare App that provides request logging :-) It can also link with ipinfo.io to provide location data. Appears to be customisable too, see [Logflare github repo](https://github.com/Logflare/cloudflare-app)

    Search query examples:

    -   'Non-bot' page requests

        ```txt
        -m.request.headers.user_agent:~"Go-http-client|[B|b]ot|spider|crawler|Minefield|Lighthouse|facebook|.html|.com|[H|h]ttp|Favicon|[H|h]eadless" -m.request.url:~"[api|dev|time|www]\.shadowcryptic|images|favicon|plugins|css|js|robots" m.response.status_code:<300 c:count(*) c:group_by(t::hour)
        ```

    -   Non-bot/Non-200 results

        ```txt
        -m.request.headers.user_agent:~"Go-http-client|[B|b]ot|spider|crawler|Minefield|Lighthouse|facebook|.html|.com|[H|h]ttp|Favicon|[H|h]eadless" -m.request.url:~"[api|dev|time|www]\.shadowcryptic" m.response.status_code:>200 c:count(*) c:group_by(t::hour)
        ```

    -   Non-dev 'cache MISS'

        ```txt
        -m.request.headers.user_agent:~"Go-http-client|[B|b]ot|spider|crawler|Minefield|Lighthouse|facebook|.html|.com|[H|h]ttp|Favicon|[H|h]eadless" -m.request.url:~"[api|dev|time]\.shadowcryptic" m.response.headers.cf_cache_status:MISS m.response.status_code:>301 c:count(*) c:group_by(t::hour)
        ```

    -   API requests

        ```txt
        m.request.url:~"://api" t:last@28day c:count(*) c:group_by(t::hour)

        m.request.url:~"://api.*/click" t:last@28day c:count(*) c:group_by(t::hour)

        m.request.url:~"://api.*/click.*SHADOW1" t:last@28day c:count(*) c:group_by(t::hour)

        m.request.url:~"://api.*/click.*Buy.*SHADOW1" t:last@28day c:count(*) c:group_by(t::hour)

        m.request.url:~"://api.*/click.*Listen.*SHADOW1" t:last@28day c:count(*) c:group_by(t::hour)
        ```

-   **Caching**: Cloudflare

    -   [Customizing Cache](https://support.cloudflare.com/hc/en-us/articles/202775670-How-do-I-cache-static-HTML-)
    -   Cache Everything Page Rule()

    Caching additional content at Cloudflare requires a `Cache Everything Page Rule`. Without creating a `Cache Everything Page Rule`, dynamic assets are never cached, even if a public Cache-Control header is returned. When combined with an `Edge Cache TTL > 0`, Cache Everything removes cookies from the origin web server response.

    1. Log in to your Cloudflare account.
    2. Choose the appropriate domain.
    3. Click the `Rules` app/tab.
    4. Create a URL pattern to differentiate your website’s static versus dynamic content or something to cover all content e.g. `*shadowcryptic.com/*` (or `shadowcryptic.com/*` to exclude `dev` subdomin)
    5. Choose the Cache Level setting and then the `Cache Everything` submenu setting.
    6. Click Save to Deploy.
    7. Verify resources are cached by checking the [cache response returned by Cloudflare](https://support.cloudflare.com/hc/articles/200172516#h_bd959d6a-39c0-4786-9bcd-6e6504dcdb97) when viewing the `cf-cache-status` header in the response e.g. `curl -v https://shadowcryptic.com`

    When deploying a new site, the cache can be purged by visiting the `Caching` app/tab.

    1.  Select `Configuration`
    2.  In the `Purge Cache` section, select the desired purge option.
    3.  Purging the cache at deploy time, new content will be available by the value of the `Browser Cache TTL` which can be set to something like `4 hours`. **Note**: `Browser Cache TTL` allows a browser session to be speedier as it doesn't call out to the Internet for data; it just uses content from a local cache.
    4.  Alternatively, use the API e.g.:

        ```sh
        # Zone ID for a domain is shown on the `Overview` app/tab > API > Zone ID
        # API Token and scope is configured at the account level
        # https://dash.cloudflare.com/profile/api-tokens
        ZoneID="<INSERT ZONE ID HERE>" && \
        APIToken="<INSERT API TOKEN HERE>" && \
        curl -X POST "https://api.cloudflare.com/client/v4/zones/${ZoneID}/purge_cache" \
            --header "Authorization: Bearer ${APIToken}" \
            --header "Content-Type:application/json" \
            --data '{"purge_everything":true}'
        ```

    **Note**: Cloudflare's "Cache Everything" simply **skips the extension check**, and all content is treated as cacheable.

-   **Hosting**: We're going to use Cloudflare Pages rather than an S3 bucket.

    -   Ensure site works (direct to <http://cloudflarepagespoc.pages.dev/>) given a simple `index.html` file.

    -   Point DNS to site, for example:

        ```txt
        CNAME    @    cloudflarepagespoc.pages.dev
        ```

        **Note**: remove any conflicting A record beforehand

    -   Ensure Staging Site isn't able to be 'listed'

        Add a [robots.txt](https://developers.google.com/search/docs/advanced/robots/create-robots-txt) file containing:

        ```txt
        # DO NOT include in production deployment!
        # Block all access to site
        User-agent: *
        Disallow: /
        ```

        Note: Prod endpoints can be tested against the `robots.txt` file using the [robots-testing-tool](https://www.google.com/webmasters/tools/robots-testing-tool?siteUrl=https%3A%2F%2Fshadowcryptic.com%2F&path=sitemap.xml) (modify the path parameter)

    -   [OPTIONAL] Check DNS routes correctly, maybe using `curl` to view, for example:

        ```sh
        curl -vvv -L -I https://poc.shadowcryptic.com

        # if regular curl doesn't work (it can't do TLS1.3 yet)
        # brew install --build-from-source curl
        /usr/local/opt/curl/bin/curl -vvv -L -I https://poc.shadowcryptic.com
        # or, to force a TLS version
        /usr/local/opt/curl/bin/curl -vvv --tlsv1.2 -L -I https://poc.shadowcryptic.com
        ```
