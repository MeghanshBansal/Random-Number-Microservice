# Random-Number-Microservice
A microservice with two routes to generate random numbers

The microservice has two one routes "/"
If the method is GET it will generate only one random number

If the method is POST it will generate N random numbers.
The body of the request will take 
``` 
{
  "count": N
}
```
as parameter to generate those numbers
