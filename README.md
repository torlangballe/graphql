NOTE: Not quite compiling yet, a utility dependency slipped in.

This is a simple command line tool written in Go, that sends a full schema dump request query to a  graphql query and shows it in a format similar to the GraphQL schema language used at graphql.org and in some graphql server implementations.

graphql \<url\> [-debug]  
add a -debug flag to dump the raw json received from the server.

![Alt text](/output.png?raw=true "Output")

the output is color formatted to be similar to the examples on graphql.org, but is optimzed more for display on a dark background.

It ignores all the built-in types.

It:
* Doesn't Handle nested types yet.
* Doesn't Show ! and [] properly if there's a list of non-nulls instead of a non-null list yet.
* Doesn't Handle interfaces or unions yet.
* Isn't tested with 'ID' type yet.
