# personal-mailinator

Host your own mailinator.

![](https://raw.githubusercontent.com/scottmotte/personal-mailinator/master/public/images/personal-mailinator.gif)

Personal Mailinator uses [SendGrid's Inbound Parse Webhook](http://sendgrid.com/docs/API_Reference/Webhooks/parse.html) & [Runscope's RequestBin](http://requestb.in) to quickly standup your own personal mailinator.

It's easy to setup and use. You just need [Go](http://golang.org).

## Usage

```
git clone https://github.com/scottmotte/personal-mailinator.git
cd personal-mailinator
go get github.com/martini-go/martini
go get github.com/martini-contrib/render
go run app.go
```

Visit [http://localhost:3000](http://localhost:3000)

That's it. You're done!
