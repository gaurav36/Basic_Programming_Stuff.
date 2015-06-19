CC = gcc
CFLAGS = -Wall

main: token-main token-packet-arrival token-helper token-token-arrival \
      token-server token-monitor token-statistics mylist
	$(CC) -o token -g -pthread -lm token-core.o token-packet-arrival.o \
	token-helper.o token-token-arrival.o token-server.o token-monitor.o \
	token-statistics.o mylist.o

token-main:
	$(CC) -g -c $(CFLAG) -pthread token-core.c
token-packet-arrival:
	$(CC) -g -c $(CFLAG) -pthread token-packet-arrival.c
token-helper:
	$(CC) -g -c $(CFLAG) -pthread token-helper.c
token-token-arrival:
	$(CC) -g -c $(CFLAG) -pthread token-token-arrival.c
token-server:
	$(CC) -g -c $(CFLAG) -pthread token-server.c
token-monitor:
	$(CC) -g -c $(CFLAG) -pthread token-monitor.c
token-statistics:
	$(CC) -g -c $(CFLAG) -pthread token-statistics.c
mylist:
	$(CC) -g -c $(CFLAG) mylist.c
clean:
	rm -f *.o token
