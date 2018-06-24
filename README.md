This is a simple command line tool written in Go, that sends a full schema dump request query to a  graphql query and shows it in a format similar to the GraphQL schema language used at graphql.org and in some graphql server implementations.

graphqlschema \<url\> [-debug]
  
add a -debug flag to dump the raw json received from the server.

the output is color formatted to be similar to the examples on graphql.org, but is optimzed more for display on a dark background.

It doesn't yet:
* Handle nested types
* Show ! and [] properly if there's a list of non-nulls instead of a non-null list.
* Handle interfaces or unions

![Alt text](/output.png?raw=true "Output")
