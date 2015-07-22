# elasticprompt

> A command-line client for Elasticsearch.

This is pre-alpha at the moment. It uses the [https://github.com/olivere/elastic](https://github.com/olivere/elastic) library to connect to Elasticsearch.

My understanding of Elasticsearch is pre-alpha as well.

## Example Usage

```
$ elasticprompt
Connecting to http://localhost:9200/...
localhost:9200 () > port 9201
Connecting to http://localhost:9201/...
localhost:9201 () > index bank
localhost:9201 (bank) > search
// results here
localhost:9201 (bank) > quit
$
```
