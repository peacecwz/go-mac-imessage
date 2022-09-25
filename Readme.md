# Mac SMS Tracker

Easy way to send and receive SMS/iMessage on Mac OS

## How it works

If you are using your same iCloud account on Mac and iPhone you can use it easily. iCloud can sync your message between iPhone, Mac and other devices on iMessage application. [See how to configure and details about this feature](https://support.apple.com/en-gb/guide/messages/icht8a28bb9a/mac)

## Use as CLI

### Send SMS/iMessage to Phone Number

```bash

$ sms send -m "Hello, World" -t "+310638122497"

```

### Receive SMS/iMessages

```bash

$ sms receive

```

And also you can set intervals like

```bash

$ sms receive -i 3000

```

## Use as Library

### Send SMS/iMessage

```go

err := sms.Send("Hello, World", "+31063812249")
if err != nil {
    log.Fatal(err)
    return
}

fmt.Println("Sent")

```

### Receive SMS/iMessage

```go
interval := 3 * 1000
err := sms.TrackSMS(interval, func(sms []sms.SMS) {
    for _, s := range sms {
        fmt.Printf("Message: %s from %s\n", s.Content, s.From)
        err := s.Read()
        if err != nil {
            log.Fatalln(err)
        }
    }
})

```

## License

Licensed under the MIT License
