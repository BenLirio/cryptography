#include <unistd.h> 
#include <stdio.h> 
#include <sys/socket.h> 
#include <stdlib.h> 
#include <netinet/in.h> 
#include <string.h> 

#define PORT 8080

int main(int argc, const char* argv[]) {
	int server_fd, new_socket, valread;
	struct sockaddr_in address;
	int opt = 1;
	char buffer[1024] = {0};
	char *hello = "Hello from server";

	if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) == 0)
	{
		perror("socket failed");
		exit(EXIT_FAILURE);
	}
	if (setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR | SO_REUSEPORT, &opt, sizeof(opt))) 
	{
		perror("setsockopt");
		exit(EXIT_FAILURE);
	}
	return 0;
}
