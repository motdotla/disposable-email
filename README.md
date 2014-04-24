# disposable-email

Create and host your own disposable email account. Great for preventing SPAM when you signup for questionable websites. Inspiration from the very excellent [mailinator](http://mailinator.com/). 

Try the [demo](http://disposable-email.herokuapp.com/).

![](https://raw.githubusercontent.com/scottmotte/disposable-email/master/public/images/disposable-email.gif)

## About

Dispoable Email uses [SendGrid's Inbound Parse Webhook](http://sendgrid.com/docs/API_Reference/Webhooks/parse.html) & [Runscope's RequestBin](http://requestb.in) to quickly standup your own personal disposable email.

It's easy to setup and use. You just need [Go](http://golang.org).

## Usage

```
git clone https://github.com/scottmotte/disposable-email.git
cd disposable-email
go get
go run app.go
```

Visit [http://localhost:3000](http://localhost:3000)

That's it. You're done!

## Deploy to Heroku

```
git clone https://github.com/scottmotte/disposable-email.git
cd disposable-email
heroku create -b https://github.com/kr/heroku-buildpack-go.git 
git push heroku master
heroku open
```

That's it. You're done!
