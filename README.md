# Intermittent Failure Service

This service returns a payload intermittently, other times responding with 500-series errors:

* 500 internal server error
* 503 service unavailable
* 504 gateway timeout

ref [RFC 7231 Section 6.6](https://tools.ietf.org/html/rfc7231#section-6.6.1)
