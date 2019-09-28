# Confessional-Christianity
The historic creeds and confessions of the Christian faith are of great interest to me as a Christian. This project is mostly intended
to build and demonstrate my skills as a Software Engineer, but also to provide interesting and helpful data concerning these historic
and time-tested documents.

## Aren't there already websites out there like this?
Yes and no. To my knowledge, there is not a public API returning JSON in a modern web-development format. The websites I've looked at (in the
network tab in the Chrome Dev Console!) appear to be making calls which return HTML pages with the data embedded. They could be server-side-rendering a modern-web application, but it appears to me they are all serving their data in such a way that it is coupled to a certain HTML structure.

What I intend to do, is return strictly JSON via HTTP, GRPC, and graphQL. The benefit of doing this, of course, is that the client can determine what to do with the data, so it is not coupled to HTML etc...

## The Intent of this REPO
First, I will implement a basic HTTP server responding to RESTful requests in JSON for the Westminster Confession of Faith.
Second, I will implement a basic gRPC server responding to RPC requests.

### Tech Used
- GoKit
- DynamoDB

## Forthcoming REPOs
- Node GraphQL API: a GraphQL API will be created for the front end to consume.
- Data Visualization in Python: I intend to interact with the gRPC server in a python project to extract, load, and transform the data into a new back end.
- SSR React: I intend to use NextJS, Apollo, and React to render a modern, performant, as a stand-alone front end.

## Hosting & Deployment
- To implement a smooth deployment process, I will be using CircleCI and the boto3 AWS python package for automatic deployment.
- For hosting, I am not completely sure, but DynamoDB and Lambda + API Gateway appear to be the most affordable options so far.

If you are interested in contributing to this project, please reach out to me `@MNK5490` on twitter!
