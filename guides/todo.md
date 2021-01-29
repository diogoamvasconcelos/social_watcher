- twitterLambda(keyword) -> searches twitter (last hour/day) and adds to ddb

  - searches last hour
    - writes to ddb (batches in parallel?)
    - make `keyword` part of the stored item data (and queriable, and post it on slack)
  - also searches for replies on main entries that are active (having a reply in the last x days)
  - posts author follower count

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

- Reddit search

  - run once a day (very few posts)
  - or native https://www.reddit.com/dev/api/#GET_search (not sure if searhes for comments)
  - use https://redditsearch.io/ (not sure if legal)

- compitcao search

  - o outro programa pro ipad

- Website Healthcheck
  - if has to retry for a few times the last X minuts, send warning to channel (without @here)
