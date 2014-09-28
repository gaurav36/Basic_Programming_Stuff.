#include <stdio.h>
#include <stdlib.h>

typedef struct mem_pool
{
        char *pool_next; //Starting address of pool
        char *pool_end; //Ending address of pool
	int  hot_count;
	int  cold_count;
}pool;

pool *pool_create (size_t size)
{
        pool *p = (pool*)malloc (size + sizeof (pool));
        p->pool_next = (char*)&p[1];
        p->pool_end  = p->pool_next + size;

        return p;
}

void pool_destroy (pool *p) {
        free (p);
        p = NULL;
}

size_t pool_available (pool *p) {
        return p->pool_end - p->pool_next;
}

void *pool_alloc (pool *p, size_t size) {
        if (pool_available(p) < size) {
                printf ("oops!!! No memory left in mempool\n");
                return NULL;
        }
        void *mem = (void *)p->pool_next;
        p->pool_next +=size;
        printf ("memeory successfully allocated from pool\n");
        return mem;
}

int main ()
{
        pool *my_pool = pool_create (100);
        pool *use1,*use2,*use3,*use4;

        printf ("pool_available is %zu\n\n",pool_available(my_pool));

        use1 = pool_alloc (my_pool, 30);
        printf ("pool_available is %zu\n",pool_available(my_pool));
	printf ("use1 pool starting address is %x\n\n",&use1->pool_next);

        use2 = pool_alloc (my_pool, 30);
        printf ("pool_available is %zu\n",pool_available(my_pool));
	printf ("use2 pool starting address is %x\n\n",&use2->pool_next);

        use3 = pool_alloc (my_pool, 30);
        printf ("pool_available is %zu\n",pool_available(my_pool));
	printf ("use3 pool starting address is %x\n\n",&use3->pool_next);

        use4 = pool_alloc (my_pool, 50);
        printf ("pool_available is %zu\n\n",pool_available(my_pool));

        return 0;
}
