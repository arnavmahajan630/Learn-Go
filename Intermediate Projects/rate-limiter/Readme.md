Footnotes
Rate Limiting: Protecting API's from DOS/ DDOS

Approaches to Rate Limiting

1) Token Bucket:

A bucket is created. Tokens are added to this bucket at a fixed rate. Each request comsumes tokens
IF there are not enough tokens in the bucket the Request is delayed/blocked/rejected

Advantages
* Controlled bursts
* Long term rate
* O(1) State

-> NewLimiter takes 1) rate at which tokens are added per sec 2) burst. i.e maximum tokens the bucket can hold


2) Per Client Limiting.
Token bucket is an efficeint algorithm but by default we limit whole app
per client is limiting based on 
* cookies
* sessions
* users
* ip
* tokens
* IP's
* geo location and more

3) Toolbooth

A strong and robust library that combines the x/time/rate and per client limiting with mroe exciting features like middleware integration, auto clean, and some more.

Not useful with advanced distributed systems. there just build your own rate limiter using the x/time/rate.

