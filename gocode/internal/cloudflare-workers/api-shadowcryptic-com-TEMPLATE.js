addEventListener("fetch", (event) => {
  event.respondWith(
    handleRequest(event).catch(
      (err) => new Response(err.stack, { status: 500 })
    )
  );
});

/**
 * @param {Request} request
 * @returns {Promise<Response>}
 */
async function handleRequest(event) {
  const { request } = event;
  const headers = Object.fromEntries([...request.headers]);
  const origin = headers["origin"];

  // var allowed_origins = [
  //   "https://shadowcryptic.com",
  //   "https://poc.shadowcryptic.com",
  //   "https://dev.shadowcryptic.com",
  //   "http://local.shadowcryptic.com:1313",
  // ];

  // if (allowed_origins.includes(origin)) {
    return handleSiteRequest(event, origin);
  // }

  // return new Response("Not Found", { status: 404 });
}

function handleSiteRequest(event, origin) {
  const { request } = event;
  const { host, pathname } = new URL(request.url);

  // const headers = Object.fromEntries([...request.headers]);
  // const echo = { headers };
  // const body = JSON.stringify(echo, null, 2);

  switch (pathname) {
    case "/click":
      const body = JSON.stringify({}, null, 2);
      return new Response(body, {
        status: 200,
        headers: {
          "Content-Type": "application/json;charset=UTF-8",
          "Cache-Control": "no-store",
          "Access-Control-Allow-Origin": origin,
        },
      });
    case "/time":
      return respondTime(request);
    default:
      return new Response("Not Found", { status: 404 });
  }
}

/**
 * Responds with the local time.
 * @param {Request} request
 * @returns {Response}
 */
function respondTime(request) {
  const { headers, cf } = request;
  const locale = (headers.get("Accept-Language") || "en-US").split(",")[0];

  const timeStr = getTime(locale, cf);
  const data = {
    time_build: "POPULATE BUILD DATA HERE",
    time_now: timeStr,
  };
  const body = JSON.stringify(data, null, 2);
  return new Response(body, {
    status: 200,
    headers: {
      "Content-Type": "application/json;charset=UTF-8",
      "Content-Language": locale,
    },
  });
}

/**
 * Gets a human-readable description of the local time.
 * @param {string} locale
 * @param {object} cf
 * @returns {string}
 */
function getTime(locale, cf) {
  // In the preview editor, the 'cf' object will be null.
  // To view the object and its variables: make sure you are logged in,
  // click the "Save and Deploy" button, then open the URL in a new tab.
  const { city, region, country, timezone: timeZone } = cf || {};
  const localTime = new Date().toLocaleTimeString(locale, { timeZone });
  const location = [city, region, country].filter(Boolean).join(", ");

  return timeZone
    ? "The local time in " + location + " is " + localTime
    : "The UTC time is " + localTime;
}
