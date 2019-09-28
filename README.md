# Confessional-Christianity
The historic creeds and confessions of the Christian faith are of great interest to me as a Christian. This project is mostly intended
to build and demonstrate my skills as a Software Engineer, but also to provide interesting and helpful data concerning these historic
and time-tested documents.

## Aren't there already websites out there like this?
Yes and no. To my knowledge, there is not a public API returning JSON in a modern web-development format. The websites I've looked at (in the
network tab in the Chrome Dev Console!) appear to be making call which return with HTML which have the data embedded. They could be server-side
rendering a modern-web application, but it appears to me they are all serving their data in such a way that it is married to a certain HTML
structure.

What I intend to do, is return strictly JSON via HTTP, GRPC, and graphQL. The benefit of doing this, of course, is that the client can determine
what to do with the data, so it is not coupled to HTML etc...

## The Intent of this REPO
This repo is intended to be the grand-parent of all other forthcoming repos. In it, I will implement a basic HTTP server responding to
RESTful requests in JSON. I chose golang because at the time I began this project, I was working with Go and found it an intersting language.

### Tech Used
- GoKit
- DynamoDB

## Forthcoming REPOs
- Node GraphQL: a GraphQL API will be created for the front end to consume.
- SSR React: I intend to use NextJS, Apollo, and React to render a modern, performant, stand-alone front end.
- Data Visualization in Python: I intend to use the gRPC server in a python project to extract, load, and transform the data into a new back end.

## Hosting & Deployment
- To implement a smooth deployment process, I will be using CircleCI and the boto3 AWS python package for automatic deployment.
- For hosting, I am not completely sure, but DynamoDB and Lambda + API Gateway appear to be the most affordable options so far.

If you are interested in contributing to this project, please reach out to me `@MNK5490` on twitter!
