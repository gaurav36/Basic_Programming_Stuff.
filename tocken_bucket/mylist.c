/* Link list manipulation*/

#include <stdio.h>
#include <stdlib.h>
#include "mylist.h"

int My402ListInit(My402List * myList){
    myList->num_members = 0;
    myList->anchor.next = &(myList->anchor);
    myList->anchor.prev = &(myList->anchor);
    
    return TRUE;
}

int My402ListLength(My402List * myList){
    return myList->num_members;
}

int My402ListEmpty(My402List * myList){
    if(myList->num_members == 0){
        return TRUE;
    }
    else{
        return FALSE;
    }
}

int My402ListAppend(My402List * myList, void * obj){
    /* Create a new My402ListElem node */
    My402ListElem *newNode = (My402ListElem *) malloc(sizeof(My402ListElem));
    if(newNode == NULL){
        printf("malloc failed while creating My402ListElem");
        return FALSE;
    }
    
    /* Set the pointers in the new node */
    newNode->obj = obj;
    newNode->next = &(myList->anchor);
    newNode->prev = myList->anchor.prev;
    
    /* Modify the pointers of the anchor and last node accordingly */
    myList->anchor.prev->next = newNode;
    myList->anchor.prev = newNode;
    
    /* Increment the count by 1 */
    myList->num_members+=1;
    
    return TRUE;
}

int My402ListPrepend(My402List * myList, void * obj){
    
    /* Create a new My402ListElem node */
    My402ListElem *newNode = (My402ListElem *) malloc(sizeof(My402ListElem));
    if(newNode == NULL){
        printf("malloc failed while creating My402ListElem");
        return FALSE;
    }
    
    /* Set the pointers in the new node */
    newNode->obj = obj;
    newNode->next = myList->anchor.next;
    newNode->prev = &(myList->anchor);
    
    /* Update the pointers in anchor */
    myList->anchor.next->prev = newNode;
    myList->anchor.next = newNode;
    
    /* Increment the count by 1 */
    myList->num_members+=1;
    
    return TRUE;
}

void My402ListUnlink(My402List * myList, My402ListElem * node){
    
    /* Update the pointers of the nodes adjacent to node */
    node->prev->next = node->next;
    node->next->prev = node->prev;
    
    /* Free the memory occupied by that node */
    free(node);
    
    /* Decrement the count by 1 */
    myList->num_members-=1;
}

void My402ListUnlinkAll(My402List * myList){
    
    /* If the list is not empty then unlink and delete the nodes */
    if(myList->num_members > 0){
        My402ListElem *anchor = &(myList->anchor);
        anchor->prev->next = NULL;
        My402ListElem *temp = anchor->next;
        My402ListElem *next = NULL;
        
        /* Separate the anchor */
        anchor->next = anchor;
        anchor->prev = anchor;
        
        /* Delete the remaining nodes */
        while(temp !=NULL){
            next = temp->next;
            free(temp);
            temp = next;
        }
        
        myList->num_members = 0;
    }
}

int My402ListInsertBefore(My402List * myList, void * obj, My402ListElem * node){
    if(node == NULL){
        
        /* Call the prepend method */
        return My402ListPrepend(myList, obj);
    }
    else{
        /* Create a new node and update pointers accordingly */
        My402ListElem *newNode = (My402ListElem *) malloc(sizeof(My402ListElem));
        if(newNode == NULL){
            printf("malloc failed while creating My402ListElem");
            return FALSE;
        }
        newNode->obj = obj;
        newNode->next = node;
        newNode->prev = node->prev;
        node->prev->next = newNode;
        node->prev = newNode;
        
        /* Increment the count by 1 */
        myList->num_members+=1;
        
        return TRUE;
    }
}

int My402ListInsertAfter(My402List * myList, void * obj, My402ListElem * node){
    if(node == NULL){
        
        /* Call the append method */
        return My402ListAppend(myList, obj);
    }
    else{
        /* Create a new node and update pointers accordingly */
        My402ListElem *newNode = (My402ListElem *) malloc(sizeof(My402ListElem));
        if(newNode == NULL){
            printf("malloc failed while creating My402ListElem");
            return FALSE;
        }
        newNode->obj = obj;
        newNode->next = node->next;
        newNode->prev = node;
        node->next->prev = newNode;
        node->next = newNode;
        
        /* Increment the count by 1 */
        myList->num_members+=1;
        
        return TRUE;
    }
}

My402ListElem * My402ListFirst(My402List * myList){
    if(My402ListEmpty(myList) == TRUE){
        return NULL;
    }else{
        return myList->anchor.next;
    }
}

My402ListElem * My402ListLast(My402List * myList){
    if(My402ListEmpty(myList) == TRUE){
        return NULL;
    }else{
        return myList->anchor.prev;
    }
}

My402ListElem * My402ListNext(My402List * myList, My402ListElem * node){
    if(node->next == &(myList->anchor)){
        return NULL;
    }
    else{
        return node->next;
    }
}
My402ListElem * My402ListPrev(My402List * myList, My402ListElem * node){
    if(node->prev== &(myList->anchor)){
        return NULL;
    }
    else{
        return node->prev;
    }
}

My402ListElem * My402ListFind(My402List * myList, void * obj){
    My402ListElem * node = NULL;
    My402ListElem * temp = NULL;
    for (node = My402ListFirst(myList); node != NULL; node = My402ListNext(myList, node)) {
        
        if(node->obj == obj){
            temp = node;
            break;
        }
    }
    return temp;
}















