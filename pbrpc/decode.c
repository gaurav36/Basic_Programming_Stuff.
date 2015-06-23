#include <stdio.h>
#include "main.pb-c.h"

#define MAX_MSG_SIZE 4096

enum {
	ADD_REQ = 1,
	MULT_REQ
};

int main (int argc, const char *argv[])
{
	DMessage *msg;
	Rpcheader *header;

	char len;
	uint8_t buf[MAX_MSG_SIZE];

	//fread(&len, 1, 1, stdin);

	printf ("\npayload len = %d\n", len);

	fread (buf, 1, len, stdin);

	msg = dmessage__unpack (NULL, len, buf);
	if (msg == NULL)
	{
		printf ("null");
		fprintf (stderr, "error unpacking incoming messages\n");
		return 1;
	}
	header = msg->header;

	printf ("msg a is %d and msg b is %d\n", msg->a, msg->b);

	switch (header->rpcreq) {
		case ADD_REQ:
			printf ("\n %d+%d=%d\n", msg->a, msg->b, msg->a + msg->b);
			break;
		case MULT_REQ:
			printf ("\n%d*%d=%d\n", msg->a, msg->b, msg->a * msg->b);
			break;
	}

	dmessage__free_unpacked (msg, NULL);

	return 0;
}
