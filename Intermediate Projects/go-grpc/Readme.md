GRPC - Google Remote Procedural Call

* This is a protocal built on top of http/2. the upgrade over rest. per say 
* Key features is the serialization techniquee which is a protobuf over json.
* http/2 supports framing and multiplexing and grpc exploits this feature perfectly
* The protobuf is well typed with better security and bug hunting mechanisms catching errors at compile times
* Streaming is one of the best ways to use grpc since it is really good with small frame transfers

* Browers don't natively support grpc. we need grpc-web
* Client side api's / browser api's should still use websockets and standard http
* For internal fast communications between two or more microservices GRPC stands out. 
* The http/3 (QUIC over UDP) also supports GRPC 