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

## TODO

- Logger factory for different logrus hooks (ElasticSearch, for example)
