# Go Image Upload/Resizer Server

* go 1.13+
* [gin-gonic/gin](https://github.com/gin-gonic/gin)
* [disintegration/imaging](https://github.com/disintegration/imaging)
### Installation & Run:
 
* Before running, you should set the projectPath information on  ***server/main.go***. 

Then on terminal;
```
scripts/run.sh
```
## API
#### /upload-file
* `POST` : Body:{ "data":FILE }
#### /files/:width/:height/filename
* `GET` : Static folder.
* `example`: /files/1000/500/KGkP4cJQKDzxcsuM7xcL23.png