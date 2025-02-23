# Logger

Generric JSON Formatted logger written in Go

## Usage

```Go
var debug = flag.Bool("debug", false, "")

init (
  logger.SetAppName("app_name")
)

func main() {
  flag.Parse()
  if *debug {
    logger.SetLevel(logger.DebugLevel)
  }
  ...

  logger.WithFields(Fields{
    "animal": "walrus",
  }).Print("hello")
}

```

## API List

```Go

SetAppName("appname") // applicaton name
SetLevel(InfoLevel) // default: InfoLevel

Type("hoge") // set "type": "hoge"
WithFields(Fields{}) // set data fields

Print("hoge")
Printf("%s","hoge")
Println("hoge")

Debug, Debugf, Debugln
Info, Infof, Infoln
Warn, Warnf, Warnln
Error, Errorf, Errorln
Fatal, Fatalf, Fatalln
Panic, Panicf, Panicln

```


## Sentry Message
To use sentry logger, please initialize as the below, you need to pass the environemnt(local/develop/stg/prod)

```go
   env := util.GetEnv()
	logger.InitSentry(env)
```

Now you can send the log to sentry.
```go
	logger.LogToSentry(fmt.Sprintf("Successfully updated portfolio chart for user %d", userID), "INFO")
```
## TODO

- Logger factory for different logrus hooks (ElasticSearch, for example)
