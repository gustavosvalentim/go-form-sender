# go-form-sender

Send forms somewhere.

This project was built to facilitate sending request forms through email.

## Build

`go build ./cmd/go-form-sender`

## Setup

Set your templates directory

`go-form-sender set templates_dir <your-templates-dir>`

## Usage

Submit a form

`go-form-sender send <form-name>`

This command will ask you to fill the form before sending.
