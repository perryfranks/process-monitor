# Process Monitor API 

The api is communicated with basic HTTP request and json payloads. Responses from the sever are generally just acknowledgements. 

For suggested use of the API see the usage of the cli. 

For the go representation see monitorapi package. 

### Start monitoring

Start monitoring a process. This will display the information in given in the webpage 
```
/api/start
```

```
  {
    "procName": "exampleProcess",
    "workspaceName": "exampleWorkspace",
    "user": "exampleUser",
    "PID": "12345"
  },

```

Response with the id assigned to process (used to id the process in the future) and acknowledgement of receipt. 


```
  {
    "ID": 1,
    "success": true
  }

```

### End monitoring 

Ending is similar to starting. output corresponds to anything like stdout that you want captured in the front-end. 

```
/api/end
```

```
{
  "ID": 1,
  "output": "Process completed successfully.",
  "exitStatus": 0
}

```
Ack: 

```
{
  "Success": true
}

```

### Check connection 
You can no-op check your connection to the webserver via this endpoint. While the return is a bool there is little reason for it ever be false.

```
/api/checkhealth
```

This endpoint needs no payload. 



Ack: 

```
{
  "connection": true
}
```
