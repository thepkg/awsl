package awsl
/*

Package https://github.com/thepkg/awsl provides helpful functions to work with AWS Lambda.

* **FromDynamoDBMap** - converts data from AWS Lambda Event for DynamoDB to regular map
(from `map[string]events.DynamoDBAttributeValue` to `map[string]interface{}`):

	func main() {
		lambda.Start(Handler)
	}

	func Handler(event events.DynamoDBEvent) {
		for _, record := range event.Records {
			rec := FromDynamoDBMap(record.Change.NewImage)
			// rec contains map[string]interface{} so it's easier to work with data!
		}
	}

*/
