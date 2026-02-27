# SSE implementation in  GO

* This SSE is really great i hear because agents use it msot of the times. need to learn this for the mega project

## Pointers to rememeber
* The content-type is text/stream this is what initiates the SSE
* Cache-Control : This is for client/browser for data cache permissions. few types are public,private,
   nostore means never store and nocache means store but always revalidate with server
* The keep-alive header is auto handled by net/http
* Access-Control-Allow-Origin * allows for all sites to talk