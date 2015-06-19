/* main function for parsing input data.*/
/* Default value of:
 * Bucket capicity           B:      10;
 * Token arrival time        r:      1.5;
 * Packet service time       mu:     0.35;
 * Packet arrival time       lambda: 0.5;
 * Token needed for a packet P:      3;
 * Total number of packet is num:    20;
 * */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <pthread.h>
#include "mylist.h"
#include <sys/time.h>
#include <unistd.h>
#include <errno.h>
#include "token-header.h"
#include <signal.h>
#include <math.h>

My402List Q1;                                                                   
My402List Q2;                                                                   
                                                                                
/* Emulation related data */                                                    
long  **tsfile_data;                                                            
long  bucket_capacity;                                                          
float token_arrival_time;                                                       
float token_arrival_time;                                                       
float packet_arrival_time;                                                      
float packet_service_time;                                                      
long  total_number_of_packets;                                                  
long  tokens_needed_for_a_packet;                                               
                                                                                
int  cntrl_c_signal;                                                            
int  server_thread_stop;
int  is_trace_driven_mode;                                                      
int  packet_arrival_thread_stop;                                                
long number_of_tokens_in_bucket;                                                
long number_of_packets_put_in_Q1;                                               
                                                                                
/* Statistics related data */                                                   
float token_drop_probability;                                                   
float packet_drop_probability;                                                  
float average_packet_service_time;                                              
float average_number_of_packets_at_S;                                           
float average_number_of_packets_in_Q1;                                          
float average_number_of_packets_in_Q2;                                          
float average_packet_inter_arrival_time;                                        
float average_time_a_packet_spent_in_system;                                    
float standard_deviation_for_time_spent_in_system;                              
                                                                                
sigset_t set;                                                                   
struct timeval emulation_start_time;                                            
                                                                                
pthread_mutex_t mutex             = PTHREAD_MUTEX_INITIALIZER;                  
pthread_cond_t condition_variable = PTHREAD_COND_INITIALIZER;

pthread_t server_thread_id;                                                     
pthread_t monitor_thread_id;                                                    
pthread_t packet_arrival_thread_id;                                             
pthread_t token_depositing_thread_id;                                           
                                                                                
void initialize () {                                                            
                                                                                
    tsfile_data                                 = NULL;                           
    cntrl_c_signal                              = 0;                            
    bucket_capacity                             = 0;                              
    server_thread_stop                          = 0;                            
    token_arrival_time                          = 0.0;                          
    is_trace_driven_mode                        = 0;                            
    token_drop_probability                      = -1.0;                         
    packet_drop_probability                     = -1.0;                         
    total_number_of_packets                     = 0;                              
    number_of_tokens_in_bucket                  = 0;                            
    packet_arrival_thread_stop                  = 0;                            
    average_packet_service_time                 = -1.0;                         
    average_number_of_packets_at_S              = -1.0;                         
    average_number_of_packets_in_Q1             = -1.0;                         
    average_number_of_packets_in_Q2             = -1.0;                         
    average_packet_inter_arrival_time           = -1.0;                         
    average_time_a_packet_spent_in_system       = -1.0;                         
    standard_deviation_for_time_spent_in_system = -1.0;                         
                                                                                
    /* Link list initilization*/                                                
    My402ListInit(&Q1);                                                         
    My402ListInit(&Q2);                                                         
                                                                                
    sigemptyset(&set);                                                          
}

