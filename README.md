# about

staff management system is a project to track worktime of each employee and 
calculate salary according to that data.

# service overview

## middlware service

that service accepts data from frontend using rest api and sends that data to 
db_service, also it has functionality to calculate salary according to data from database

### technologies
go, grpc

## db service
that service stores data in sql database and able to send data to logging service 
and db service using grpc

### technologies
go, gorm, postgres, grpc
