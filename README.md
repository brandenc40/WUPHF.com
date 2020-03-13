# WUPHF.com

### You must add a secrets.yaml file to the config/ folder with the following:

```
twilio:
  account_sid: <TWILIO SID>
  auth_token: <TWILIO AUTH TOKEN>
  phone_number: <TWILIO PHONE NUMBER>

gmail:
  email: <EMAIL ADDRESS>
  password: <PASSWORD>
```

### You'll also need a twIML template for the phone call. This is what I use:

```
<?xml version="1.0" encoding="UTF-8"?>
<Response>
<Say>
Woof from {{from_name}}, {{message}}.
</Say>
</Response>
```
