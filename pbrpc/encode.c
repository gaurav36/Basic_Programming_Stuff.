#include <stdio.h>
#include <stdlib.h>
#include "main.pb-c.h"

enum {
	ADD_REQ = 1,
	MULT_REQ
};

int main (int argc, const char * argv[])
{
	DMessage msg = DMESSAGE__INIT;
	Rpcheader header = RPCHEADER__INIT;

	uint8_t *buf;
	unsigned len;

	header.rpcreq = MULT_REQ;
	msg.header = &header;

	len = dmessage__get_packed_size (&msg); // length of packet
	buf = (uint8_t*)malloc (len + 1);
	buf[0] = len;

	msg.a = 10;
	msg.b = 20;
	msg.len = len;

	dmessage__pack (&msg, &buf[1]);	// packing message

	fprintf (stderr, "Writing %d serialized bytes\n", len);

	fwrite (buf, 1, len+1, stdout);

	free(buf);

	return 0;
}
