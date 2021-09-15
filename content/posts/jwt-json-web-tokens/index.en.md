---
weight: 4
title: "JWT: JSON Web Tokens"
date: 2019-10-18T12:00:00+01:00
# draft: true
tags: ["jwt"]
categories: ["Code"]
toc:
  enable: true
resources:
  - name: "featured-image"
    src: "featured-image.png"
  # - name: "featured-image-preview"
  #   src: "featured-image-preview.png"
---

JSON Web Token (JWT) is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object. This information can be verified and trusted because it is digitally signed. JWTs can be signed using a secret (with the HMAC algorithm) or a public/private key pair using RSA or ECDSA.

<!--more-->

JWTs can of course also be encrypted to provide secrecy between parties. Signed tokens can verify the integrity of the claims contained within it, while encrypted tokens hide those claims from other parties. When tokens are signed using public/private key pairs, the signature also certifies that only the party holding the private key is the one that signed it.

Checkout the [jwt.io debugger / playground](https://jwt.io/) for an online encoder/decoder.

## When should you use JSON Web Tokens

Here are some scenarios where JSON Web Tokens are useful:

**Authorization**: This is the most common scenario for using JWT. Once the user is logged in, each subsequent request will include the JWT, allowing the user to access routes, services, and resources that are permitted with that token. Single Sign On is a feature that widely uses JWT nowadays, because of its small overhead and its ability to be easily used across different domains.

**Information Exchange**: JSON Web Tokens are a good way of securely transmitting information between parties. Because JWTs can be signed — for example, using public/private key pairs — you can be sure the senders are who they say they are. Additionally, as the signature is calculated using the header and the payload, you can also verify that the content hasn't been tampered with.

## The JSON Web Token structure

In its compact form, JSON Web Tokens consist of three parts separated by dots (.), which are:

-   Header
-   Payload
-   Signature

Therefore, a JWT typically looks like `xxxxx.yyyyy.zzzzz`

**Note**: For signed tokens, although protected against tampering, the information is readable by anyone.
**Do not put secret information in the payload or header elements** of a JWT unless it is encrypted.

## Example

```js
myJWTSecret = 'thisismysecret!'

jwtOptions = {
    expiresIn: '10 seconds'
}

jwtObject = {
    _id: 'abc123'
}

token = jwt.sign(jwtObject, myJWTSecret, jwtOptions)

//  => base64 encoded header.payload.signature
// e.g. eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiJhYmMxMjMiLCJpYXQiOjE1NzEzMjY3MjcsImV4cCI6MTU3MTMyNjczN30.7BV9302v6BtbOslWEa1pD5BXLHI2kuq0iYFUcQ4J2M4

// header:          eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
//  => {
//      "alg":"HS256",
//      "typ":"JWT"
//  }

// payload:         eyJfaWQiOiJhYmMxMjMiLCJpYXQiOjE1NzEzMjY3MjcsImV4cCI6MTU3MTMyNjczN30
//  => {
//      "_id": "abc123",
//      "iat":1571326727
//      "exp":1571326737
//  }

// signature:       7BV9302v6BtbOslWEa1pD5BXLHI2kuq0iYFUcQ4J2M4

jwt.verify(token, myJWTSecret)
// => true|false
```

Use [base64decode.org](https://www.base64decode.org/) to decode a JWT's header or body

## Reference

-   [https://jwt.io/introduction/](https://jwt.io/introduction/)
