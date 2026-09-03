- sudo apt install -y redis-server redis-tools 
- sudo systemctl enable  redis-server  
-  sudo systemctl status redis-server  
-  redis-cli ping 


- my one doubt is when i ran the code with the installed value kaushik it ran succefully ,and it main function ended then in redi-slli when ran the query command i got the values stored there too  . 

- The answer is go programm and redis are two seprate process .
  client.Set(ctx, "user:1:name", "Kaushik", time.Minute)

  sends the data to the Redis server. Redis stores it in its own memory.

  When main() ends:

  defer client.Close()

  only closes the Go client connection. It does not delete the Redis key.

  Therefore, you can still access the data using:
