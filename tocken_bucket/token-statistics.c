
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

void display_statistics(){
    struct timeval current_time;
    gettimeofday(&current_time, NULL);
    float temp = calculate_time_difference(current_time,emulation_start_time);
    printf("%012.3fms: emulation ends\n",temp);
    pthread_mutex_lock(&mutex);
    printf("Total Q1 : %f\n",average_number_of_packets_in_Q1);
    printf("Total Q2 : %f\n",average_number_of_packets_in_Q2);
    average_number_of_packets_in_Q1 = average_number_of_packets_in_Q1 / temp;
    average_number_of_packets_in_Q2 = average_number_of_packets_in_Q2 / temp;
    average_number_of_packets_at_S = average_number_of_packets_at_S / temp;
    printf("\nStatistics:\n\n");
    printf("    average packet inter-arrival time = ");
    if(average_packet_inter_arrival_time >= 0.0){
        printf("%f\n",average_packet_inter_arrival_time / 1000.0);
    }
    else {
        printf("N/A\n");
    }
    printf("    average packet service time = ");
    if(average_packet_service_time >= 0.0){
        printf("%f\n\n",average_packet_service_time / 1000.0);
    }
    else {
        printf("N/A\n\n");
    }
    printf("    average number of packets in Q1 = ");
    if(average_number_of_packets_in_Q1 >= 0.0){
        printf("%f\n",average_number_of_packets_in_Q1);
    }
    else {
        printf("N/A\n");
    }
    printf("    average number of packets in Q2 = ");
    if(average_number_of_packets_in_Q2 >= 0.0){
        printf("%f\n",average_number_of_packets_in_Q2);
    }
    else {
        printf("N/A\n");
    }
    printf("    average number of packets in S = ");
    if(average_number_of_packets_at_S >= 0.0){
        printf("%f\n\n",average_number_of_packets_at_S);
    }
    else {
        printf("N/A\n\n");
    }
    printf("    average time a packet spent in system = ");
    if(average_time_a_packet_spent_in_system >= 0.0){
        printf("%f\n",average_time_a_packet_spent_in_system / 1000.0);
    }
    else {
        printf("N/A\n");
    }
    printf("    standard deviation for the time spent in system = ");
    if(standard_deviation_for_time_spent_in_system >= 0.0){
        printf("%f\n\n",standard_deviation_for_time_spent_in_system / 1000.0);
    }
    else {
        printf("N/A\n\n");
    }
    printf("    token drop probability = ");
    if(token_drop_probability >= 0.0){
        printf("%f\n",token_drop_probability);
    }
    else {
        printf("N/A\n");
    }
    printf("    packet drop probability = ");
    if(packet_drop_probability >= 0.0){
        printf("%f\n",packet_drop_probability);
    }
    else {
        printf("N/A\n");
    }
    pthread_mutex_unlock(&mutex);

}
