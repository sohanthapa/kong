# SERVICE API
This project implements the Services API that can be used to retrieve services from back end
## Tech stack used
Golang, mongodb atlas (cloud)
## How to run this project
NOTE: You must have Golang and mongodb installed in your machine.  <br />
   1.to build the project use "go build main.go"  <br />
   2. to run the project use "./main"  <br />
   3. to unit test this project use "go test ./..." <br />
   4. You can use Postman or curl command to send requests while the project is running
## API reference
We have the following two APIs for the server <br />
    1. GET "/services/{serviceID}" - need to provide the serviceID to retrieve that specific service document <br />
    2. GET "/services?name={name}&sort=desc&offset={offset}&limit={limit} - need to provide the int values for the offset and limit <br />
        &emsp;&emsp; 1. offset and limit is for pagination. <br /> 
        &emsp;&emsp; &emsp;&emsp; example : if the FE only wants to get the first 4 documents then, we can use offset as 0 and limit as 4 <br />
        &emsp;&emsp; 2. name - is used as filter for searching. <br /> 
        &emsp;&emsp; 3. sort - use desc value. By default, it is sorted in asc order <br /> 
    3.
## Assumptions
I have made below assumptions for this project: <br />
&emsp;&emsp; 1. Created services and versions and stored in db while the server loads. <br />
&emsp;&emsp; 2. Used in-memory users for authorization and aunthentication. <br />
&emsp;&emsp; 3. Can only use name (from Service object) as search filter. <br />
&emsp;&emsp; 4. Can only sort using name (from Service object) for sorting in ascending or descending order. <br />
&emsp;&emsp; 5. Stored all the version objects inside each service. <br />
&emsp;&emsp; 6. Do not need to implement any APIs for Versions (the exercise did not specify). <br />
&emsp;&emsp; 7. The FE or any one that makes a request to BE has full freedom in terms of pagination. That means <br />
&emsp;&emsp;&emsp;&emsp; they will provide the offset and limit value to get how many documents they want as per their requirement.
&emsp;&emsp;&emsp;&emsp; BE does not have any business logic to figure the offset and limit.

## Notes
I have added few comments with NOTES in the code to clarify few things and why I choose to code it that way. <br />
Also, for the simplicity of this exercise I choose not unit test database functions because I need to mock the database call which <br />
will involve more work. I have stated this in my code.

## Screenshots 
Below are few screenshots of the results of the APIs using Postman:
![image](https://user-images.githubusercontent.com/7494475/140621180-c1c4e97c-80e0-43a3-ad06-fae888ebdc19.png)
![image](https://user-images.githubusercontent.com/7494475/140621232-82aabf0f-2be2-4e3b-85df-0d5f50231141.png)
![image](https://user-images.githubusercontent.com/7494475/140621267-80592d44-2b87-4214-86d1-bfd5858c6d1e.png)




                                           
                                                 