void process(float lambda, float mu, float r, long B,                           
             long P, long num, char * tsfile) {                                 
                                                                                
    sigaddset (&set, SIGINT);                                                   
    pthread_sigmask(SIG_BLOCK, &set, 0);                                        
                                                                                
    if(tsfile != NULL){                                                         
        bucket_capacity      = B;                                               
        token_arrival_time   = 1/r;                                             
        is_trace_driven_mode = 1;                                               
                                                                                
        if(token_arrival_time > 10) token_arrival_time = 10.0;                  
    }                                                                           
    else {                                                                      
        bucket_capacity            = B;                                         
        token_arrival_time         = 1 / r;                                     
        packet_arrival_time        = 1 / lambda;                                
        packet_service_time        = 1 / mu;                                    
        total_number_of_packets    = num;                                       
        tokens_needed_for_a_packet = P;                                         
                                                                                
        if(packet_arrival_time > 10) packet_arrival_time = 10.0;                
        if(packet_service_time > 10) packet_service_time = 10.0;                
        if(token_arrival_time > 10) token_arrival_time = 10.0;                  
    }                                                                           
                                                                                
    if(is_trace_driven_mode){                                                   
        FILE *fp = NULL;                                                        
        fp = fopen(tsfile, "r");                                                
        if (fp == NULL) {                                                       
            fprintf(stderr, "\nThe input file %s cannot be opened for reading " 
                    "or it doesn't exist.\n", tsfile);                          
            exit(1);                                                            
        }                                                                       
                                                                                
        digest_file(fp);                                                          
    }                                                                           
                                                                                
    printf("Emulation Parameters:\n");                                          
    printf("Total number of packet to arrive = %ld\n",total_number_of_packets); 
                                                                                
    if(!is_trace_driven_mode){                                                  
        printf("Packet arrival time lambda = %f\n",lambda);                     
    }                                                                           
    if(!is_trace_driven_mode){                                                  
        printf("Packet service time  mu = %f\n",mu);                            
    }                                                                           
    if(!is_trace_driven_mode){                                                  
        printf(" Tokens needed for packet P = %ld\n\n",P);                      
    }                                                                           
                                                                                
    if(is_trace_driven_mode){                                                   
        printf("Test file name for processing input data  = %s\n\n",tsfile);    
    }                                                                           
                                                                                
    printf("Tocken arrival time r = %f\n",r);                                   
    printf("Bucket capicity  B = %ld\n",B);                                     
                                                                                
    gettimeofday(&emulation_start_time, NULL);                                  
    printf("00000000.000ms: emulation begins\n");                               
                                                                                
    pthread_create(&packet_arrival_thread_id, 0,                                
                   packet_arrival_thread_function, 0);                          
                                                                                
    pthread_create(&token_depositing_thread_id, 0,                              
                   token_depositing_thread_function, 0);                        
                                                                                
    pthread_create(&server_thread_id, 0, server_thread_function, 0);            
                                                                                
    //pthread_create(&monitor_thread_id, 0, monitor, 0);                          
                                                                                
    pthread_join(packet_arrival_thread_id, NULL);                               
    pthread_join(token_depositing_thread_id, NULL);                             
    pthread_join(server_thread_id, NULL);        
                               
    display_statistics();                                                         
}                                

int main (int argc, char *argv[])
{
    float lambda = 0.5;   /*Packet arriaval time*/
    float mu     = 0.35;  /*Packet service time*/
    //float r      = 0.2;   /*Token arrival time*/
    float r      = 1.5;   /*Token arrival time*/
    long  B      = 5;    /*Bucket capicity*/ 
    long  P      = 3;     /*Token needed for a packet*/
    long  num    = 5;    /*Total number of packet*/
    char *tsfile = NULL;  /*Name of test file*/
    int   i      = 1;

    for(i=1;i<argc;i++){
        /* Packet arrival time is 1/lambda */
        if(strcmp("-lambda", argv[i])==0){
            if(i!=argc && strstr(argv[i+1],"-")==NULL){
                lambda = atof(argv[i+1]);
            }
            else{
                fprintf(stderr, "Please provide a valid value with -lambda "
		        "argument, else remove -lambda to use default value\n");
                exit(1);
            }
        }

        /* Packet service time is 1/mu */
        if(strcmp("-mu", argv[i])==0){
            if(i!=argc && strstr(argv[i+1],"-")==NULL){
                mu = atof(argv[i+1]);
            }
            else{
                fprintf(stderr,"Please provide a valid value with -mu "
		        "argument, else remove -mu to use default value\n");
                exit(1);
            }
        }

        /* Token arrival time is 1/r */
        if(strcmp("-r", argv[i])==0){
           if(i!=argc && strstr(argv[i+1],"-")==NULL){
               r = atof(argv[i+1]);
           }
           else{
               fprintf(stderr,"Please provide a valid value with -r argument, "
	               "else remove -r to use default value\n");
               exit(1);
           }
        }

        /* Bucket capicity is B*/
        if(strcmp("-B", argv[i])==0){
            if(i!=argc && strstr(argv[i+1],"-")==NULL){
                B = atol(argv[i+1]);
            }
            else{
                fprintf(stderr,"Please provide a valid value with -B argument,"
		        " else remove -B to use default value\n");
                exit(1);
            }
        }

        /* tokens needed for a packet is P*/
        if(strcmp("-P", argv[i])==0){
            if(i!=argc && strstr(argv[i+1],"-")==NULL){
            P = atof(argv[i+1]);
            }
            else{
                fprintf(stderr,"Please provide a valid value with -P argument, "
		        "else remove -P to use default value\n");
                exit(1);
            }
        }

        /* Total number of packets is num. you can give total num of packet
         * in test file it self
         */
        if(strcmp("-n", argv[i])==0){
            if(i!=argc && strstr(argv[i+1],"-")==NULL){
                num = atof(argv[i+1]);
            }
            else{
                fprintf(stderr,"Please provide a valid value with -num "
		        "argument, else remove -num to use default value\n");
                exit(1);
            }
        }

        /* When you give file name then use -t during execution time  */
        if(strcmp("-t", argv[i])==0){
           if(i!=argc){
               tsfile = malloc(sizeof(argv[i+1]));
               strcpy(tsfile,argv[i+1]);
           }
           else{
               fprintf(stderr,"Please provide a file name with -t argument");
               exit(1);
           }
        }
    }

    /* Initialize all variable's*/
    initialize ();

    /* Process tockens and create thread for further processing
     * Defination of process can be found in token-core.c file.
     */
    process(lambda, mu, r, B, P, num, tsfile);

    return 0;
}


