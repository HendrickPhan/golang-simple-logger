# golang-simple-logger


# Overview
This is a simple golang logger package

# Usage
## Config
### 1. Flag:
- Only message < flag will be log  
  
### 2. Output:
- List of output stream. Default is terminal
- Check test for add log to file  
  
### 3. MessageChanSize:
- Size of chan that hold message for async write log

## Sample

```
    f, _ := os.OpenFile("./test_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    config := &logger.LoggerConfig{
        Flag:            logger.FLAG_DEBUGP,
        Outputs:         []*os.File{os.Stdout, f}, // For both file and terminal
        MessageChanSize: 1000,
    }
    logger.SetConfig(config)
    logger.Debug("Log debug")
```


