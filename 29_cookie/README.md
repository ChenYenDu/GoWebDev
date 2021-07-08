# Cookie

A cookie is a file that store on the client's machine.

Cookie are written by domains to store information on the client's machine.

## Domain specific

Cookies are only sent by the browser to the domain which wrote them.

With every request to a specific domain, a client's browser looks to see if there is a cookie from that domain on the client's machine. If there is a cookie that has been written by that particular domain, then the browser will send the cookie with every request to that domain.

Cookies are domain specific.

## Limits

We can store whatever information we would like in a cookie up to a certain size limit. The size limit of a cookie is dependent upon the browser but is usually around 4096 characters.

There is also a limit to the number of cookies one domain can write. This limit is also browser specific.

See [the source](http://browsercookielimits.squawky.net/) for more information.

## Expiring a cookie

If the __Expires__ or __MaxAge__ field isn't set, the cookie is deleted when the browser is closed. This is colloquially known as a "session cookie".

You can expire a cookie by setting one of these two fields: __Expires__ or __MaxAge__

__Expires__ sets the exact time when the cookie expires. Expires is __Deprecated__.

__MaxAge__ sets how long the cookie should live (in seconds).
