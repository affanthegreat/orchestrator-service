#  Task orchestration

This is the code follow up for the home exercise.

## steps to run:
- For part 2 and part 3 folders, copy "getuserbyname" and "mockdata" folders into your GOPATH
- In client folder of each part 2 and part 3, run server.go file to start the RPC server and client.go for RPC Client
- open http://localhost:8080/getUser/name and http://localhost:10000/getUser/name  for testing each server respectively.
- For part 3 (Orchestration process) copy all the folders in"Proto" folder into your GOPATH
- In logic folder, from each  orchestrator folder,  run each orchestrator file to start the respective servers. 
--  orchestrator 1 contains the implementation for 'GetUserByName' function and for the purposes of demo, There are four names {"James", "Wagner", "Christene", "Mike"} to replicate the find user by name. if given name exists in the list, then Orchestrator 1 will call 'GetUser' function of Orchestrator 2 operating on port 9001.
-- Orchestrator 2 contains implementation for GetUser function which then connects to 'GetMockUserdata' function of Orchestrator 3. As given, it returns a user object with given name as name, Class as length of string name and roll as length of string name multiplied by 10.



