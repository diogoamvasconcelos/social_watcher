- cloudwatch event trigger for lambda
- twitterLambda(keyword) -> searches twitter (last hour/day) and adds to ddb
  - v1: searches last hour
    - writes to ddb (batches in parallel?)
    - make `keyword` part of the stored item data (and queriable, and post it on slack)
  - v2: also searches for replies on main entries 7days
- ddbstream consumer lambda -> consumes and posts discord
  - v3: translates if not english?

- Go improvements:
  - Use the AWS SDK v2: https://aws.github.io/aws-sdk-go-v2/docs/

- CI/Dev
  - Run lambda locally
  - Unit Tests
  - Integration Tests
  - Environment Tests
  - Test coverage
  - Vulnerability scans
  - Behaviour test (cucumber?)

Pureref.com healthcheck
- similar to https://downforeveryoneorjustme.com/pureref.com
  - curl 'https://api-prod.downfor.cloud/httpcheck/http://pureref.com' -H 'User-Agent: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0' -H 'Accept: */*' -H 'Accept-Language: en-US,en;q=0.5' --compressed -H 'Referer: https://downforeveryoneorjustme.com/pureref.com' -H 'Origin: https://downforeveryoneorjustme.com' -H 'Connection: keep-alive'
    - {"statusCode":0,"statusText":"timed out","isDown":true,"returnedUrl":"http://pureref.com","requestedDomain":"http://pureref.com","lastChecked":1609600275103}
    - {"statusCode":200,"statusText":"OK","isDown":false,"returnedUrl":"http://www.google.com/","requestedDomain":"http://google.com","lastChecked":1609600320737}