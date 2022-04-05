# WebApps 

WebApps that can take to gRPC endpoints via gRPC-Web transport 

_protobuf-ts_ protoc plugin has built-in support for Angular, including:

1. a Twirp transport that uses the Angular `HttpClient`
2. a date pipe that supports `google.protobuf.Timestamp` and `google.type.DateTime`
3. annotations for dependency injection

## Usage

in web/angular project, install npm modules 
```shell
npm i @protobuf-ts/runtime @protobuf-ts/runtime-rpc @protobuf-ts/runtime-angular @protobuf-ts/twirp-transport
```

Update your `app.module.ts` with the following:

```typescript
// app.module.ts

@NgModule({
  imports: [
    // ...

    // Registers the `PbDatePipe`.
    // This pipe overrides the standard "date" pipe and adds support
    // for `google.protobuf.Timestamp` and `google.type.DateTime`.
    PbDatePipeModule,

    // Registers the `TwirpTransport` with the given options
    // and sets up dependency injection.
    TwirpModule.forRoot({
      // don't forget the "twirp" prefix if your server requires it
      baseUrl: "http://localhost:8080/twirp/",
    })

  ],
  providers: [

    // Make a service available for dependency injection.
    // Now you can use it as a constructor argument of your component.
    PolicyServerClient,
    
    // ...
  ],
  // ...
})
export class AppModule {
}
```

If you want to use a different RPC transport, you can wire it up using the `RPC_TRANSPORT` injection token. The following example uses the `GrpcWebFetchTransport` from `@protobuf-ts/grpcweb-transport`:
```typescript
// app.module.ts

@NgModule({
  // ...
  providers: [

    // Make this service available for injection in all components:
    MyServerStreamingServiceClient,

    // Configure gRPC web as transport for all services. 
    {provide: RPC_TRANSPORT, useValue: new GrpcWebFetchTransport({
        baseUrl: "http://localhost:4200"
    })},
  ],
  // ...
})
export class AppModule {
}
```

For more information, have a look at the example angular app in [packages/example-angular-app](https://github.com/timostamm/protobuf-ts/blob/master/packages/example-angular-app). 
It shows how the pipe is used, how Twirp is setup and can be run against an example gRPC-web or Twirp server (also included in the examples).
