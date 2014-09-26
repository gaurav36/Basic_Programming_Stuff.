#include <stdio.h>
#include <stdlib.h>

typedef struct POOL
{
        char *next;
        char *end;
}pool;

pool *pool_create (size_t size)
{
        pool *p = (pool*)malloc (size + sizeof (pool));
        p->next = (char*)&p[1];
        p->end  = p->next + size;

        return p;
}

void pool_destroy (pool *p) {
        free (p);
        p = NULL;
}

size_t pool_available (pool *p) {
        return p->end - p->next;
}

void *pool_alloc (pool *p, size_t size) {
        if (pool_available(p) < size) {
                printf ("oops!!! No memory left in mempool\n");
                return NULL;
        }
        void *mem = (void *)p->next;
        p->next +=size;
        printf ("memeory successfully allocated from pool\n");
        return mem;
}

int main ()
{
        pool *my_pool = pool_create (100);
        pool *use1,*use2,*use3,*use4;

        printf ("pool_available is %d\n\n",pool_available(my_pool));

        use1 = pool_alloc (my_pool, 30);
        printf ("pool_available is %d\n\n",pool_available(my_pool));

        use2 = pool_alloc (my_pool, 30);
        printf ("pool_available is %d\n\n",pool_available(my_pool));

        use3 = pool_alloc (my_pool, 30);
        printf ("pool_available is %d\n\n",pool_available(my_pool));

        use4 = pool_alloc (my_pool, 50);
        printf ("pool_available is %d\n\n",pool_available(my_pool));

        return 0;
}
