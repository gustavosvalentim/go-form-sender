# go-form-sender

Send forms somewhere.

This project was built to facilitate sending request forms through email.

## Build

In `goformsender.go` configure your SMTP credentials by changing

```go
smtp_server_name   string = "smtp.gmail.com"
smtp_server_port   int    = 587
smtp_user_name     string = "email@test.com"
smtp_user_password string = "*****"
```

`go build goformsender.go`

## Usage

Copy the sample form from `samples/simple-template.yaml` to `~/.formtemplates`.

Then run `goformsender <form-name>` fill and send the form.
