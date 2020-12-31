# Configuration (installation - Ubuntu)

## Terraform
Extract the .zip from the latest terraform build and move it to `/usr/bin`. Try running `terraform` in the terminal. If you can't, make sure the `/usr/bin` was added to the #PATH.

## AWS CLI
```
apt install awscli
```

## GO
https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
```
go get github.com/aws/aws-lambda-go/lambda
go get github.com/aws/aws-lambda-go/events
go get github.com/bwmarrin/discordgo
```

### GOLANG tuts
- https://golang.org/doc/tutorial/create-module
- https://tour.golang.org/welcome/1
- https://golang.org/doc/effective_go.html

## Discord bot
  - guide: https://medium.com/@thomlom/how-to-create-a-discord-bot-under-15-minutes-fb2fd0083844
  - client/sdk: https://github.com/bwmarrin/discordgo