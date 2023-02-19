# ipify-go

An unofficial and very simple Go library for using [ipify API](https://www.ipify.org/).

## Features
Get IP v4 and IP v6 in different formats such as:

### Text
```
127.0.0.1
```

### JSON
  ```json
    {"ip": "127.0.0.1"}
  ```

### JSONP
```jsonp
callback({"ip": "127.0.0.1"})
```

### JSONP with custom callback

```jsonp
my_function({"ip": "127.0.0.1"})
```
